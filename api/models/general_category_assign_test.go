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

func testGeneralCategoryAssigns(t *testing.T) {
	t.Parallel()

	query := GeneralCategoryAssigns()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGeneralCategoryAssignsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
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

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGeneralCategoryAssignsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := GeneralCategoryAssigns().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGeneralCategoryAssignsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GeneralCategoryAssignSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGeneralCategoryAssignsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GeneralCategoryAssignExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if GeneralCategoryAssign exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GeneralCategoryAssignExists to return true, but got false.")
	}
}

func testGeneralCategoryAssignsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	generalCategoryAssignFound, err := FindGeneralCategoryAssign(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if generalCategoryAssignFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGeneralCategoryAssignsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = GeneralCategoryAssigns().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGeneralCategoryAssignsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := GeneralCategoryAssigns().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGeneralCategoryAssignsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	generalCategoryAssignOne := &GeneralCategoryAssign{}
	generalCategoryAssignTwo := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, generalCategoryAssignOne, generalCategoryAssignDBTypes, false, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}
	if err = randomize.Struct(seed, generalCategoryAssignTwo, generalCategoryAssignDBTypes, false, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = generalCategoryAssignOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = generalCategoryAssignTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GeneralCategoryAssigns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGeneralCategoryAssignsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	generalCategoryAssignOne := &GeneralCategoryAssign{}
	generalCategoryAssignTwo := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, generalCategoryAssignOne, generalCategoryAssignDBTypes, false, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}
	if err = randomize.Struct(seed, generalCategoryAssignTwo, generalCategoryAssignDBTypes, false, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = generalCategoryAssignOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = generalCategoryAssignTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func generalCategoryAssignBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func generalCategoryAssignAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GeneralCategoryAssign) error {
	*o = GeneralCategoryAssign{}
	return nil
}

func testGeneralCategoryAssignsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &GeneralCategoryAssign{}
	o := &GeneralCategoryAssign{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign object: %s", err)
	}

	AddGeneralCategoryAssignHook(boil.BeforeInsertHook, generalCategoryAssignBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignBeforeInsertHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.AfterInsertHook, generalCategoryAssignAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignAfterInsertHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.AfterSelectHook, generalCategoryAssignAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignAfterSelectHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.BeforeUpdateHook, generalCategoryAssignBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignBeforeUpdateHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.AfterUpdateHook, generalCategoryAssignAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignAfterUpdateHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.BeforeDeleteHook, generalCategoryAssignBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignBeforeDeleteHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.AfterDeleteHook, generalCategoryAssignAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignAfterDeleteHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.BeforeUpsertHook, generalCategoryAssignBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignBeforeUpsertHooks = []GeneralCategoryAssignHook{}

	AddGeneralCategoryAssignHook(boil.AfterUpsertHook, generalCategoryAssignAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	generalCategoryAssignAfterUpsertHooks = []GeneralCategoryAssignHook{}
}

func testGeneralCategoryAssignsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGeneralCategoryAssignsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(generalCategoryAssignColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGeneralCategoryAssignsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
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

func testGeneralCategoryAssignsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GeneralCategoryAssignSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGeneralCategoryAssignsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GeneralCategoryAssigns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	generalCategoryAssignDBTypes = map[string]string{`ID`: `int`, `ItemType`: `varchar`, `ItemID`: `int`, `CategoryID`: `int`, `CreatedAt`: `datetime`}
	_                            = bytes.MinRead
)

func testGeneralCategoryAssignsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(generalCategoryAssignPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(generalCategoryAssignAllColumns) == len(generalCategoryAssignPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGeneralCategoryAssignsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(generalCategoryAssignAllColumns) == len(generalCategoryAssignPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GeneralCategoryAssign{}
	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, generalCategoryAssignDBTypes, true, generalCategoryAssignPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(generalCategoryAssignAllColumns, generalCategoryAssignPrimaryKeyColumns) {
		fields = generalCategoryAssignAllColumns
	} else {
		fields = strmangle.SetComplement(
			generalCategoryAssignAllColumns,
			generalCategoryAssignPrimaryKeyColumns,
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

	slice := GeneralCategoryAssignSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGeneralCategoryAssignsUpsert(t *testing.T) {
	t.Parallel()

	if len(generalCategoryAssignAllColumns) == len(generalCategoryAssignPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLGeneralCategoryAssignUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := GeneralCategoryAssign{}
	if err = randomize.Struct(seed, &o, generalCategoryAssignDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GeneralCategoryAssign: %s", err)
	}

	count, err := GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, generalCategoryAssignDBTypes, false, generalCategoryAssignPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GeneralCategoryAssign struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GeneralCategoryAssign: %s", err)
	}

	count, err = GeneralCategoryAssigns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}