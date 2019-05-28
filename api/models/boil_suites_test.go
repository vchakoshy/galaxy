// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Authors", testAuthors)
	t.Run("Books", testBooks)
	t.Run("BookCategories", testBookCategories)
	t.Run("BookCategoryAssigns", testBookCategoryAssigns)
	t.Run("BookStats", testBookStats)
	t.Run("BookTags", testBookTags)
	t.Run("BooksTags", testBooksTags)
	t.Run("FlexComponents", testFlexComponents)
	t.Run("FlexPageComponents", testFlexPageComponents)
	t.Run("FlexPages", testFlexPages)
	t.Run("ProposeBookLists", testProposeBookLists)
	t.Run("ProposeBookListItems", testProposeBookListItems)
	t.Run("Publishers", testPublishers)
}

func TestDelete(t *testing.T) {
	t.Run("Authors", testAuthorsDelete)
	t.Run("Books", testBooksDelete)
	t.Run("BookCategories", testBookCategoriesDelete)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsDelete)
	t.Run("BookStats", testBookStatsDelete)
	t.Run("BookTags", testBookTagsDelete)
	t.Run("BooksTags", testBooksTagsDelete)
	t.Run("FlexComponents", testFlexComponentsDelete)
	t.Run("FlexPageComponents", testFlexPageComponentsDelete)
	t.Run("FlexPages", testFlexPagesDelete)
	t.Run("ProposeBookLists", testProposeBookListsDelete)
	t.Run("ProposeBookListItems", testProposeBookListItemsDelete)
	t.Run("Publishers", testPublishersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Authors", testAuthorsQueryDeleteAll)
	t.Run("Books", testBooksQueryDeleteAll)
	t.Run("BookCategories", testBookCategoriesQueryDeleteAll)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsQueryDeleteAll)
	t.Run("BookStats", testBookStatsQueryDeleteAll)
	t.Run("BookTags", testBookTagsQueryDeleteAll)
	t.Run("BooksTags", testBooksTagsQueryDeleteAll)
	t.Run("FlexComponents", testFlexComponentsQueryDeleteAll)
	t.Run("FlexPageComponents", testFlexPageComponentsQueryDeleteAll)
	t.Run("FlexPages", testFlexPagesQueryDeleteAll)
	t.Run("ProposeBookLists", testProposeBookListsQueryDeleteAll)
	t.Run("ProposeBookListItems", testProposeBookListItemsQueryDeleteAll)
	t.Run("Publishers", testPublishersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Authors", testAuthorsSliceDeleteAll)
	t.Run("Books", testBooksSliceDeleteAll)
	t.Run("BookCategories", testBookCategoriesSliceDeleteAll)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsSliceDeleteAll)
	t.Run("BookStats", testBookStatsSliceDeleteAll)
	t.Run("BookTags", testBookTagsSliceDeleteAll)
	t.Run("BooksTags", testBooksTagsSliceDeleteAll)
	t.Run("FlexComponents", testFlexComponentsSliceDeleteAll)
	t.Run("FlexPageComponents", testFlexPageComponentsSliceDeleteAll)
	t.Run("FlexPages", testFlexPagesSliceDeleteAll)
	t.Run("ProposeBookLists", testProposeBookListsSliceDeleteAll)
	t.Run("ProposeBookListItems", testProposeBookListItemsSliceDeleteAll)
	t.Run("Publishers", testPublishersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Authors", testAuthorsExists)
	t.Run("Books", testBooksExists)
	t.Run("BookCategories", testBookCategoriesExists)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsExists)
	t.Run("BookStats", testBookStatsExists)
	t.Run("BookTags", testBookTagsExists)
	t.Run("BooksTags", testBooksTagsExists)
	t.Run("FlexComponents", testFlexComponentsExists)
	t.Run("FlexPageComponents", testFlexPageComponentsExists)
	t.Run("FlexPages", testFlexPagesExists)
	t.Run("ProposeBookLists", testProposeBookListsExists)
	t.Run("ProposeBookListItems", testProposeBookListItemsExists)
	t.Run("Publishers", testPublishersExists)
}

