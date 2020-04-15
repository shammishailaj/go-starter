// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testNotificationTemplates(t *testing.T) {
	t.Parallel()

	query := NotificationTemplates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNotificationTemplatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
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

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationTemplatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NotificationTemplates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationTemplatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationTemplateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationTemplatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NotificationTemplateExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if NotificationTemplate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NotificationTemplateExists to return true, but got false.")
	}
}

func testNotificationTemplatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	notificationTemplateFound, err := FindNotificationTemplate(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if notificationTemplateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNotificationTemplatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NotificationTemplates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNotificationTemplatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NotificationTemplates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNotificationTemplatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	notificationTemplateOne := &NotificationTemplate{}
	notificationTemplateTwo := &NotificationTemplate{}
	if err = randomize.Struct(seed, notificationTemplateOne, notificationTemplateDBTypes, false, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}
	if err = randomize.Struct(seed, notificationTemplateTwo, notificationTemplateDBTypes, false, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = notificationTemplateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = notificationTemplateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NotificationTemplates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNotificationTemplatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notificationTemplateOne := &NotificationTemplate{}
	notificationTemplateTwo := &NotificationTemplate{}
	if err = randomize.Struct(seed, notificationTemplateOne, notificationTemplateDBTypes, false, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}
	if err = randomize.Struct(seed, notificationTemplateTwo, notificationTemplateDBTypes, false, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = notificationTemplateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = notificationTemplateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testNotificationTemplatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotificationTemplatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(notificationTemplateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotificationTemplateToManyRoleNotificationTemplates(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NotificationTemplate
	var b, c RoleNotificationTemplate

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleNotificationTemplateDBTypes, false, roleNotificationTemplateColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.NotificationTemplateID = a.ID
	c.NotificationTemplateID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RoleNotificationTemplates().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.NotificationTemplateID == b.NotificationTemplateID {
			bFound = true
		}
		if v.NotificationTemplateID == c.NotificationTemplateID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := NotificationTemplateSlice{&a}
	if err = a.L.LoadRoleNotificationTemplates(ctx, tx, false, (*[]*NotificationTemplate)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleNotificationTemplates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RoleNotificationTemplates = nil
	if err = a.L.LoadRoleNotificationTemplates(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleNotificationTemplates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testNotificationTemplateToManyAddOpRoleNotificationTemplates(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NotificationTemplate
	var b, c, d, e RoleNotificationTemplate

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, notificationTemplateDBTypes, false, strmangle.SetComplement(notificationTemplatePrimaryKeyColumns, notificationTemplateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RoleNotificationTemplate{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleNotificationTemplateDBTypes, false, strmangle.SetComplement(roleNotificationTemplatePrimaryKeyColumns, roleNotificationTemplateColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*RoleNotificationTemplate{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRoleNotificationTemplates(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.NotificationTemplateID {
			t.Error("foreign key was wrong value", a.ID, first.NotificationTemplateID)
		}
		if a.ID != second.NotificationTemplateID {
			t.Error("foreign key was wrong value", a.ID, second.NotificationTemplateID)
		}

		if first.R.NotificationTemplate != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.NotificationTemplate != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RoleNotificationTemplates[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RoleNotificationTemplates[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RoleNotificationTemplates().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testNotificationTemplatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
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

func testNotificationTemplatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationTemplateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNotificationTemplatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NotificationTemplates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	notificationTemplateDBTypes = map[string]string{`ID`: `uuid`, `Text`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                           = bytes.MinRead
)

func testNotificationTemplatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(notificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(notificationTemplateAllColumns) == len(notificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNotificationTemplatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(notificationTemplateAllColumns) == len(notificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NotificationTemplate{}
	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, notificationTemplateDBTypes, true, notificationTemplatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(notificationTemplateAllColumns, notificationTemplatePrimaryKeyColumns) {
		fields = notificationTemplateAllColumns
	} else {
		fields = strmangle.SetComplement(
			notificationTemplateAllColumns,
			notificationTemplatePrimaryKeyColumns,
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

	slice := NotificationTemplateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNotificationTemplatesUpsert(t *testing.T) {
	t.Parallel()

	if len(notificationTemplateAllColumns) == len(notificationTemplatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NotificationTemplate{}
	if err = randomize.Struct(seed, &o, notificationTemplateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NotificationTemplate: %s", err)
	}

	count, err := NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, notificationTemplateDBTypes, false, notificationTemplatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NotificationTemplate struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NotificationTemplate: %s", err)
	}

	count, err = NotificationTemplates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}