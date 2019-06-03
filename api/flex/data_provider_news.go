package flex

import (
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

const (
	dataProviderTypeNews = "NEWS"
)

// NewsDataProvider struct
type NewsDataProvider struct {
	ComponentSettings ComponentSettings
	ComponentType     string
}

func NewNewsDataProvider(cs ComponentSettings) NewsDataProvider {
	return NewsDataProvider{
		ComponentSettings: cs,
		ComponentType:     "news",
	}
}

func (b NewsDataProvider) GetOutputComponent() OutputComponent {
	return b.getOutputComponent()
}

func (b NewsDataProvider) getOutputComponent() OutputComponent {

	com, _ := NewOutputComponent(
		OutputComponentSetTitle(b.ComponentSettings.Elements.Title.Value.Static),
		OutputComponentSetType(b.ComponentType),
		OutputComponentSetResourceType(dataProviderTypeNews),
	)

	// com := OutputComponent{
	// 	Type:         b.ComponentType,
	// 	ResourceType: dataProviderTypeNews,
	// 	Title:        b.ComponentSettings.Elements.Title.Value.Static,
	// }

	a := b.ComponentSettings.Elements.MoreTitle.Action.getAction()
	if a.Type != "" {
		com.Action = a
	}

	if b.ComponentSettings.Elements.MoreTitle.Value != "" {
		com.ActionTitle = b.ComponentSettings.Elements.MoreTitle.Value
	}

	n, err := b.Models()
	if err != nil {
		log.Println(err.Error())
		return *com
	}

	for i, v := range n.(models.NewsSlice) {
		g := Generic{
			Title: v.Title,
			Image: v.Thumbnail.String,
			// IconSubtitle: , // TODO reading time
			Action: &Action{ // TODO fix actions
				Type:   "news_page",
				Method: "/general/news/get",
				Input: []InputAction{
					InputAction{
						Key:   "newsId",
						Value: v.ID,
					},
					InputAction{
						Key:   "pageName",
						Value: "NEWS_DETAIL_PAGE",
					},
				},
			},
		}

		if v.R.Publisher != nil {
			g.IconTitle = v.R.Publisher.Title
			g.Icon = v.R.Publisher.Logo.String
		}

		com.Data.Items.Generic[i] = g

		// com.Data.Items.Model = append(com.Data.Items.Model, v)
	}

	return *com
}

func (b NewsDataProvider) Models() (r interface{}, err error) {
	queries := []qm.QueryMod{}

	newsIDs := b.ComponentSettings.Settings.Setup.News.GetIdis()
	if len(newsIDs) > 0 {
		queries = append(queries, qm.WhereIn("news.id in ?", newsIDs...))
	}

	queries = append(queries, qm.Limit(10), qm.Load("Publisher"))

	r, err = models.AllNews(queries...).AllG(context.Background())

	return
}
