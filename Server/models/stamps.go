// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Stamp is an object representing the database table.
type Stamp struct {
	ID        int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time  `boil:"createdAt" json:"createdAt,omitempty" toml:"createdAt" yaml:"createdAt,omitempty"`
	UpdatedAt null.Time  `boil:"updatedAt" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`
	PersonID  null.Int64 `boil:"person_id" json:"person_id,omitempty" toml:"person_id" yaml:"person_id,omitempty"`
	Checkin   bool       `boil:"checkin" json:"checkin" toml:"checkin" yaml:"checkin"`
	Stamp     time.Time  `boil:"stamp" json:"stamp" toml:"stamp" yaml:"stamp"`

	R *stampR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stampL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StampColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	PersonID  string
	Checkin   string
	Stamp     string
}{
	ID:        "id",
	CreatedAt: "createdAt",
	UpdatedAt: "updatedAt",
	PersonID:  "person_id",
	Checkin:   "checkin",
	Stamp:     "stamp",
}

var StampTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	PersonID  string
	Checkin   string
	Stamp     string
}{
	ID:        "stamps.id",
	CreatedAt: "stamps.createdAt",
	UpdatedAt: "stamps.updatedAt",
	PersonID:  "stamps.person_id",
	Checkin:   "stamps.checkin",
	Stamp:     "stamps.stamp",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var StampWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
	PersonID  whereHelpernull_Int64
	Checkin   whereHelperbool
	Stamp     whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"stamps\".\"id\""},
	CreatedAt: whereHelpernull_Time{field: "\"stamps\".\"createdAt\""},
	UpdatedAt: whereHelpernull_Time{field: "\"stamps\".\"updatedAt\""},
	PersonID:  whereHelpernull_Int64{field: "\"stamps\".\"person_id\""},
	Checkin:   whereHelperbool{field: "\"stamps\".\"checkin\""},
	Stamp:     whereHelpertime_Time{field: "\"stamps\".\"stamp\""},
}

// StampRels is where relationship names are stored.
var StampRels = struct {
	Person string
}{
	Person: "Person",
}

// stampR is where relationships are stored.
type stampR struct {
	Person *Person `boil:"Person" json:"Person" toml:"Person" yaml:"Person"`
}

// NewStruct creates a new relationship struct
func (*stampR) NewStruct() *stampR {
	return &stampR{}
}

func (r *stampR) GetPerson() *Person {
	if r == nil {
		return nil
	}
	return r.Person
}

// stampL is where Load methods for each relationship are stored.
type stampL struct{}

var (
	stampAllColumns            = []string{"id", "createdAt", "updatedAt", "person_id", "checkin", "stamp"}
	stampColumnsWithoutDefault = []string{"stamp"}
	stampColumnsWithDefault    = []string{"id", "createdAt", "updatedAt", "person_id", "checkin"}
	stampPrimaryKeyColumns     = []string{"id"}
	stampGeneratedColumns      = []string{}
)

type (
	// StampSlice is an alias for a slice of pointers to Stamp.
	// This should almost always be used instead of []Stamp.
	StampSlice []*Stamp
	// StampHook is the signature for custom Stamp hook methods
	StampHook func(context.Context, boil.ContextExecutor, *Stamp) error

	stampQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stampType                 = reflect.TypeOf(&Stamp{})
	stampMapping              = queries.MakeStructMapping(stampType)
	stampPrimaryKeyMapping, _ = queries.BindMapping(stampType, stampMapping, stampPrimaryKeyColumns)
	stampInsertCacheMut       sync.RWMutex
	stampInsertCache          = make(map[string]insertCache)
	stampUpdateCacheMut       sync.RWMutex
	stampUpdateCache          = make(map[string]updateCache)
	stampUpsertCacheMut       sync.RWMutex
	stampUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var stampAfterSelectHooks []StampHook

var stampBeforeInsertHooks []StampHook
var stampAfterInsertHooks []StampHook

var stampBeforeUpdateHooks []StampHook
var stampAfterUpdateHooks []StampHook

var stampBeforeDeleteHooks []StampHook
var stampAfterDeleteHooks []StampHook

var stampBeforeUpsertHooks []StampHook
var stampAfterUpsertHooks []StampHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stamp) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stamp) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stamp) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stamp) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stamp) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stamp) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stamp) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stamp) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stamp) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stampAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStampHook registers your hook function for all future operations.
func AddStampHook(hookPoint boil.HookPoint, stampHook StampHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		stampAfterSelectHooks = append(stampAfterSelectHooks, stampHook)
	case boil.BeforeInsertHook:
		stampBeforeInsertHooks = append(stampBeforeInsertHooks, stampHook)
	case boil.AfterInsertHook:
		stampAfterInsertHooks = append(stampAfterInsertHooks, stampHook)
	case boil.BeforeUpdateHook:
		stampBeforeUpdateHooks = append(stampBeforeUpdateHooks, stampHook)
	case boil.AfterUpdateHook:
		stampAfterUpdateHooks = append(stampAfterUpdateHooks, stampHook)
	case boil.BeforeDeleteHook:
		stampBeforeDeleteHooks = append(stampBeforeDeleteHooks, stampHook)
	case boil.AfterDeleteHook:
		stampAfterDeleteHooks = append(stampAfterDeleteHooks, stampHook)
	case boil.BeforeUpsertHook:
		stampBeforeUpsertHooks = append(stampBeforeUpsertHooks, stampHook)
	case boil.AfterUpsertHook:
		stampAfterUpsertHooks = append(stampAfterUpsertHooks, stampHook)
	}
}

