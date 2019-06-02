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
	Type              string
}

func (b NewsDataProvider) getOutputComponent() OutputComponent {
	com := OutputComponent{
		Type:         b.Type,
		ResourceType: dataProviderTypeNews,
		Title:        b.ComponentSettings.Elements.Title.Value.Static,
	}

	a := b.ComponentSettings.Elements.MoreTitle.Action.getAction()
	if a.Type != "" {
		com.Action = a
	}

	if b.ComponentSettings.Elements.MoreTitle.Value != "" {
		com.ActionTitle = b.ComponentSettings.Elements.MoreTitle.Value
	}

	news, err := b.Models()
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

	return com
}

func (b NewsDataProvider) Models() (r models.NewsSlice, err error) {
	queries := []qm.QueryMod{}

	newsIDs := b.ComponentSettings.Settings.Setup.News.GetIdis()
	if len(newsIDs) > 0 {
		queries = append(queries, qm.WhereIn("news.id in ?", newsIDs...))
	}

	queries = append(queries, qm.Limit(10), qm.Load("Publisher"))

	r, err = models.AllNews(queries...).AllG(context.Background())

	return
}
