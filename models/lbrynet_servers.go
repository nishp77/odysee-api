// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
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

// LbrynetServer is an object representing the database table.
type LbrynetServer struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Address   string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	Weight    int       `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *lbrynetServerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lbrynetServerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LbrynetServerColumns = struct {
	ID        string
	Name      string
	Address   string
	Weight    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Address:   "address",
	Weight:    "weight",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

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

var LbrynetServerWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	Address   whereHelperstring
	Weight    whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"lbrynet_servers\".\"id\""},
	Name:      whereHelperstring{field: "\"lbrynet_servers\".\"name\""},
	Address:   whereHelperstring{field: "\"lbrynet_servers\".\"address\""},
	Weight:    whereHelperint{field: "\"lbrynet_servers\".\"weight\""},
	CreatedAt: whereHelpertime_Time{field: "\"lbrynet_servers\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"lbrynet_servers\".\"updated_at\""},
}

// LbrynetServerRels is where relationship names are stored.
var LbrynetServerRels = struct {
	Users string
}{
	Users: "Users",
}

// lbrynetServerR is where relationships are stored.
type lbrynetServerR struct {
	Users UserSlice
}

// NewStruct creates a new relationship struct
func (*lbrynetServerR) NewStruct() *lbrynetServerR {
	return &lbrynetServerR{}
}

// lbrynetServerL is where Load methods for each relationship are stored.
type lbrynetServerL struct{}

var (
	lbrynetServerAllColumns            = []string{"id", "name", "address", "weight", "created_at", "updated_at"}
	lbrynetServerColumnsWithoutDefault = []string{"name", "address"}
	lbrynetServerColumnsWithDefault    = []string{"id", "weight", "created_at", "updated_at"}
	lbrynetServerPrimaryKeyColumns     = []string{"id"}
)

type (
	// LbrynetServerSlice is an alias for a slice of pointers to LbrynetServer.
	// This should generally be used opposed to []LbrynetServer.
	LbrynetServerSlice []*LbrynetServer
	// LbrynetServerHook is the signature for custom LbrynetServer hook methods
	LbrynetServerHook func(boil.Executor, *LbrynetServer) error

	lbrynetServerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lbrynetServerType                 = reflect.TypeOf(&LbrynetServer{})
	lbrynetServerMapping              = queries.MakeStructMapping(lbrynetServerType)
	lbrynetServerPrimaryKeyMapping, _ = queries.BindMapping(lbrynetServerType, lbrynetServerMapping, lbrynetServerPrimaryKeyColumns)
	lbrynetServerInsertCacheMut       sync.RWMutex
	lbrynetServerInsertCache          = make(map[string]insertCache)
	lbrynetServerUpdateCacheMut       sync.RWMutex
	lbrynetServerUpdateCache          = make(map[string]updateCache)
	lbrynetServerUpsertCacheMut       sync.RWMutex
	lbrynetServerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var lbrynetServerBeforeInsertHooks []LbrynetServerHook
var lbrynetServerBeforeUpdateHooks []LbrynetServerHook
var lbrynetServerBeforeDeleteHooks []LbrynetServerHook
var lbrynetServerBeforeUpsertHooks []LbrynetServerHook

var lbrynetServerAfterInsertHooks []LbrynetServerHook
var lbrynetServerAfterSelectHooks []LbrynetServerHook
var lbrynetServerAfterUpdateHooks []LbrynetServerHook
var lbrynetServerAfterDeleteHooks []LbrynetServerHook
var lbrynetServerAfterUpsertHooks []LbrynetServerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LbrynetServer) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LbrynetServer) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LbrynetServer) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LbrynetServer) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LbrynetServer) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LbrynetServer) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LbrynetServer) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LbrynetServer) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LbrynetServer) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range lbrynetServerAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLbrynetServerHook registers your hook function for all future operations.
func AddLbrynetServerHook(hookPoint boil.HookPoint, lbrynetServerHook LbrynetServerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		lbrynetServerBeforeInsertHooks = append(lbrynetServerBeforeInsertHooks, lbrynetServerHook)
	case boil.BeforeUpdateHook:
		lbrynetServerBeforeUpdateHooks = append(lbrynetServerBeforeUpdateHooks, lbrynetServerHook)
	case boil.BeforeDeleteHook:
		lbrynetServerBeforeDeleteHooks = append(lbrynetServerBeforeDeleteHooks, lbrynetServerHook)
	case boil.BeforeUpsertHook:
		lbrynetServerBeforeUpsertHooks = append(lbrynetServerBeforeUpsertHooks, lbrynetServerHook)
	case boil.AfterInsertHook:
		lbrynetServerAfterInsertHooks = append(lbrynetServerAfterInsertHooks, lbrynetServerHook)
	case boil.AfterSelectHook:
		lbrynetServerAfterSelectHooks = append(lbrynetServerAfterSelectHooks, lbrynetServerHook)
	case boil.AfterUpdateHook:
		lbrynetServerAfterUpdateHooks = append(lbrynetServerAfterUpdateHooks, lbrynetServerHook)
	case boil.AfterDeleteHook:
		lbrynetServerAfterDeleteHooks = append(lbrynetServerAfterDeleteHooks, lbrynetServerHook)
	case boil.AfterUpsertHook:
		lbrynetServerAfterUpsertHooks = append(lbrynetServerAfterUpsertHooks, lbrynetServerHook)
	}
}

// OneG returns a single lbrynetServer record from the query using the global executor.
func (q lbrynetServerQuery) OneG() (*LbrynetServer, error) {
	return q.One(boil.GetDB())
}

// One returns a single lbrynetServer record from the query.
func (q lbrynetServerQuery) One(exec boil.Executor) (*LbrynetServer, error) {
	o := &LbrynetServer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for lbrynet_servers")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all LbrynetServer records from the query using the global executor.
func (q lbrynetServerQuery) AllG() (LbrynetServerSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all LbrynetServer records from the query.
func (q lbrynetServerQuery) All(exec boil.Executor) (LbrynetServerSlice, error) {
	var o []*LbrynetServer

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LbrynetServer slice")
	}

	if len(lbrynetServerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all LbrynetServer records in the query, and panics on error.
func (q lbrynetServerQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all LbrynetServer records in the query.
func (q lbrynetServerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count lbrynet_servers rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q lbrynetServerQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q lbrynetServerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if lbrynet_servers exists")
	}

	return count > 0, nil
}

// Users retrieves all the user's Users with an executor.
func (o *LbrynetServer) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"users\".\"lbrynet_server_id\"=?", o.ID),
	)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"users\".*"})
	}

	return query
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (lbrynetServerL) LoadUsers(e boil.Executor, singular bool, maybeLbrynetServer interface{}, mods queries.Applicator) error {
	var slice []*LbrynetServer
	var object *LbrynetServer

	if singular {
		object = maybeLbrynetServer.(*LbrynetServer)
	} else {
		slice = *maybeLbrynetServer.(*[]*LbrynetServer)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lbrynetServerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lbrynetServerR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`lbrynet_server_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.LbrynetServer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.LbrynetServerID) {
				local.R.Users = append(local.R.Users, foreign)
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.LbrynetServer = local
				break
			}
		}
	}

	return nil
}

