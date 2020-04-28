// Code generated by SQLBoiler 4.0.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// NotificationTemplate is an object representing the database table.
type NotificationTemplate struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Text      string    `boil:"text" json:"text" toml:"text" yaml:"text"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *notificationTemplateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L notificationTemplateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NotificationTemplateColumns = struct {
	ID        string
	Text      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Text:      "text",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var NotificationTemplateWhere = struct {
	ID        whereHelperstring
	Text      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"notification_templates\".\"id\""},
	Text:      whereHelperstring{field: "\"notification_templates\".\"text\""},
	CreatedAt: whereHelpertime_Time{field: "\"notification_templates\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"notification_templates\".\"updated_at\""},
}

// NotificationTemplateRels is where relationship names are stored.
var NotificationTemplateRels = struct {
	RoleNotificationTemplates string
}{
	RoleNotificationTemplates: "RoleNotificationTemplates",
}

// notificationTemplateR is where relationships are stored.
type notificationTemplateR struct {
	RoleNotificationTemplates RoleNotificationTemplateSlice
}

// NewStruct creates a new relationship struct
func (*notificationTemplateR) NewStruct() *notificationTemplateR {
	return &notificationTemplateR{}
}

// notificationTemplateL is where Load methods for each relationship are stored.
type notificationTemplateL struct{}

var (
	notificationTemplateAllColumns            = []string{"id", "text", "created_at", "updated_at"}
	notificationTemplateColumnsWithoutDefault = []string{"text", "created_at", "updated_at"}
	notificationTemplateColumnsWithDefault    = []string{"id"}
	notificationTemplatePrimaryKeyColumns     = []string{"id"}
)

type (
	// NotificationTemplateSlice is an alias for a slice of pointers to NotificationTemplate.
	// This should generally be used opposed to []NotificationTemplate.
	NotificationTemplateSlice []*NotificationTemplate

	notificationTemplateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	notificationTemplateType                 = reflect.TypeOf(&NotificationTemplate{})
	notificationTemplateMapping              = queries.MakeStructMapping(notificationTemplateType)
	notificationTemplatePrimaryKeyMapping, _ = queries.BindMapping(notificationTemplateType, notificationTemplateMapping, notificationTemplatePrimaryKeyColumns)
	notificationTemplateInsertCacheMut       sync.RWMutex
	notificationTemplateInsertCache          = make(map[string]insertCache)
	notificationTemplateUpdateCacheMut       sync.RWMutex
	notificationTemplateUpdateCache          = make(map[string]updateCache)
	notificationTemplateUpsertCacheMut       sync.RWMutex
	notificationTemplateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single notificationTemplate record from the query.
func (q notificationTemplateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NotificationTemplate, error) {
	o := &NotificationTemplate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for notification_templates")
	}

	return o, nil
}

// All returns all NotificationTemplate records from the query.
func (q notificationTemplateQuery) All(ctx context.Context, exec boil.ContextExecutor) (NotificationTemplateSlice, error) {
	var o []*NotificationTemplate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NotificationTemplate slice")
	}

	return o, nil
}

// Count returns the count of all NotificationTemplate records in the query.
func (q notificationTemplateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count notification_templates rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q notificationTemplateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if notification_templates exists")
	}

	return count > 0, nil
}

// RoleNotificationTemplates retrieves all the role_notification_template's RoleNotificationTemplates with an executor.
func (o *NotificationTemplate) RoleNotificationTemplates(mods ...qm.QueryMod) roleNotificationTemplateQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"role_notification_templates\".\"notification_template_id\"=?", o.ID),
	)

	query := RoleNotificationTemplates(queryMods...)
	queries.SetFrom(query.Query, "\"role_notification_templates\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"role_notification_templates\".*"})
	}

	return query
}

// LoadRoleNotificationTemplates allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (notificationTemplateL) LoadRoleNotificationTemplates(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNotificationTemplate interface{}, mods queries.Applicator) error {
	var slice []*NotificationTemplate
	var object *NotificationTemplate

	if singular {
		object = maybeNotificationTemplate.(*NotificationTemplate)
	} else {
		slice = *maybeNotificationTemplate.(*[]*NotificationTemplate)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &notificationTemplateR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &notificationTemplateR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`role_notification_templates`),
		qm.WhereIn(`role_notification_templates.notification_template_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load role_notification_templates")
	}

	var resultSlice []*RoleNotificationTemplate
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice role_notification_templates")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on role_notification_templates")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for role_notification_templates")
	}

	if singular {
		object.R.RoleNotificationTemplates = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &roleNotificationTemplateR{}
			}
			foreign.R.NotificationTemplate = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.NotificationTemplateID {
				local.R.RoleNotificationTemplates = append(local.R.RoleNotificationTemplates, foreign)
				if foreign.R == nil {
					foreign.R = &roleNotificationTemplateR{}
				}
				foreign.R.NotificationTemplate = local
				break
			}
		}
	}

	return nil
}

