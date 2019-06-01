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
type NewsDataProvider struct{}

func (b NewsDataProvider) getOutputComponent(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: dataProviderTypeNews,
		Title:        cs.Elements.Title.Value.Static,
	}

	a := cs.Elements.MoreTitle.Action.getAction()
	if a.Type != "" {
		com.Action = a
	}

	if cs.Elements.MoreTitle.Value != "" {
		com.ActionTitle = cs.Elements.MoreTitle.Value
	}

	queries := []qm.QueryMod{}

	newsIDs := cs.Settings.Setup.News.GetIdis()
	if len(newsIDs) > 0 {
		queries = append(queries, qm.WhereIn("news.id in ?", newsIDs...))
	}

	queries = append(queries, qm.Limit(10), qm.Load("Publisher"))

	news, err := models.AllNews(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
		return com
	}

	com.Data.Items.Generic = make([]Generic, len(news))
	for i, v := range news {
		g := Generic{
			Title: v.Title,
			Image: v.Thumbnail.String,
			// IconSubtitle: , // TODO reading time
			Action: &BaseAction{ // TODO fix actions
				Type:   "news_page",
				Method: "/general/news/get",
				Input: []ComponentAction{
					ComponentAction{
						Key:   "newsId",
						Value: v.ID,
					},
					ComponentAction{
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

	return com
}
