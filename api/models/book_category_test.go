// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBookCategories(t *testing.T) {
	t.Parallel()

	query := BookCategories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BookCategories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookCategorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookCategoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BookCategory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookCategoryExists to return true, but got false.")
	}
}

func testBookCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookCategoryFound, err := FindBookCategory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookCategoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BookCategories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BookCategories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookCategoryOne := &BookCategory{}
	bookCategoryTwo := &BookCategory{}
	if err = randomize.Struct(seed, bookCategoryOne, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, bookCategoryTwo, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookCategoryOne := &BookCategory{}
	bookCategoryTwo := &BookCategory{}
	if err = randomize.Struct(seed, bookCategoryOne, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, bookCategoryTwo, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookCategoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func testBookCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BookCategory{}
	o := &BookCategory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookCategory object: %s", err)
	}

	AddBookCategoryHook(boil.BeforeInsertHook, bookCategoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeInsertHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterInsertHook, bookCategoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterInsertHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterSelectHook, bookCategoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterSelectHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.BeforeUpdateHook, bookCategoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeUpdateHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterUpdateHook, bookCategoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterUpdateHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.BeforeDeleteHook, bookCategoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeDeleteHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterDeleteHook, bookCategoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterDeleteHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.BeforeUpsertHook, bookCategoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeUpsertHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterUpsertHook, bookCategoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterUpsertHooks = []BookCategoryHook{}
}

func testBookCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookCategoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookCategoryToManyParentBookCategories(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ParentID, a.ID)
	queries.Assign(&c.ParentID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ParentBookCategories().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ParentID, b.ParentID) {
			bFound = true
		}
		if queries.Equal(v.ParentID, c.ParentID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookCategorySlice{&a}
	if err = a.L.LoadParentBookCategories(ctx, tx, false, (*[]*BookCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ParentBookCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ParentBookCategories = nil
	if err = a.L.LoadParentBookCategories(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ParentBookCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookCategoryToManyCategoryBookCategoryAssigns(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c BookCategoryAssign

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookCategoryAssignDBTypes, false, bookCategoryAssignColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookCategoryAssignDBTypes, false, bookCategoryAssignColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CategoryID = a.ID
	c.CategoryID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CategoryBookCategoryAssigns().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.CategoryID == b.CategoryID {
			bFound = true
		}
		if v.CategoryID == c.CategoryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookCategorySlice{&a}
	if err = a.L.LoadCategoryBookCategoryAssigns(ctx, tx, false, (*[]*BookCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryBookCategoryAssigns); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CategoryBookCategoryAssigns = nil
	if err = a.L.LoadCategoryBookCategoryAssigns(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryBookCategoryAssigns); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookCategoryToManyAddOpParentBookCategories(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c, d, e BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookCategory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BookCategory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddParentBookCategories(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ParentID) {
			t.Error("foreign key was wrong value", a.ID, first.ParentID)
		}
		if !queries.Equal(a.ID, second.ParentID) {
			t.Error("foreign key was wrong value", a.ID, second.ParentID)
		}

		if first.R.Parent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Parent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ParentBookCategories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ParentBookCategories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ParentBookCategories().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBookCategoryToManySetOpParentBookCategories(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c, d, e BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookCategory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetParentBookCategories(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ParentBookCategories().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetParentBookCategories(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ParentBookCategories().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ParentID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ParentID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ParentID) {
		t.Error("foreign key was wrong value", a.ID, d.ParentID)
	}
	if !queries.Equal(a.ID, e.ParentID) {
		t.Error("foreign key was wrong value", a.ID, e.ParentID)
	}

	if b.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Parent != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Parent != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ParentBookCategories[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ParentBookCategories[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBookCategoryToManyRemoveOpParentBookCategories(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c, d, e BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookCategory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddParentBookCategories(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ParentBookCategories().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveParentBookCategories(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ParentBookCategories().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ParentID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ParentID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Parent != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Parent != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ParentBookCategories) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ParentBookCategories[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ParentBookCategories[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBookCategoryToManyAddOpCategoryBookCategoryAssigns(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c, d, e BookCategoryAssign

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookCategoryAssign{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookCategoryAssignDBTypes, false, strmangle.SetComplement(bookCategoryAssignPrimaryKeyColumns, bookCategoryAssignColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BookCategoryAssign{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCategoryBookCategoryAssigns(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CategoryID {
			t.Error("foreign key was wrong value", a.ID, first.CategoryID)
		}
		if a.ID != second.CategoryID {
			t.Error("foreign key was wrong value", a.ID, second.CategoryID)
		}

		if first.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CategoryBookCategoryAssigns[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CategoryBookCategoryAssigns[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CategoryBookCategoryAssigns().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBookCategoryToOneBookCategoryUsingParent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BookCategory
	var foreign BookCategory

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ParentID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Parent().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BookCategorySlice{&local}
	if err = local.L.LoadParent(ctx, tx, false, (*[]*BookCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Parent = nil
	if err = local.L.LoadParent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBookCategoryToOneSetOpBookCategoryUsingParent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BookCategory{&b, &c} {
		err = a.SetParent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Parent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentBookCategories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ParentID, x.ID) {
			t.Error("foreign key was wrong value", a.ParentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ParentID))
		reflect.Indirect(reflect.ValueOf(&a.ParentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ParentID, x.ID) {
			t.Error("foreign key was wrong value", a.ParentID, x.ID)
		}
	}
}

func testBookCategoryToOneRemoveOpBookCategoryUsingParent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetParent(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveParent(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Parent().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Parent != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ParentID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ParentBookCategories) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBookCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookCategorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookCategoryDBTypes = map[string]string{`ID`: `int`, `Level`: `int`, `Title`: `varchar`, `EnglishTitle`: `varchar`, `Slug`: `varchar`, `LongSlug`: `varchar`, `Subject`: `varchar`, `ParentID`: `int`, `Position`: `int`, `StoresID`: `int`, `BisacID`: `int`, `Status`: `tinyint`, `Code`: `varchar`, `BisacTitle`: `varchar`, `ContentFormat`: `enum('EPUB','PDF','AUDIO')`, `Image`: `varchar`, `TotalBooks`: `int`, `SeoTitle`: `varchar`, `SeoDescription`: `text`, `Canonical`: `text`, `SeoFrontShow`: `tinyint`}
	_                   = bytes.MinRead
)

func testBookCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookCategoryColumns) == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookCategoryColumns) == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookCategoryColumns, bookCategoryPrimaryKeyColumns) {
		fields = bookCategoryColumns
	} else {
		fields = strmangle.SetComplement(
			bookCategoryColumns,
			bookCategoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BookCategorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(bookCategoryColumns) == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBookCategoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BookCategory{}
	if err = randomize.Struct(seed, &o, bookCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookCategory: %s", err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookCategoryDBTypes, false, bookCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookCategory: %s", err)
	}

	count, err = BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}