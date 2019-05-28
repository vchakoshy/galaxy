// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("Authors", testAuthorsUpsert)

	t.Run("Books", testBooksUpsert)

	t.Run("BookCategories", testBookCategoriesUpsert)

	t.Run("BookCategoryAssigns", testBookCategoryAssignsUpsert)

	t.Run("BookStats", testBookStatsUpsert)

	t.Run("BookTags", testBookTagsUpsert)

	t.Run("BooksTags", testBooksTagsUpsert)

	t.Run("FlexComponents", testFlexComponentsUpsert)

	t.Run("FlexPageComponents", testFlexPageComponentsUpsert)

	t.Run("FlexPages", testFlexPagesUpsert)

	t.Run("ProposeBookLists", testProposeBookListsUpsert)

	t.Run("ProposeBookListItems", testProposeBookListItemsUpsert)

	t.Run("Publishers", testPublishersUpsert)
}