func TestFind(t *testing.T) {
	t.Run("Authors", testAuthorsFind)
	t.Run("Books", testBooksFind)
	t.Run("BookCategories", testBookCategoriesFind)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsFind)
	t.Run("BookStats", testBookStatsFind)
	t.Run("BookTags", testBookTagsFind)
	t.Run("BooksTags", testBooksTagsFind)
	t.Run("FlexComponents", testFlexComponentsFind)
	t.Run("FlexPageComponents", testFlexPageComponentsFind)
	t.Run("FlexPages", testFlexPagesFind)
	t.Run("ProposeBookLists", testProposeBookListsFind)
	t.Run("ProposeBookListItems", testProposeBookListItemsFind)
	t.Run("Publishers", testPublishersFind)
}

func TestBind(t *testing.T) {
	t.Run("Authors", testAuthorsBind)
	t.Run("Books", testBooksBind)
	t.Run("BookCategories", testBookCategoriesBind)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsBind)
	t.Run("BookStats", testBookStatsBind)
	t.Run("BookTags", testBookTagsBind)
	t.Run("BooksTags", testBooksTagsBind)
	t.Run("FlexComponents", testFlexComponentsBind)
	t.Run("FlexPageComponents", testFlexPageComponentsBind)
	t.Run("FlexPages", testFlexPagesBind)
	t.Run("ProposeBookLists", testProposeBookListsBind)
	t.Run("ProposeBookListItems", testProposeBookListItemsBind)
	t.Run("Publishers", testPublishersBind)
}

func TestOne(t *testing.T) {
	t.Run("Authors", testAuthorsOne)
	t.Run("Books", testBooksOne)
	t.Run("BookCategories", testBookCategoriesOne)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsOne)
	t.Run("BookStats", testBookStatsOne)
	t.Run("BookTags", testBookTagsOne)
	t.Run("BooksTags", testBooksTagsOne)
	t.Run("FlexComponents", testFlexComponentsOne)
	t.Run("FlexPageComponents", testFlexPageComponentsOne)
	t.Run("FlexPages", testFlexPagesOne)
	t.Run("ProposeBookLists", testProposeBookListsOne)
	t.Run("ProposeBookListItems", testProposeBookListItemsOne)
	t.Run("Publishers", testPublishersOne)
}

func TestAll(t *testing.T) {
	t.Run("Authors", testAuthorsAll)
	t.Run("Books", testBooksAll)
	t.Run("BookCategories", testBookCategoriesAll)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsAll)
	t.Run("BookStats", testBookStatsAll)
	t.Run("BookTags", testBookTagsAll)
	t.Run("BooksTags", testBooksTagsAll)
	t.Run("FlexComponents", testFlexComponentsAll)
	t.Run("FlexPageComponents", testFlexPageComponentsAll)
	t.Run("FlexPages", testFlexPagesAll)
	t.Run("ProposeBookLists", testProposeBookListsAll)
	t.Run("ProposeBookListItems", testProposeBookListItemsAll)
	t.Run("Publishers", testPublishersAll)
}