// AddRoleNotificationTemplates adds the given related objects to the existing relationships
// of the notification_template, optionally inserting them as new records.
// Appends related to o.R.RoleNotificationTemplates.
// Sets related.R.NotificationTemplate appropriately.
func (o *NotificationTemplate) AddRoleNotificationTemplates(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RoleNotificationTemplate) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.NotificationTemplateID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"role_notification_templates\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"notification_template_id"}),
				strmangle.WhereClause("\"", "\"", 2, roleNotificationTemplatePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.RoleID, rel.NotificationTemplateID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.NotificationTemplateID = o.ID
		}
	}

	if o.R == nil {
		o.R = &notificationTemplateR{
			RoleNotificationTemplates: related,
		}
	} else {
		o.R.RoleNotificationTemplates = append(o.R.RoleNotificationTemplates, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &roleNotificationTemplateR{
				NotificationTemplate: o,
			}
		} else {
			rel.R.NotificationTemplate = o
		}
	}
	return nil
}

// NotificationTemplates retrieves all the records using an executor.
func NotificationTemplates(mods ...qm.QueryMod) notificationTemplateQuery {
	mods = append(mods, qm.From("\"notification_templates\""))
	return notificationTemplateQuery{NewQuery(mods...)}
}

// FindNotificationTemplate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNotificationTemplate(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*NotificationTemplate, error) {
	notificationTemplateObj := &NotificationTemplate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"notification_templates\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, notificationTemplateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from notification_templates")
	}

	return notificationTemplateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NotificationTemplate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no notification_templates provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(notificationTemplateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	notificationTemplateInsertCacheMut.RLock()
	cache, cached := notificationTemplateInsertCache[key]
	notificationTemplateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			notificationTemplateAllColumns,
			notificationTemplateColumnsWithDefault,
			notificationTemplateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(notificationTemplateType, notificationTemplateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(notificationTemplateType, notificationTemplateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"notification_templates\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"notification_templates\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into notification_templates")
	}

	if !cached {
		notificationTemplateInsertCacheMut.Lock()
		notificationTemplateInsertCache[key] = cache
		notificationTemplateInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the NotificationTemplate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NotificationTemplate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	notificationTemplateUpdateCacheMut.RLock()
	cache, cached := notificationTemplateUpdateCache[key]
	notificationTemplateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			notificationTemplateAllColumns,
			notificationTemplatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update notification_templates, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"notification_templates\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, notificationTemplatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(notificationTemplateType, notificationTemplateMapping, append(wl, notificationTemplatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update notification_templates row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for notification_templates")
	}

	if !cached {
		notificationTemplateUpdateCacheMut.Lock()
		notificationTemplateUpdateCache[key] = cache
		notificationTemplateUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q notificationTemplateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for notification_templates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for notification_templates")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NotificationTemplateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notificationTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"notification_templates\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, notificationTemplatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in notificationTemplate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all notificationTemplate")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NotificationTemplate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no notification_templates provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(notificationTemplateColumnsWithDefault, o)

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

	notificationTemplateUpsertCacheMut.RLock()
	cache, cached := notificationTemplateUpsertCache[key]
	notificationTemplateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			notificationTemplateAllColumns,
			notificationTemplateColumnsWithDefault,
			notificationTemplateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			notificationTemplateAllColumns,
			notificationTemplatePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert notification_templates, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(notificationTemplatePrimaryKeyColumns))
			copy(conflict, notificationTemplatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"notification_templates\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(notificationTemplateType, notificationTemplateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(notificationTemplateType, notificationTemplateMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert notification_templates")
	}

	if !cached {
		notificationTemplateUpsertCacheMut.Lock()
		notificationTemplateUpsertCache[key] = cache
		notificationTemplateUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single NotificationTemplate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NotificationTemplate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NotificationTemplate provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), notificationTemplatePrimaryKeyMapping)
	sql := "DELETE FROM \"notification_templates\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from notification_templates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for notification_templates")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q notificationTemplateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no notificationTemplateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from notification_templates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for notification_templates")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NotificationTemplateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notificationTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"notification_templates\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, notificationTemplatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from notificationTemplate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for notification_templates")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NotificationTemplate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNotificationTemplate(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NotificationTemplateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NotificationTemplateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notificationTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"notification_templates\".* FROM \"notification_templates\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, notificationTemplatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NotificationTemplateSlice")
	}

	*o = slice

	return nil
}

// NotificationTemplateExists checks if the NotificationTemplate row exists.
func NotificationTemplateExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"notification_templates\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if notification_templates exists")
	}

	return exists, nil
}
