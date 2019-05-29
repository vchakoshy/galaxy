package flex

import (
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

type DataProvidersProposedList struct{}

func (b DataProvidersProposedList) getGeneric(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: "PROPOSED_LIST",
		Title:        cs.Elements.Title.Value.Static,
	}

	a := getAction(cs.Elements.MoreTitle.Action)
	if a.Type != "" {
		com.Action = a
	}

	if cs.Elements.MoreTitle.Value != "" {
		com.ActionTitle = cs.Elements.MoreTitle.Value
	}

	queries := []qm.QueryMod{}

	catIds := cs.Settings.Setup.Category.GetIdis()
	if len(catIds) > 0 {
		qm.InnerJoin("general_category_assign ON general_category_assign.item_id = propose_book_list.id AND item_type='ProposeBookList'")
		qm.WhereIn("general_category_assign.category_id", catIds)
	}

	authorIDs := cs.Settings.Setup.Author.GetIdis()
	if len(authorIDs) > 0 {
		qm.InnerJoin("book ON book.id = propose_book_list_item.book_id")
		qm.WhereIn("book.author_id", authorIDs)
	}

	plIDs := cs.Settings.Setup.ProposedList.GetIdis()
	if len(plIDs) > 0 {
		qm.WhereIn("propose_book_list.id", plIDs)
	}

	queries = append(queries, qm.Limit(10))
	queries = append(queries, qm.Load("Author"))

	proposeList, err := models.ProposeBookLists(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
		return com
	}

	com.Data.Items.Generic = make([]Generic, len(proposeList))
	for i, v := range proposeList {
		com.Data.Items.Generic[i] = Generic{
			Title:     v.Title.String,
			SubTitle:  v.SubTitle.String,
			Image:     v.CoverImage.String,
			IconTitle: v.R.Author.Name,
			Icon:      v.R.Author.Logo.String,
			// FooterText: v., // TODO count
			Action: &BaseAction{ // TODO fix actions
				Type:   "proposed_list_page",
				Method: "/general/proposed-list/get",
				Input: []ComponentAction{
					ComponentAction{
						Key:   "proposedListId",
						Value: v.ID,
					},
					ComponentAction{
						Key:   "pageName",
						Value: "READING_LIST_PAGE",
					},
				},
			},
			ChildAction: &BaseAction{
				Type:   "proposed_list_page",
				Method: "/general/proposed-list/get",
				Input: []ComponentAction{
					ComponentAction{
						Key:   "proposedListId",
						Value: v.ID,
					},
					ComponentAction{
						Key:   "pageName",
						Value: "READING_LIST_PAGE",
					},
				},
			},
		}
	}

	return com
}
