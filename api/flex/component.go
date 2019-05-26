package flex

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
	"gitlab.fidibo.com/backend/galaxy/hubble"
)

// Component for flex
type Component struct {
	queries []qm.QueryMod
	pageID  int
}

func NewComponentByPage(pageID int) *Component {
	return &Component{
		pageID: pageID,
	}
}

func (c *Component) GetData() []flexComponent {
	qComponent := qm.Where("page_id=?", c.pageID)

	fpc, err := models.FlexPageComponents(
		qComponent,
		qm.OrderBy("crud_order asc")).
		AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	componenets := make([]flexComponent, 0)

	for _, comp := range fpc {

		compModel, err := models.
			FlexComponents(qm.Where("id=?", comp.ComponentID)).
			OneG(context.Background())
		if err != nil {
			log.Println(err.Error(), comp.ComponentID)
			continue
		}

		cs := flexComponentSettings{}

		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)

		if cs.Settings.DataProvider == "BOOK" {
			com := flexComponent{
				Type:         compModel.Type,
				ResourceType: "BOOK",
				Title:        cs.Elements.Title.Value.Static,
				Action:       getAction(cs),
			}

			if cs.Elements.MoreTitle.Value != "" {
				com.ActionTitle = cs.Elements.MoreTitle.Value
			}

			bookIdis := QueryIdis{}

			ss := esClient.Search(hubble.ProductIndexName)

			switch cs.Settings.Setup.Sort.Value {
			case "BESTSELLER":
				ss.Sort("all_sales_count", false)

			case "WEEK_BESTSELLER":
				ss.Sort("week_sales_count", false)

			case "POPULAR":
				ss.Sort("month_download_count", false)

			case "RECENT":
				ss.Sort("publish_time", false)

			case "MOST_COMMENTED":
				ss.Sort("all_comment_count", false)

			default:
				ss.Sort("publish_time", false)
			}

			inList := cs.Settings.Setup.Format.Value.getInterfaceList()
			if len(inList) > 0 {
				elastic.NewBoolQuery().Should(
					elastic.NewTermsQuery("format", inList),
				)
				// queries = append(queries, qm.WhereIn("format in ?", inList...))
			}

			esres, err := ss.StoredFields("_id").Size(8).Do(context.Background())
			if err != nil {
				log.Println(err.Error())
			}

			// spew.Dump(esres)

			for _, item := range esres.Hits.Hits {
				itemID := strings.Replace(item.Id, "book-", "", 1)
				bookIdis = append(bookIdis, itemID)
			}

			queries := []qm.QueryMod{}

			if len(bookIdis) > 0 {
				queries = append(queries, qm.WhereIn("id in ?", bookIdis...))
			}

			queries = append(queries, qm.Limit(8))

			res := newGenericBookByQuery(queries)

			com.Data.Items.Generic = make([]interface{}, len(res))
			for i, v := range res {
				com.Data.Items.Generic[i] = v
			}

			componenets = append(componenets, com)
		}
	}
	return componenets
}

func getAction(cs flexComponentSettings) flexBaseAction {
	action := GetActionByPanelType(cs.Elements.MoreTitle.Action.Type)

	return flexBaseAction{
		Type:      action.ClientType,
		Input:     cs.Elements.MoreTitle.Action.Setup.getInputActions(),
		ExtraData: nil,
		Method:    action.Method,
	}
}
