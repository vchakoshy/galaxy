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
type ProposedListDataProvider struct {
	ComponentSettings ComponentSettings
	Type              string
}

func (b ProposedListDataProvider) getOutputComponent() OutputComponent {
	com := OutputComponent{
		Type:         b.Type,
		ResourceType: dataProviderTypeProposedList,
		Title:        b.ComponentSettings.Elements.Title.Value.Static,
	}

	a := b.ComponentSettings.Elements.MoreTitle.Action.getAction()
	if a.Type != "" {
		com.Action = a
	}

	if b.ComponentSettings.Elements.MoreTitle.Value != "" {
		com.ActionTitle = b.ComponentSettings.Elements.MoreTitle.Value
	}

	proposeLists, err := b.Models()
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
			// FooterText: v.R.ProposeBookListItems, // TODO Count proposeBookListItems
			Action: &Action{ // TODO fix actions
				Type:   "proposed_list_page",
				Method: "/general/proposed-list/get",
				Input: []InputAction{
					InputAction{
						Key:   "proposedListId",
						Value: v.ID,
					},
					InputAction{
						Key:   "pageName",
						Value: "READING_LIST_PAGE",
					},
				},
			},
			ChildAction: &Action{
				Type:   "proposed_list_page",
				Method: "/general/proposed-list/get",
				Input: []InputAction{
					InputAction{
						Key:   "proposedListId",
						Value: v.ID,
					},
					InputAction{
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

func (b ProposedListDataProvider) Models() (r models.ProposeBookListSlice, err error) {
	queries := []qm.QueryMod{}

	catIds := b.ComponentSettings.Settings.Setup.Category.GetIdis()
	if len(catIds) > 0 {
		queries = append(queries, qm.InnerJoin("general_category_assign ON general_category_assign.item_id = propose_book_list.id AND item_type='ProposeBookList'"))
		queries = append(queries, qm.WhereIn("general_category_assign.category_id = ?", catIds...))
	}

	authorIDs := b.ComponentSettings.Settings.Setup.Author.GetIdis()
	if len(authorIDs) > 0 {
		queries = append(queries, qm.InnerJoin("book ON book.id = propose_book_list_item.book_id"))
		queries = append(queries, qm.WhereIn("book.author_id in ?", authorIDs...))
	}

	plIDs := b.ComponentSettings.Settings.Setup.ProposedList.GetIdis()
	if len(plIDs) > 0 {
		queries = append(queries, qm.WhereIn("propose_book_list.id in ?", plIDs...))
	}

	queries = append(queries, qm.Limit(10))
	queries = append(queries, qm.Load("Author"))

	r, err = models.ProposeBookLists(queries...).AllG(context.Background())

	return
}