func TestCount(t *testing.T) {
	t.Run("Authors", testAuthorsCount)
	t.Run("Books", testBooksCount)
	t.Run("BookCategories", testBookCategoriesCount)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsCount)
	t.Run("BookStats", testBookStatsCount)
	t.Run("BookTags", testBookTagsCount)
	t.Run("BooksTags", testBooksTagsCount)
	t.Run("FlexComponents", testFlexComponentsCount)
	t.Run("FlexPageComponents", testFlexPageComponentsCount)
	t.Run("FlexPages", testFlexPagesCount)
	t.Run("ProposeBookLists", testProposeBookListsCount)
	t.Run("ProposeBookListItems", testProposeBookListItemsCount)
	t.Run("Publishers", testPublishersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Authors", testAuthorsHooks)
	t.Run("Books", testBooksHooks)
	t.Run("BookCategories", testBookCategoriesHooks)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsHooks)
	t.Run("BookStats", testBookStatsHooks)
	t.Run("BookTags", testBookTagsHooks)
	t.Run("BooksTags", testBooksTagsHooks)
	t.Run("FlexComponents", testFlexComponentsHooks)
	t.Run("FlexPageComponents", testFlexPageComponentsHooks)
	t.Run("FlexPages", testFlexPagesHooks)
	t.Run("ProposeBookLists", testProposeBookListsHooks)
	t.Run("ProposeBookListItems", testProposeBookListItemsHooks)
	t.Run("Publishers", testPublishersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Authors", testAuthorsInsert)
	t.Run("Authors", testAuthorsInsertWhitelist)
	t.Run("Books", testBooksInsert)
	t.Run("Books", testBooksInsertWhitelist)
	t.Run("BookCategories", testBookCategoriesInsert)
	t.Run("BookCategories", testBookCategoriesInsertWhitelist)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsInsert)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsInsertWhitelist)
	t.Run("BookStats", testBookStatsInsert)
	t.Run("BookStats", testBookStatsInsertWhitelist)
	t.Run("BookTags", testBookTagsInsert)
	t.Run("BookTags", testBookTagsInsertWhitelist)
	t.Run("BooksTags", testBooksTagsInsert)
	t.Run("BooksTags", testBooksTagsInsertWhitelist)
	t.Run("FlexComponents", testFlexComponentsInsert)
	t.Run("FlexComponents", testFlexComponentsInsertWhitelist)
	t.Run("FlexPageComponents", testFlexPageComponentsInsert)
	t.Run("FlexPageComponents", testFlexPageComponentsInsertWhitelist)
	t.Run("FlexPages", testFlexPagesInsert)
	t.Run("FlexPages", testFlexPagesInsertWhitelist)
	t.Run("ProposeBookLists", testProposeBookListsInsert)
	t.Run("ProposeBookLists", testProposeBookListsInsertWhitelist)
	t.Run("ProposeBookListItems", testProposeBookListItemsInsert)
	t.Run("ProposeBookListItems", testProposeBookListItemsInsertWhitelist)
	t.Run("Publishers", testPublishersInsert)
	t.Run("Publishers", testPublishersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BookToPublisherUsingPublisher", testBookToOnePublisherUsingPublisher)
	t.Run("BookToAuthorUsingAuthor", testBookToOneAuthorUsingAuthor)
	t.Run("BookToAuthorUsingTranslator", testBookToOneAuthorUsingTranslator)
	t.Run("BookToAuthorUsingAuthor2", testBookToOneAuthorUsingAuthor2)
	t.Run("BookToAuthorUsingAuthor3", testBookToOneAuthorUsingAuthor3)
	t.Run("BookToAuthorUsingTranslator2", testBookToOneAuthorUsingTranslator2)
	t.Run("BookToAuthorUsingTranslator3", testBookToOneAuthorUsingTranslator3)
	t.Run("BookToPublisherUsingOriginalPublisher", testBookToOnePublisherUsingOriginalPublisher)
	t.Run("BookCategoryToBookCategoryUsingParent", testBookCategoryToOneBookCategoryUsingParent)
	t.Run("BookCategoryAssignToBookCategoryUsingCategory", testBookCategoryAssignToOneBookCategoryUsingCategory)
	t.Run("BookCategoryAssignToBookUsingBook", testBookCategoryAssignToOneBookUsingBook)
	t.Run("BookStatToBookUsingBook", testBookStatToOneBookUsingBook)
	t.Run("FlexPageComponentToFlexComponentUsingComponent", testFlexPageComponentToOneFlexComponentUsingComponent)
	t.Run("FlexPageComponentToFlexPageUsingPage", testFlexPageComponentToOneFlexPageUsingPage)
	t.Run("ProposeBookListToAuthorUsingAuthor", testProposeBookListToOneAuthorUsingAuthor)
	t.Run("ProposeBookListItemToBookUsingBook", testProposeBookListItemToOneBookUsingBook)
	t.Run("ProposeBookListItemToProposeBookListUsingProposeBookList", testProposeBookListItemToOneProposeBookListUsingProposeBookList)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("BookToBookStatUsingBookStat", testBookOneToOneBookStatUsingBookStat)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AuthorToBooks", testAuthorToManyBooks)
	t.Run("AuthorToTranslatorBooks", testAuthorToManyTranslatorBooks)
	t.Run("AuthorToAuthor2Books", testAuthorToManyAuthor2Books)
	t.Run("AuthorToAuthor3Books", testAuthorToManyAuthor3Books)
	t.Run("AuthorToTranslator2Books", testAuthorToManyTranslator2Books)
	t.Run("AuthorToTranslator3Books", testAuthorToManyTranslator3Books)
	t.Run("AuthorToProposeBookLists", testAuthorToManyProposeBookLists)
	t.Run("BookToBookCategoryAssigns", testBookToManyBookCategoryAssigns)
	t.Run("BookToProposeBookListItems", testBookToManyProposeBookListItems)
	t.Run("BookCategoryToParentBookCategories", testBookCategoryToManyParentBookCategories)
	t.Run("BookCategoryToCategoryBookCategoryAssigns", testBookCategoryToManyCategoryBookCategoryAssigns)
	t.Run("FlexComponentToComponentFlexPageComponents", testFlexComponentToManyComponentFlexPageComponents)
	t.Run("FlexPageToPageFlexPageComponents", testFlexPageToManyPageFlexPageComponents)
	t.Run("ProposeBookListToProposeBookListItems", testProposeBookListToManyProposeBookListItems)
	t.Run("PublisherToBooks", testPublisherToManyBooks)
	t.Run("PublisherToOriginalPublisherBooks", testPublisherToManyOriginalPublisherBooks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BookToPublisherUsingBooks", testBookToOneSetOpPublisherUsingPublisher)
	t.Run("BookToAuthorUsingBooks", testBookToOneSetOpAuthorUsingAuthor)
	t.Run("BookToAuthorUsingTranslatorBooks", testBookToOneSetOpAuthorUsingTranslator)
	t.Run("BookToAuthorUsingAuthor2Books", testBookToOneSetOpAuthorUsingAuthor2)
	t.Run("BookToAuthorUsingAuthor3Books", testBookToOneSetOpAuthorUsingAuthor3)
	t.Run("BookToAuthorUsingTranslator2Books", testBookToOneSetOpAuthorUsingTranslator2)
	t.Run("BookToAuthorUsingTranslator3Books", testBookToOneSetOpAuthorUsingTranslator3)
	t.Run("BookToPublisherUsingOriginalPublisherBooks", testBookToOneSetOpPublisherUsingOriginalPublisher)
	t.Run("BookCategoryToBookCategoryUsingParentBookCategories", testBookCategoryToOneSetOpBookCategoryUsingParent)
	t.Run("BookCategoryAssignToBookCategoryUsingCategoryBookCategoryAssigns", testBookCategoryAssignToOneSetOpBookCategoryUsingCategory)
	t.Run("BookCategoryAssignToBookUsingBookCategoryAssigns", testBookCategoryAssignToOneSetOpBookUsingBook)
	t.Run("BookStatToBookUsingBookStat", testBookStatToOneSetOpBookUsingBook)
	t.Run("FlexPageComponentToFlexComponentUsingComponentFlexPageComponents", testFlexPageComponentToOneSetOpFlexComponentUsingComponent)
	t.Run("FlexPageComponentToFlexPageUsingPageFlexPageComponents", testFlexPageComponentToOneSetOpFlexPageUsingPage)
	t.Run("ProposeBookListToAuthorUsingProposeBookLists", testProposeBookListToOneSetOpAuthorUsingAuthor)
	t.Run("ProposeBookListItemToBookUsingProposeBookListItems", testProposeBookListItemToOneSetOpBookUsingBook)
	t.Run("ProposeBookListItemToProposeBookListUsingProposeBookListItems", testProposeBookListItemToOneSetOpProposeBookListUsingProposeBookList)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("BookToPublisherUsingBooks", testBookToOneRemoveOpPublisherUsingPublisher)
	t.Run("BookToAuthorUsingBooks", testBookToOneRemoveOpAuthorUsingAuthor)
	t.Run("BookToAuthorUsingTranslatorBooks", testBookToOneRemoveOpAuthorUsingTranslator)
	t.Run("BookToAuthorUsingAuthor2Books", testBookToOneRemoveOpAuthorUsingAuthor2)
	t.Run("BookToAuthorUsingAuthor3Books", testBookToOneRemoveOpAuthorUsingAuthor3)
	t.Run("BookToAuthorUsingTranslator2Books", testBookToOneRemoveOpAuthorUsingTranslator2)
	t.Run("BookToAuthorUsingTranslator3Books", testBookToOneRemoveOpAuthorUsingTranslator3)
	t.Run("BookToPublisherUsingOriginalPublisherBooks", testBookToOneRemoveOpPublisherUsingOriginalPublisher)
	t.Run("BookCategoryToBookCategoryUsingParentBookCategories", testBookCategoryToOneRemoveOpBookCategoryUsingParent)
	t.Run("ProposeBookListToAuthorUsingProposeBookLists", testProposeBookListToOneRemoveOpAuthorUsingAuthor)
	t.Run("ProposeBookListItemToProposeBookListUsingProposeBookListItems", testProposeBookListItemToOneRemoveOpProposeBookListUsingProposeBookList)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("BookToBookStatUsingBookStat", testBookOneToOneSetOpBookStatUsingBookStat)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AuthorToBooks", testAuthorToManyAddOpBooks)
	t.Run("AuthorToTranslatorBooks", testAuthorToManyAddOpTranslatorBooks)
	t.Run("AuthorToAuthor2Books", testAuthorToManyAddOpAuthor2Books)
	t.Run("AuthorToAuthor3Books", testAuthorToManyAddOpAuthor3Books)
	t.Run("AuthorToTranslator2Books", testAuthorToManyAddOpTranslator2Books)
	t.Run("AuthorToTranslator3Books", testAuthorToManyAddOpTranslator3Books)
	t.Run("AuthorToProposeBookLists", testAuthorToManyAddOpProposeBookLists)
	t.Run("BookToBookCategoryAssigns", testBookToManyAddOpBookCategoryAssigns)
	t.Run("BookToProposeBookListItems", testBookToManyAddOpProposeBookListItems)
	t.Run("BookCategoryToParentBookCategories", testBookCategoryToManyAddOpParentBookCategories)
	t.Run("BookCategoryToCategoryBookCategoryAssigns", testBookCategoryToManyAddOpCategoryBookCategoryAssigns)
	t.Run("FlexComponentToComponentFlexPageComponents", testFlexComponentToManyAddOpComponentFlexPageComponents)
	t.Run("FlexPageToPageFlexPageComponents", testFlexPageToManyAddOpPageFlexPageComponents)
	t.Run("ProposeBookListToProposeBookListItems", testProposeBookListToManyAddOpProposeBookListItems)
	t.Run("PublisherToBooks", testPublisherToManyAddOpBooks)
	t.Run("PublisherToOriginalPublisherBooks", testPublisherToManyAddOpOriginalPublisherBooks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AuthorToBooks", testAuthorToManySetOpBooks)
	t.Run("AuthorToTranslatorBooks", testAuthorToManySetOpTranslatorBooks)
	t.Run("AuthorToAuthor2Books", testAuthorToManySetOpAuthor2Books)
	t.Run("AuthorToAuthor3Books", testAuthorToManySetOpAuthor3Books)
	t.Run("AuthorToTranslator2Books", testAuthorToManySetOpTranslator2Books)
	t.Run("AuthorToTranslator3Books", testAuthorToManySetOpTranslator3Books)
	t.Run("AuthorToProposeBookLists", testAuthorToManySetOpProposeBookLists)
	t.Run("BookCategoryToParentBookCategories", testBookCategoryToManySetOpParentBookCategories)
	t.Run("ProposeBookListToProposeBookListItems", testProposeBookListToManySetOpProposeBookListItems)
	t.Run("PublisherToBooks", testPublisherToManySetOpBooks)
	t.Run("PublisherToOriginalPublisherBooks", testPublisherToManySetOpOriginalPublisherBooks)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AuthorToBooks", testAuthorToManyRemoveOpBooks)
	t.Run("AuthorToTranslatorBooks", testAuthorToManyRemoveOpTranslatorBooks)
	t.Run("AuthorToAuthor2Books", testAuthorToManyRemoveOpAuthor2Books)
	t.Run("AuthorToAuthor3Books", testAuthorToManyRemoveOpAuthor3Books)
	t.Run("AuthorToTranslator2Books", testAuthorToManyRemoveOpTranslator2Books)
	t.Run("AuthorToTranslator3Books", testAuthorToManyRemoveOpTranslator3Books)
	t.Run("AuthorToProposeBookLists", testAuthorToManyRemoveOpProposeBookLists)
	t.Run("BookCategoryToParentBookCategories", testBookCategoryToManyRemoveOpParentBookCategories)
	t.Run("ProposeBookListToProposeBookListItems", testProposeBookListToManyRemoveOpProposeBookListItems)
	t.Run("PublisherToBooks", testPublisherToManyRemoveOpBooks)
	t.Run("PublisherToOriginalPublisherBooks", testPublisherToManyRemoveOpOriginalPublisherBooks)
}

func TestReload(t *testing.T) {
	t.Run("Authors", testAuthorsReload)
	t.Run("Books", testBooksReload)
	t.Run("BookCategories", testBookCategoriesReload)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsReload)
	t.Run("BookStats", testBookStatsReload)
	t.Run("BookTags", testBookTagsReload)
	t.Run("BooksTags", testBooksTagsReload)
	t.Run("FlexComponents", testFlexComponentsReload)
	t.Run("FlexPageComponents", testFlexPageComponentsReload)
	t.Run("FlexPages", testFlexPagesReload)
	t.Run("ProposeBookLists", testProposeBookListsReload)
	t.Run("ProposeBookListItems", testProposeBookListItemsReload)
	t.Run("Publishers", testPublishersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Authors", testAuthorsReloadAll)
	t.Run("Books", testBooksReloadAll)
	t.Run("BookCategories", testBookCategoriesReloadAll)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsReloadAll)
	t.Run("BookStats", testBookStatsReloadAll)
	t.Run("BookTags", testBookTagsReloadAll)
	t.Run("BooksTags", testBooksTagsReloadAll)
	t.Run("FlexComponents", testFlexComponentsReloadAll)
	t.Run("FlexPageComponents", testFlexPageComponentsReloadAll)
	t.Run("FlexPages", testFlexPagesReloadAll)
	t.Run("ProposeBookLists", testProposeBookListsReloadAll)
	t.Run("ProposeBookListItems", testProposeBookListItemsReloadAll)
	t.Run("Publishers", testPublishersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Authors", testAuthorsSelect)
	t.Run("Books", testBooksSelect)
	t.Run("BookCategories", testBookCategoriesSelect)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsSelect)
	t.Run("BookStats", testBookStatsSelect)
	t.Run("BookTags", testBookTagsSelect)
	t.Run("BooksTags", testBooksTagsSelect)
	t.Run("FlexComponents", testFlexComponentsSelect)
	t.Run("FlexPageComponents", testFlexPageComponentsSelect)
	t.Run("FlexPages", testFlexPagesSelect)
	t.Run("ProposeBookLists", testProposeBookListsSelect)
	t.Run("ProposeBookListItems", testProposeBookListItemsSelect)
	t.Run("Publishers", testPublishersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Authors", testAuthorsUpdate)
	t.Run("Books", testBooksUpdate)
	t.Run("BookCategories", testBookCategoriesUpdate)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsUpdate)
	t.Run("BookStats", testBookStatsUpdate)
	t.Run("BookTags", testBookTagsUpdate)
	t.Run("BooksTags", testBooksTagsUpdate)
	t.Run("FlexComponents", testFlexComponentsUpdate)
	t.Run("FlexPageComponents", testFlexPageComponentsUpdate)
	t.Run("FlexPages", testFlexPagesUpdate)
	t.Run("ProposeBookLists", testProposeBookListsUpdate)
	t.Run("ProposeBookListItems", testProposeBookListItemsUpdate)
	t.Run("Publishers", testPublishersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Authors", testAuthorsSliceUpdateAll)
	t.Run("Books", testBooksSliceUpdateAll)
	t.Run("BookCategories", testBookCategoriesSliceUpdateAll)
	t.Run("BookCategoryAssigns", testBookCategoryAssignsSliceUpdateAll)
	t.Run("BookStats", testBookStatsSliceUpdateAll)
	t.Run("BookTags", testBookTagsSliceUpdateAll)
	t.Run("BooksTags", testBooksTagsSliceUpdateAll)
	t.Run("FlexComponents", testFlexComponentsSliceUpdateAll)
	t.Run("FlexPageComponents", testFlexPageComponentsSliceUpdateAll)
	t.Run("FlexPages", testFlexPagesSliceUpdateAll)
	t.Run("ProposeBookLists", testProposeBookListsSliceUpdateAll)
	t.Run("ProposeBookListItems", testProposeBookListItemsSliceUpdateAll)
	t.Run("Publishers", testPublishersSliceUpdateAll)
}