// One returns a single stamp record from the query.
func (q stampQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Stamp, error) {
	o := &Stamp{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stamps")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Stamp records from the query.
func (q stampQuery) All(ctx context.Context, exec boil.ContextExecutor) (StampSlice, error) {
	var o []*Stamp

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stamp slice")
	}

	if len(stampAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Stamp records in the query.
func (q stampQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stamps rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stampQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stamps exists")
	}

	return count > 0, nil
}

// Person pointed to by the foreign key.
func (o *Stamp) Person(mods ...qm.QueryMod) personQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PersonID),
	}

	queryMods = append(queryMods, mods...)

	return Persons(queryMods...)
}

// LoadPerson allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stampL) LoadPerson(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStamp interface{}, mods queries.Applicator) error {
	var slice []*Stamp
	var object *Stamp

	if singular {
		object = maybeStamp.(*Stamp)
	} else {
		slice = *maybeStamp.(*[]*Stamp)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stampR{}
		}
		if !queries.IsNil(object.PersonID) {
			args = append(args, object.PersonID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stampR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.PersonID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.PersonID) {
				args = append(args, obj.PersonID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`persons`),
		qm.WhereIn(`persons.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Person")
	}

	var resultSlice []*Person
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Person")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for persons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for persons")
	}

	if len(stampAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Person = foreign
		if foreign.R == nil {
			foreign.R = &personR{}
		}
		foreign.R.Stamps = append(foreign.R.Stamps, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.PersonID, foreign.ID) {
				local.R.Person = foreign
				if foreign.R == nil {
					foreign.R = &personR{}
				}
				foreign.R.Stamps = append(foreign.R.Stamps, local)
				break
			}
		}
	}

	return nil
}

// SetPerson of the stamp to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Stamps.
func (o *Stamp) SetPerson(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Person) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stamps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"person_id"}),
		strmangle.WhereClause("\"", "\"", 2, stampPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.PersonID, related.ID)
	if o.R == nil {
		o.R = &stampR{
			Person: related,
		}
	} else {
		o.R.Person = related
	}

	if related.R == nil {
		related.R = &personR{
			Stamps: StampSlice{o},
		}
	} else {
		related.R.Stamps = append(related.R.Stamps, o)
	}

	return nil
}

// RemovePerson relationship.
// Sets o.R.Person to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Stamp) RemovePerson(ctx context.Context, exec boil.ContextExecutor, related *Person) error {
	var err error

	queries.SetScanner(&o.PersonID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("person_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Person = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Stamps {
		if queries.Equal(o.PersonID, ri.PersonID) {
			continue
		}

		ln := len(related.R.Stamps)
		if ln > 1 && i < ln-1 {
			related.R.Stamps[i] = related.R.Stamps[ln-1]
		}
		related.R.Stamps = related.R.Stamps[:ln-1]
		break
	}
	return nil
}

// Stamps retrieves all the records using an executor.
func Stamps(mods ...qm.QueryMod) stampQuery {
	mods = append(mods, qm.From("\"stamps\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"stamps\".*"})
	}

	return stampQuery{q}
}

// FindStamp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStamp(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Stamp, error) {
	stampObj := &Stamp{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stamps\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stampObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stamps")
	}

	if err = stampObj.doAfterSelectHooks(ctx, exec); err != nil {
		return stampObj, err
	}

	return stampObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Stamp) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stamps provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stampColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stampInsertCacheMut.RLock()
	cache, cached := stampInsertCache[key]
	stampInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stampAllColumns,
			stampColumnsWithDefault,
			stampColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stampType, stampMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stampType, stampMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"stamps\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"stamps\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into stamps")
	}

	if !cached {
		stampInsertCacheMut.Lock()
		stampInsertCache[key] = cache
		stampInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Stamp.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Stamp) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stampUpdateCacheMut.RLock()
	cache, cached := stampUpdateCache[key]
	stampUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stampAllColumns,
			stampPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stamps, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stamps\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stampPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stampType, stampMapping, append(wl, stampPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update stamps row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stamps")
	}

	if !cached {
		stampUpdateCacheMut.Lock()
		stampUpdateCache[key] = cache
		stampUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q stampQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stamps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stamps")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StampSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stampPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"stamps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stampPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stamp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stamp")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Stamp) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stamps provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stampColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	stampUpsertCacheMut.RLock()
	cache, cached := stampUpsertCache[key]
	stampUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stampAllColumns,
			stampColumnsWithDefault,
			stampColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			stampAllColumns,
			stampPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stamps, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stampPrimaryKeyColumns))
			copy(conflict, stampPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"stamps\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stampType, stampMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stampType, stampMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert stamps")
	}

	if !cached {
		stampUpsertCacheMut.Lock()
		stampUpsertCache[key] = cache
		stampUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Stamp record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stamp) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stamp provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stampPrimaryKeyMapping)
	sql := "DELETE FROM \"stamps\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stamps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stamps")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stampQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stampQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stamps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stamps")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StampSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(stampBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stampPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"stamps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stampPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stamp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stamps")
	}

	if len(stampAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stamp) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStamp(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StampSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StampSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stampPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"stamps\".* FROM \"stamps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stampPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StampSlice")
	}

	*o = slice

	return nil
}

// StampExists checks if the Stamp row exists.
func StampExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"stamps\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stamps exists")
	}

	return exists, nil
}
