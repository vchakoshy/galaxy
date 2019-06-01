package flex

import (
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

const (
	dataProviderTypeProposedList = "PROPOSED_LIST"
)

// ProposedListDataProvider struct
type ProposedListDataProvider struct{}

func (b ProposedListDataProvider) getOutputComponent(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: dataProviderTypeProposedList,
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

	catIds := cs.Settings.Setup.Category.GetIdis()
	if len(catIds) > 0 {
		queries = append(queries, qm.InnerJoin("general_category_assign ON general_category_assign.item_id = propose_book_list.id AND item_type='ProposeBookList'"))
		queries = append(queries, qm.WhereIn("general_category_assign.category_id = ?", catIds...))
	}

	authorIDs := cs.Settings.Setup.Author.GetIdis()
	if len(authorIDs) > 0 {
		queries = append(queries, qm.InnerJoin("book ON book.id = propose_book_list_item.book_id"))
		queries = append(queries, qm.WhereIn("book.author_id in ?", authorIDs...))
	}

	plIDs := cs.Settings.Setup.ProposedList.GetIdis()
	if len(plIDs) > 0 {
		queries = append(queries, qm.WhereIn("propose_book_list.id in ?", plIDs...))
	}

	queries = append(queries, qm.Limit(10))
	queries = append(queries, qm.Load("Author"))

	proposeLists, err := models.ProposeBookLists(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
		return com
	}

	com.Data.Items.Generic = make([]Generic, len(proposeLists))
	for i, v := range proposeLists {
		com.Data.Items.Generic[i] = Generic{
			Title:     v.Title.String,
			SubTitle:  v.SubTitle.String,
			Image:     v.CoverImage.String,
			IconTitle: v.R.Author.Name,
			Icon:      v.R.Author.Logo.String,
			// FooterText: v.R.ProposeBookListItems, // Count proposeBookListItems
			Action: &Action{ // TODO fix actions
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
			ChildAction: &Action{
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

		// com.Data.Items.Model = append(com.Data.Items.Model, v)
	}

	return com
}
