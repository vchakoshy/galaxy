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

func testFlexPages(t *testing.T) {
	t.Parallel()

	query := FlexPages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFlexPagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
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

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlexPagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FlexPages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlexPagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FlexPageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlexPagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FlexPageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if FlexPage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FlexPageExists to return true, but got false.")
	}
}

func testFlexPagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	flexPageFound, err := FindFlexPage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if flexPageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFlexPagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FlexPages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFlexPagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FlexPages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFlexPagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	flexPageOne := &FlexPage{}
	flexPageTwo := &FlexPage{}
	if err = randomize.Struct(seed, flexPageOne, flexPageDBTypes, false, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}
	if err = randomize.Struct(seed, flexPageTwo, flexPageDBTypes, false, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = flexPageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = flexPageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FlexPages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFlexPagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	flexPageOne := &FlexPage{}
	flexPageTwo := &FlexPage{}
	if err = randomize.Struct(seed, flexPageOne, flexPageDBTypes, false, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}
	if err = randomize.Struct(seed, flexPageTwo, flexPageDBTypes, false, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = flexPageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = flexPageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func flexPageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func flexPageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FlexPage) error {
	*o = FlexPage{}
	return nil
}

func testFlexPagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FlexPage{}
	o := &FlexPage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, flexPageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FlexPage object: %s", err)
	}

	AddFlexPageHook(boil.BeforeInsertHook, flexPageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	flexPageBeforeInsertHooks = []FlexPageHook{}

	AddFlexPageHook(boil.AfterInsertHook, flexPageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	flexPageAfterInsertHooks = []FlexPageHook{}

	AddFlexPageHook(boil.AfterSelectHook, flexPageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	flexPageAfterSelectHooks = []FlexPageHook{}

	AddFlexPageHook(boil.BeforeUpdateHook, flexPageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	flexPageBeforeUpdateHooks = []FlexPageHook{}

	AddFlexPageHook(boil.AfterUpdateHook, flexPageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	flexPageAfterUpdateHooks = []FlexPageHook{}

	AddFlexPageHook(boil.BeforeDeleteHook, flexPageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	flexPageBeforeDeleteHooks = []FlexPageHook{}

	AddFlexPageHook(boil.AfterDeleteHook, flexPageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	flexPageAfterDeleteHooks = []FlexPageHook{}

	AddFlexPageHook(boil.BeforeUpsertHook, flexPageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	flexPageBeforeUpsertHooks = []FlexPageHook{}

	AddFlexPageHook(boil.AfterUpsertHook, flexPageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	flexPageAfterUpsertHooks = []FlexPageHook{}
}

func testFlexPagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFlexPagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(flexPageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFlexPageToManyPageFlexPageComponents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FlexPage
	var b, c FlexPageComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, flexPageComponentDBTypes, false, flexPageComponentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, flexPageComponentDBTypes, false, flexPageComponentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PageID = a.ID
	c.PageID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PageFlexPageComponents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PageID == b.PageID {
			bFound = true
		}
		if v.PageID == c.PageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FlexPageSlice{&a}
	if err = a.L.LoadPageFlexPageComponents(ctx, tx, false, (*[]*FlexPage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PageFlexPageComponents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PageFlexPageComponents = nil
	if err = a.L.LoadPageFlexPageComponents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PageFlexPageComponents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testFlexPageToManyAddOpPageFlexPageComponents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FlexPage
	var b, c, d, e FlexPageComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, flexPageDBTypes, false, strmangle.SetComplement(flexPagePrimaryKeyColumns, flexPageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*FlexPageComponent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, flexPageComponentDBTypes, false, strmangle.SetComplement(flexPageComponentPrimaryKeyColumns, flexPageComponentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*FlexPageComponent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPageFlexPageComponents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PageID {
			t.Error("foreign key was wrong value", a.ID, first.PageID)
		}
		if a.ID != second.PageID {
			t.Error("foreign key was wrong value", a.ID, second.PageID)
		}

		if first.R.Page != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Page != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PageFlexPageComponents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PageFlexPageComponents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PageFlexPageComponents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testFlexPagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
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

func testFlexPagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FlexPageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFlexPagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FlexPages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	flexPageDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `StoreID`: `int`, `TemplateID`: `int`, `CreatedAt`: `datetime`, `Title`: `varchar`, `SettingData`: `text`, `Type`: `enum('STATIC','DYNAMIC')`, `Description`: `text`, `Image`: `varchar`}
	_               = bytes.MinRead
)

func testFlexPagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(flexPagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(flexPageColumns) == len(flexPagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFlexPagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(flexPageColumns) == len(flexPagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FlexPage{}
	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, flexPageDBTypes, true, flexPagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(flexPageColumns, flexPagePrimaryKeyColumns) {
		fields = flexPageColumns
	} else {
		fields = strmangle.SetComplement(
			flexPageColumns,
			flexPagePrimaryKeyColumns,
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

	slice := FlexPageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFlexPagesUpsert(t *testing.T) {
	t.Parallel()

	if len(flexPageColumns) == len(flexPagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFlexPageUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FlexPage{}
	if err = randomize.Struct(seed, &o, flexPageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FlexPage: %s", err)
	}

	count, err := FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, flexPageDBTypes, false, flexPagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FlexPage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FlexPage: %s", err)
	}

	count, err = FlexPages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
