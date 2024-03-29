// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// BookTag is an object representing the database table.
type BookTag struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	BookID    int       `boil:"book_id" json:"book_id" toml:"book_id" yaml:"book_id"`
	TagID     int       `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *bookTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookTagColumns = struct {
	ID        string
	BookID    string
	TagID     string
	CreatedAt string
}{
	ID:        "id",
	BookID:    "book_id",
	TagID:     "tag_id",
	CreatedAt: "created_at",
}

// Generated where

var BookTagWhere = struct {
	ID        whereHelperint
	BookID    whereHelperint
	TagID     whereHelperint
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "`book_tags`.`id`"},
	BookID:    whereHelperint{field: "`book_tags`.`book_id`"},
	TagID:     whereHelperint{field: "`book_tags`.`tag_id`"},
	CreatedAt: whereHelpertime_Time{field: "`book_tags`.`created_at`"},
}

// BookTagRels is where relationship names are stored.
var BookTagRels = struct {
}{}

// bookTagR is where relationships are stored.
type bookTagR struct {
}

// NewStruct creates a new relationship struct
func (*bookTagR) NewStruct() *bookTagR {
	return &bookTagR{}
}

// bookTagL is where Load methods for each relationship are stored.
type bookTagL struct{}

var (
	bookTagAllColumns            = []string{"id", "book_id", "tag_id", "created_at"}
	bookTagColumnsWithoutDefault = []string{"book_id", "tag_id"}
	bookTagColumnsWithDefault    = []string{"id", "created_at"}
	bookTagPrimaryKeyColumns     = []string{"id"}
)

type (
	// BookTagSlice is an alias for a slice of pointers to BookTag.
	// This should generally be used opposed to []BookTag.
	BookTagSlice []*BookTag
	// BookTagHook is the signature for custom BookTag hook methods
	BookTagHook func(context.Context, boil.ContextExecutor, *BookTag) error

	bookTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookTagType                 = reflect.TypeOf(&BookTag{})
	bookTagMapping              = queries.MakeStructMapping(bookTagType)
	bookTagPrimaryKeyMapping, _ = queries.BindMapping(bookTagType, bookTagMapping, bookTagPrimaryKeyColumns)
	bookTagInsertCacheMut       sync.RWMutex
	bookTagInsertCache          = make(map[string]insertCache)
	bookTagUpdateCacheMut       sync.RWMutex
	bookTagUpdateCache          = make(map[string]updateCache)
	bookTagUpsertCacheMut       sync.RWMutex
	bookTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookTagBeforeInsertHooks []BookTagHook
var bookTagBeforeUpdateHooks []BookTagHook
var bookTagBeforeDeleteHooks []BookTagHook
var bookTagBeforeUpsertHooks []BookTagHook

var bookTagAfterInsertHooks []BookTagHook
var bookTagAfterSelectHooks []BookTagHook
var bookTagAfterUpdateHooks []BookTagHook
var bookTagAfterDeleteHooks []BookTagHook
var bookTagAfterUpsertHooks []BookTagHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookTag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookTag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookTag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookTag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookTag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookTag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookTag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookTag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookTag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookTagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookTagHook registers your hook function for all future operations.
func AddBookTagHook(hookPoint boil.HookPoint, bookTagHook BookTagHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bookTagBeforeInsertHooks = append(bookTagBeforeInsertHooks, bookTagHook)
	case boil.BeforeUpdateHook:
		bookTagBeforeUpdateHooks = append(bookTagBeforeUpdateHooks, bookTagHook)
	case boil.BeforeDeleteHook:
		bookTagBeforeDeleteHooks = append(bookTagBeforeDeleteHooks, bookTagHook)
	case boil.BeforeUpsertHook:
		bookTagBeforeUpsertHooks = append(bookTagBeforeUpsertHooks, bookTagHook)
	case boil.AfterInsertHook:
		bookTagAfterInsertHooks = append(bookTagAfterInsertHooks, bookTagHook)
	case boil.AfterSelectHook:
		bookTagAfterSelectHooks = append(bookTagAfterSelectHooks, bookTagHook)
	case boil.AfterUpdateHook:
		bookTagAfterUpdateHooks = append(bookTagAfterUpdateHooks, bookTagHook)
	case boil.AfterDeleteHook:
		bookTagAfterDeleteHooks = append(bookTagAfterDeleteHooks, bookTagHook)
	case boil.AfterUpsertHook:
		bookTagAfterUpsertHooks = append(bookTagAfterUpsertHooks, bookTagHook)
	}
}

// OneG returns a single bookTag record from the query using the global executor.
func (q bookTagQuery) OneG(ctx context.Context) (*BookTag, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single bookTag record from the query.
func (q bookTagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookTag, error) {
	o := &BookTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for book_tags")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all BookTag records from the query using the global executor.
func (q bookTagQuery) AllG(ctx context.Context) (BookTagSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all BookTag records from the query.
func (q bookTagQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookTagSlice, error) {
	var o []*BookTag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookTag slice")
	}

	if len(bookTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all BookTag records in the query, and panics on error.
func (q bookTagQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all BookTag records in the query.
func (q bookTagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count book_tags rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q bookTagQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q bookTagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if book_tags exists")
	}

	return count > 0, nil
}

// BookTags retrieves all the records using an executor.
func BookTags(mods ...qm.QueryMod) bookTagQuery {
	mods = append(mods, qm.From("`book_tags`"))
	return bookTagQuery{NewQuery(mods...)}
}

// FindBookTagG retrieves a single record by ID.
func FindBookTagG(ctx context.Context, iD int, selectCols ...string) (*BookTag, error) {
	return FindBookTag(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindBookTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookTag(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BookTag, error) {
	bookTagObj := &BookTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `book_tags` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookTagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from book_tags")
	}

	return bookTagObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BookTag) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookTag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_tags provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookTagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookTagInsertCacheMut.RLock()
	cache, cached := bookTagInsertCache[key]
	bookTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookTagAllColumns,
			bookTagColumnsWithDefault,
			bookTagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookTagType, bookTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookTagType, bookTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `book_tags` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `book_tags` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `book_tags` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bookTagPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into book_tags")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bookTagMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for book_tags")
	}

CacheNoHooks:
	if !cached {
		bookTagInsertCacheMut.Lock()
		bookTagInsertCache[key] = cache
		bookTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single BookTag record using the global executor.
// See Update for more documentation.
func (o *BookTag) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the BookTag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookTag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookTagUpdateCacheMut.RLock()
	cache, cached := bookTagUpdateCache[key]
	bookTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookTagAllColumns,
			bookTagPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update book_tags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `book_tags` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bookTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookTagType, bookTagMapping, append(wl, bookTagPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update book_tags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for book_tags")
	}

	if !cached {
		bookTagUpdateCacheMut.Lock()
		bookTagUpdateCache[key] = cache
		bookTagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q bookTagQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q bookTagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for book_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for book_tags")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BookTagSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookTagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `book_tags` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookTagPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookTag")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BookTag) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLBookTagUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookTag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_tags provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookTagColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBookTagUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bookTagUpsertCacheMut.RLock()
	cache, cached := bookTagUpsertCache[key]
	bookTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookTagAllColumns,
			bookTagColumnsWithDefault,
			bookTagColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bookTagAllColumns,
			bookTagPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert book_tags, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "book_tags", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `book_tags` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(bookTagType, bookTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookTagType, bookTagMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for book_tags")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bookTagMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(bookTagType, bookTagMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for book_tags")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for book_tags")
	}

CacheNoHooks:
	if !cached {
		bookTagUpsertCacheMut.Lock()
		bookTagUpsertCache[key] = cache
		bookTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single BookTag record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BookTag) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single BookTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookTag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookTagPrimaryKeyMapping)
	sql := "DELETE FROM `book_tags` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from book_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for book_tags")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookTagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from book_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_tags")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BookTagSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookTagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `book_tags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookTagPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_tags")
	}

	if len(bookTagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BookTag) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no BookTag provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BookTag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookTag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookTagSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty BookTagSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookTagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `book_tags`.* FROM `book_tags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookTagSlice")
	}

	*o = slice

	return nil
}

// BookTagExistsG checks if the BookTag row exists.
func BookTagExistsG(ctx context.Context, iD int) (bool, error) {
	return BookTagExists(ctx, boil.GetContextDB(), iD)
}

// BookTagExists checks if the BookTag row exists.
func BookTagExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `book_tags` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if book_tags exists")
	}

	return exists, nil
}