// AddUsersG adds the given related objects to the existing relationships
// of the lbrynet_server, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.LbrynetServer appropriately.
// Uses the global database handle.
func (o *LbrynetServer) AddUsersG(insert bool, related ...*User) error {
	return o.AddUsers(boil.GetDB(), insert, related...)
}

// AddUsers adds the given related objects to the existing relationships
// of the lbrynet_server, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.LbrynetServer appropriately.
func (o *LbrynetServer) AddUsers(exec boil.Executor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.LbrynetServerID, o.ID)
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"lbrynet_server_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.LbrynetServerID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &lbrynetServerR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				LbrynetServer: o,
			}
		} else {
			rel.R.LbrynetServer = o
		}
	}
	return nil
}

// SetUsersG removes all previously related items of the
// lbrynet_server replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.LbrynetServer's Users accordingly.
// Replaces o.R.Users with related.
// Sets related.R.LbrynetServer's Users accordingly.
// Uses the global database handle.
func (o *LbrynetServer) SetUsersG(insert bool, related ...*User) error {
	return o.SetUsers(boil.GetDB(), insert, related...)
}

// SetUsers removes all previously related items of the
// lbrynet_server replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.LbrynetServer's Users accordingly.
// Replaces o.R.Users with related.
// Sets related.R.LbrynetServer's Users accordingly.
func (o *LbrynetServer) SetUsers(exec boil.Executor, insert bool, related ...*User) error {
	query := "update \"users\" set \"lbrynet_server_id\" = null where \"lbrynet_server_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Users {
			queries.SetScanner(&rel.LbrynetServerID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.LbrynetServer = nil
		}

		o.R.Users = nil
	}
	return o.AddUsers(exec, insert, related...)
}

// RemoveUsersG relationships from objects passed in.
// Removes related items from R.Users (uses pointer comparison, removal does not keep order)
// Sets related.R.LbrynetServer.
// Uses the global database handle.
func (o *LbrynetServer) RemoveUsersG(related ...*User) error {
	return o.RemoveUsers(boil.GetDB(), related...)
}

// RemoveUsers relationships from objects passed in.
// Removes related items from R.Users (uses pointer comparison, removal does not keep order)
// Sets related.R.LbrynetServer.
func (o *LbrynetServer) RemoveUsers(exec boil.Executor, related ...*User) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.LbrynetServerID, nil)
		if rel.R != nil {
			rel.R.LbrynetServer = nil
		}
		if _, err = rel.Update(exec, boil.Whitelist("lbrynet_server_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Users {
			if rel != ri {
				continue
			}

			ln := len(o.R.Users)
			if ln > 1 && i < ln-1 {
				o.R.Users[i] = o.R.Users[ln-1]
			}
			o.R.Users = o.R.Users[:ln-1]
			break
		}
	}

	return nil
}

