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

func testBookStats(t *testing.T) {
	t.Parallel()

	query := BookStats()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookStatsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
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

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookStatsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BookStats().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookStatsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookStatSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookStatsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookStatExists(ctx, tx, o.BookID)
	if err != nil {
		t.Errorf("Unable to check if BookStat exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookStatExists to return true, but got false.")
	}
}

func testBookStatsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookStatFound, err := FindBookStat(ctx, tx, o.BookID)
	if err != nil {
		t.Error(err)
	}

	if bookStatFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookStatsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BookStats().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookStatsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BookStats().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookStatsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookStatOne := &BookStat{}
	bookStatTwo := &BookStat{}
	if err = randomize.Struct(seed, bookStatOne, bookStatDBTypes, false, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}
	if err = randomize.Struct(seed, bookStatTwo, bookStatDBTypes, false, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookStatOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookStatTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookStats().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookStatsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookStatOne := &BookStat{}
	bookStatTwo := &BookStat{}
	if err = randomize.Struct(seed, bookStatOne, bookStatDBTypes, false, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}
	if err = randomize.Struct(seed, bookStatTwo, bookStatDBTypes, false, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookStatOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookStatTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookStatBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func bookStatAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookStat) error {
	*o = BookStat{}
	return nil
}

func testBookStatsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BookStat{}
	o := &BookStat{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookStatDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookStat object: %s", err)
	}

	AddBookStatHook(boil.BeforeInsertHook, bookStatBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookStatBeforeInsertHooks = []BookStatHook{}

	AddBookStatHook(boil.AfterInsertHook, bookStatAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookStatAfterInsertHooks = []BookStatHook{}

	AddBookStatHook(boil.AfterSelectHook, bookStatAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookStatAfterSelectHooks = []BookStatHook{}

	AddBookStatHook(boil.BeforeUpdateHook, bookStatBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookStatBeforeUpdateHooks = []BookStatHook{}

	AddBookStatHook(boil.AfterUpdateHook, bookStatAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookStatAfterUpdateHooks = []BookStatHook{}

	AddBookStatHook(boil.BeforeDeleteHook, bookStatBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookStatBeforeDeleteHooks = []BookStatHook{}

	AddBookStatHook(boil.AfterDeleteHook, bookStatAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookStatAfterDeleteHooks = []BookStatHook{}

	AddBookStatHook(boil.BeforeUpsertHook, bookStatBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookStatBeforeUpsertHooks = []BookStatHook{}

	AddBookStatHook(boil.AfterUpsertHook, bookStatAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookStatAfterUpsertHooks = []BookStatHook{}
}

func testBookStatsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookStatsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookStatColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookStatToOneBookUsingBook(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BookStat
	var foreign Book

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookStatDBTypes, false, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BookID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Book().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BookStatSlice{&local}
	if err = local.L.LoadBook(ctx, tx, false, (*[]*BookStat)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Book = nil
	if err = local.L.LoadBook(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBookStatToOneSetOpBookUsingBook(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookStat
	var b, c Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookStatDBTypes, false, strmangle.SetComplement(bookStatPrimaryKeyColumns, bookStatColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Book{&b, &c} {
		err = a.SetBook(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Book != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BookStat != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BookID != x.ID {
			t.Error("foreign key was wrong value", a.BookID)
		}

		if exists, err := BookStatExists(ctx, tx, a.BookID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testBookStatsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
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

func testBookStatsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookStatSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookStatsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookStats().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookStatDBTypes = map[string]string{`BookID`: `int`, `DayDownloadCount`: `int`, `WeekDownloadCount`: `int`, `MonthDownloadCount`: `int`, `AllDownloadCount`: `int`, `DaySalesCount`: `int`, `WeekSalesCount`: `int`, `MonthSalesCount`: `int`, `AllSalesCount`: `int`, `AllCommentCount`: `int`, `LastStatsUpdate`: `timestamp`, `LastIndexTime`: `timestamp`, `IndexStatus`: `tinyint`}
	_               = bytes.MinRead
)

func testBookStatsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookStatPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookStatAllColumns) == len(bookStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookStatsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookStatAllColumns) == len(bookStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookStat{}
	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookStatDBTypes, true, bookStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookStatAllColumns, bookStatPrimaryKeyColumns) {
		fields = bookStatAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookStatAllColumns,
			bookStatPrimaryKeyColumns,
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

	slice := BookStatSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookStatsUpsert(t *testing.T) {
	t.Parallel()

	if len(bookStatAllColumns) == len(bookStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBookStatUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BookStat{}
	if err = randomize.Struct(seed, &o, bookStatDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookStat: %s", err)
	}

	count, err := BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookStatDBTypes, false, bookStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookStat struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookStat: %s", err)
	}

	count, err = BookStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
