// Code generated by SQLBoiler 4.0.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testApplicationStateTransitions(t *testing.T) {
	t.Parallel()

	query := ApplicationStateTransitions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testApplicationStateTransitionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
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

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApplicationStateTransitionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ApplicationStateTransitions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApplicationStateTransitionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ApplicationStateTransitionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApplicationStateTransitionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ApplicationStateTransitionExists(ctx, tx, o.ApplicantID, o.FromApplicationStateID, o.ToApplicationStateID)
	if err != nil {
		t.Errorf("Unable to check if ApplicationStateTransition exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ApplicationStateTransitionExists to return true, but got false.")
	}
}

func testApplicationStateTransitionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	applicationStateTransitionFound, err := FindApplicationStateTransition(ctx, tx, o.ApplicantID, o.FromApplicationStateID, o.ToApplicationStateID)
	if err != nil {
		t.Error(err)
	}

	if applicationStateTransitionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testApplicationStateTransitionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ApplicationStateTransitions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testApplicationStateTransitionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ApplicationStateTransitions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testApplicationStateTransitionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	applicationStateTransitionOne := &ApplicationStateTransition{}
	applicationStateTransitionTwo := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, applicationStateTransitionOne, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}
	if err = randomize.Struct(seed, applicationStateTransitionTwo, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = applicationStateTransitionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = applicationStateTransitionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ApplicationStateTransitions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testApplicationStateTransitionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	applicationStateTransitionOne := &ApplicationStateTransition{}
	applicationStateTransitionTwo := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, applicationStateTransitionOne, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}
	if err = randomize.Struct(seed, applicationStateTransitionTwo, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = applicationStateTransitionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = applicationStateTransitionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testApplicationStateTransitionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testApplicationStateTransitionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(applicationStateTransitionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testApplicationStateTransitionToOneApplicantUsingApplicant(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ApplicationStateTransition
	var foreign Applicant

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, applicantDBTypes, false, applicantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Applicant struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ApplicantID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Applicant().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ApplicationStateTransitionSlice{&local}
	if err = local.L.LoadApplicant(ctx, tx, false, (*[]*ApplicationStateTransition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Applicant == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Applicant = nil
	if err = local.L.LoadApplicant(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Applicant == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testApplicationStateTransitionToOneApplicationStateUsingFromApplicationState(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ApplicationStateTransition
	var foreign ApplicationState

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FromApplicationStateID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FromApplicationState().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ApplicationStateTransitionSlice{&local}
	if err = local.L.LoadFromApplicationState(ctx, tx, false, (*[]*ApplicationStateTransition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromApplicationState == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FromApplicationState = nil
	if err = local.L.LoadFromApplicationState(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromApplicationState == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testApplicationStateTransitionToOneApplicationStateUsingToApplicationState(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ApplicationStateTransition
	var foreign ApplicationState

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ToApplicationStateID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ToApplicationState().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ApplicationStateTransitionSlice{&local}
	if err = local.L.LoadToApplicationState(ctx, tx, false, (*[]*ApplicationStateTransition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ToApplicationState == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ToApplicationState = nil
	if err = local.L.LoadToApplicationState(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ToApplicationState == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testApplicationStateTransitionToOneSetOpApplicantUsingApplicant(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationStateTransition
	var b, c Applicant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateTransitionDBTypes, false, strmangle.SetComplement(applicationStateTransitionPrimaryKeyColumns, applicationStateTransitionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, applicantDBTypes, false, strmangle.SetComplement(applicantPrimaryKeyColumns, applicantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicantDBTypes, false, strmangle.SetComplement(applicantPrimaryKeyColumns, applicantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Applicant{&b, &c} {
		err = a.SetApplicant(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Applicant != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ApplicationStateTransitions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ApplicantID != x.ID {
			t.Error("foreign key was wrong value", a.ApplicantID)
		}

		if exists, err := ApplicationStateTransitionExists(ctx, tx, a.ApplicantID, a.FromApplicationStateID, a.ToApplicationStateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testApplicationStateTransitionToOneSetOpApplicationStateUsingFromApplicationState(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationStateTransition
	var b, c ApplicationState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateTransitionDBTypes, false, strmangle.SetComplement(applicationStateTransitionPrimaryKeyColumns, applicationStateTransitionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ApplicationState{&b, &c} {
		err = a.SetFromApplicationState(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FromApplicationState != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FromApplicationStateApplicationStateTransitions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FromApplicationStateID != x.ID {
			t.Error("foreign key was wrong value", a.FromApplicationStateID)
		}

		if exists, err := ApplicationStateTransitionExists(ctx, tx, a.ApplicantID, a.FromApplicationStateID, a.ToApplicationStateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testApplicationStateTransitionToOneSetOpApplicationStateUsingToApplicationState(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationStateTransition
	var b, c ApplicationState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateTransitionDBTypes, false, strmangle.SetComplement(applicationStateTransitionPrimaryKeyColumns, applicationStateTransitionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ApplicationState{&b, &c} {
		err = a.SetToApplicationState(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ToApplicationState != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ToApplicationStateApplicationStateTransitions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ToApplicationStateID != x.ID {
			t.Error("foreign key was wrong value", a.ToApplicationStateID)
		}

		if exists, err := ApplicationStateTransitionExists(ctx, tx, a.ApplicantID, a.FromApplicationStateID, a.ToApplicationStateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testApplicationStateTransitionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
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

func testApplicationStateTransitionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ApplicationStateTransitionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testApplicationStateTransitionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ApplicationStateTransitions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	applicationStateTransitionDBTypes = map[string]string{`ApplicantID`: `uuid`, `FromApplicationStateID`: `uuid`, `ToApplicationStateID`: `uuid`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                                 = bytes.MinRead
)

func testApplicationStateTransitionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(applicationStateTransitionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(applicationStateTransitionAllColumns) == len(applicationStateTransitionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testApplicationStateTransitionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(applicationStateTransitionAllColumns) == len(applicationStateTransitionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationStateTransition{}
	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, applicationStateTransitionDBTypes, true, applicationStateTransitionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(applicationStateTransitionAllColumns, applicationStateTransitionPrimaryKeyColumns) {
		fields = applicationStateTransitionAllColumns
	} else {
		fields = strmangle.SetComplement(
			applicationStateTransitionAllColumns,
			applicationStateTransitionPrimaryKeyColumns,
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

	slice := ApplicationStateTransitionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testApplicationStateTransitionsUpsert(t *testing.T) {
	t.Parallel()

	if len(applicationStateTransitionAllColumns) == len(applicationStateTransitionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ApplicationStateTransition{}
	if err = randomize.Struct(seed, &o, applicationStateTransitionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ApplicationStateTransition: %s", err)
	}

	count, err := ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, applicationStateTransitionDBTypes, false, applicationStateTransitionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ApplicationStateTransition struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ApplicationStateTransition: %s", err)
	}

	count, err = ApplicationStateTransitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