// LbrynetServers retrieves all the records using an executor.
func LbrynetServers(mods ...qm.QueryMod) lbrynetServerQuery {
	mods = append(mods, qm.From("\"lbrynet_servers\""))
	return lbrynetServerQuery{NewQuery(mods...)}
}

// FindLbrynetServerG retrieves a single record by ID.
func FindLbrynetServerG(iD int, selectCols ...string) (*LbrynetServer, error) {
	return FindLbrynetServer(boil.GetDB(), iD, selectCols...)
}

// FindLbrynetServer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLbrynetServer(exec boil.Executor, iD int, selectCols ...string) (*LbrynetServer, error) {
	lbrynetServerObj := &LbrynetServer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lbrynet_servers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, lbrynetServerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from lbrynet_servers")
	}

	return lbrynetServerObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *LbrynetServer) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LbrynetServer) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lbrynet_servers provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(lbrynetServerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lbrynetServerInsertCacheMut.RLock()
	cache, cached := lbrynetServerInsertCache[key]
	lbrynetServerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lbrynetServerAllColumns,
			lbrynetServerColumnsWithDefault,
			lbrynetServerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lbrynetServerType, lbrynetServerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lbrynetServerType, lbrynetServerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lbrynet_servers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lbrynet_servers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into lbrynet_servers")
	}

	if !cached {
		lbrynetServerInsertCacheMut.Lock()
		lbrynetServerInsertCache[key] = cache
		lbrynetServerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single LbrynetServer record using the global executor.
// See Update for more documentation.
func (o *LbrynetServer) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the LbrynetServer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LbrynetServer) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	lbrynetServerUpdateCacheMut.RLock()
	cache, cached := lbrynetServerUpdateCache[key]
	lbrynetServerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lbrynetServerAllColumns,
			lbrynetServerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update lbrynet_servers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lbrynet_servers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lbrynetServerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lbrynetServerType, lbrynetServerMapping, append(wl, lbrynetServerPrimaryKeyColumns...))
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
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update lbrynet_servers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for lbrynet_servers")
	}

	if !cached {
		lbrynetServerUpdateCacheMut.Lock()
		lbrynetServerUpdateCache[key] = cache
		lbrynetServerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q lbrynetServerQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q lbrynetServerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for lbrynet_servers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for lbrynet_servers")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o LbrynetServerSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LbrynetServerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lbrynetServerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lbrynet_servers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lbrynetServerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in lbrynetServer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all lbrynetServer")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *LbrynetServer) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LbrynetServer) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lbrynet_servers provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(lbrynetServerColumnsWithDefault, o)

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

	lbrynetServerUpsertCacheMut.RLock()
	cache, cached := lbrynetServerUpsertCache[key]
	lbrynetServerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lbrynetServerAllColumns,
			lbrynetServerColumnsWithDefault,
			lbrynetServerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lbrynetServerAllColumns,
			lbrynetServerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert lbrynet_servers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lbrynetServerPrimaryKeyColumns))
			copy(conflict, lbrynetServerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lbrynet_servers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lbrynetServerType, lbrynetServerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lbrynetServerType, lbrynetServerMapping, ret)
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert lbrynet_servers")
	}

	if !cached {
		lbrynetServerUpsertCacheMut.Lock()
		lbrynetServerUpsertCache[key] = cache
		lbrynetServerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single LbrynetServer record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *LbrynetServer) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single LbrynetServer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LbrynetServer) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LbrynetServer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lbrynetServerPrimaryKeyMapping)
	sql := "DELETE FROM \"lbrynet_servers\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from lbrynet_servers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for lbrynet_servers")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lbrynetServerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no lbrynetServerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lbrynet_servers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lbrynet_servers")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o LbrynetServerSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LbrynetServerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(lbrynetServerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lbrynetServerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lbrynet_servers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lbrynetServerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lbrynetServer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lbrynet_servers")
	}

	if len(lbrynetServerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *LbrynetServer) ReloadG() error {
	if o == nil {
		return errors.New("models: no LbrynetServer provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LbrynetServer) Reload(exec boil.Executor) error {
	ret, err := FindLbrynetServer(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LbrynetServerSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty LbrynetServerSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LbrynetServerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LbrynetServerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lbrynetServerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lbrynet_servers\".* FROM \"lbrynet_servers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lbrynetServerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LbrynetServerSlice")
	}

	*o = slice

	return nil
}

// LbrynetServerExistsG checks if the LbrynetServer row exists.
func LbrynetServerExistsG(iD int) (bool, error) {
	return LbrynetServerExists(boil.GetDB(), iD)
}

// LbrynetServerExists checks if the LbrynetServer row exists.
func LbrynetServerExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lbrynet_servers\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if lbrynet_servers exists")
	}

	return exists, nil
}
