// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testUserLocations(t *testing.T) {
	t.Parallel()

	query := UserLocations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserLocationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
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

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserLocationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserLocations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserLocationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserLocationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserLocationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserLocationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserLocation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserLocationExists to return true, but got false.")
	}
}

func testUserLocationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userLocationFound, err := FindUserLocation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userLocationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserLocationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserLocations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserLocationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserLocations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserLocationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userLocationOne := &UserLocation{}
	userLocationTwo := &UserLocation{}
	if err = randomize.Struct(seed, userLocationOne, userLocationDBTypes, false, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}
	if err = randomize.Struct(seed, userLocationTwo, userLocationDBTypes, false, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userLocationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userLocationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserLocations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserLocationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userLocationOne := &UserLocation{}
	userLocationTwo := &UserLocation{}
	if err = randomize.Struct(seed, userLocationOne, userLocationDBTypes, false, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}
	if err = randomize.Struct(seed, userLocationTwo, userLocationDBTypes, false, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userLocationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userLocationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userLocationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func userLocationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserLocation) error {
	*o = UserLocation{}
	return nil
}

func testUserLocationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserLocation{}
	o := &UserLocation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userLocationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserLocation object: %s", err)
	}

	AddUserLocationHook(boil.BeforeInsertHook, userLocationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userLocationBeforeInsertHooks = []UserLocationHook{}

	AddUserLocationHook(boil.AfterInsertHook, userLocationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userLocationAfterInsertHooks = []UserLocationHook{}

	AddUserLocationHook(boil.AfterSelectHook, userLocationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userLocationAfterSelectHooks = []UserLocationHook{}

	AddUserLocationHook(boil.BeforeUpdateHook, userLocationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userLocationBeforeUpdateHooks = []UserLocationHook{}

	AddUserLocationHook(boil.AfterUpdateHook, userLocationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userLocationAfterUpdateHooks = []UserLocationHook{}

	AddUserLocationHook(boil.BeforeDeleteHook, userLocationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userLocationBeforeDeleteHooks = []UserLocationHook{}

	AddUserLocationHook(boil.AfterDeleteHook, userLocationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userLocationAfterDeleteHooks = []UserLocationHook{}

	AddUserLocationHook(boil.BeforeUpsertHook, userLocationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userLocationBeforeUpsertHooks = []UserLocationHook{}

	AddUserLocationHook(boil.AfterUpsertHook, userLocationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userLocationAfterUpsertHooks = []UserLocationHook{}
}

func testUserLocationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserLocationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userLocationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserLocationToManyCheckouts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserLocation
	var b, c Checkout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, checkoutDBTypes, false, checkoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, checkoutDBTypes, false, checkoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.UserLocationID, a.ID)
	queries.Assign(&c.UserLocationID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Checkouts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.UserLocationID, b.UserLocationID) {
			bFound = true
		}
		if queries.Equal(v.UserLocationID, c.UserLocationID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UserLocationSlice{&a}
	if err = a.L.LoadCheckouts(ctx, tx, false, (*[]*UserLocation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Checkouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Checkouts = nil
	if err = a.L.LoadCheckouts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Checkouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testUserLocationToManyAddOpCheckouts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserLocation
	var b, c, d, e Checkout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userLocationDBTypes, false, strmangle.SetComplement(userLocationPrimaryKeyColumns, userLocationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Checkout{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, checkoutDBTypes, false, strmangle.SetComplement(checkoutPrimaryKeyColumns, checkoutColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Checkout{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCheckouts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.UserLocationID) {
			t.Error("foreign key was wrong value", a.ID, first.UserLocationID)
		}
		if !queries.Equal(a.ID, second.UserLocationID) {
			t.Error("foreign key was wrong value", a.ID, second.UserLocationID)
		}

		if first.R.UserLocation != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.UserLocation != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Checkouts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Checkouts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Checkouts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testUserLocationToManySetOpCheckouts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserLocation
	var b, c, d, e Checkout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userLocationDBTypes, false, strmangle.SetComplement(userLocationPrimaryKeyColumns, userLocationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Checkout{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, checkoutDBTypes, false, strmangle.SetComplement(checkoutPrimaryKeyColumns, checkoutColumnsWithoutDefault)...); err != nil {
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

	err = a.SetCheckouts(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Checkouts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCheckouts(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Checkouts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserLocationID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserLocationID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.UserLocationID) {
		t.Error("foreign key was wrong value", a.ID, d.UserLocationID)
	}
	if !queries.Equal(a.ID, e.UserLocationID) {
		t.Error("foreign key was wrong value", a.ID, e.UserLocationID)
	}

	if b.R.UserLocation != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.UserLocation != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.UserLocation != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.UserLocation != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Checkouts[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Checkouts[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testUserLocationToManyRemoveOpCheckouts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserLocation
	var b, c, d, e Checkout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userLocationDBTypes, false, strmangle.SetComplement(userLocationPrimaryKeyColumns, userLocationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Checkout{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, checkoutDBTypes, false, strmangle.SetComplement(checkoutPrimaryKeyColumns, checkoutColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCheckouts(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Checkouts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCheckouts(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Checkouts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserLocationID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserLocationID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.UserLocation != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.UserLocation != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.UserLocation != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.UserLocation != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Checkouts) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Checkouts[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Checkouts[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testUserLocationToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserLocation
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserLocationSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserLocation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserLocationToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserLocation
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userLocationDBTypes, false, strmangle.SetComplement(userLocationPrimaryKeyColumns, userLocationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserLocations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testUserLocationToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserLocation
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userLocationDBTypes, false, strmangle.SetComplement(userLocationPrimaryKeyColumns, userLocationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.UserID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.UserLocations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testUserLocationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
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

func testUserLocationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserLocationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserLocationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserLocations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userLocationDBTypes = map[string]string{`ID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `UserID`: `bigint`, `Area`: `text`, `Street`: `text`, `House`: `text`, `PostOffice`: `text`, `PostCode`: `text`, `PoliceStation`: `text`, `City`: `text`}
	_                   = bytes.MinRead
)

func testUserLocationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userLocationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userLocationAllColumns) == len(userLocationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserLocationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userLocationAllColumns) == len(userLocationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserLocation{}
	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userLocationDBTypes, true, userLocationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userLocationAllColumns, userLocationPrimaryKeyColumns) {
		fields = userLocationAllColumns
	} else {
		fields = strmangle.SetComplement(
			userLocationAllColumns,
			userLocationPrimaryKeyColumns,
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

	slice := UserLocationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserLocationsUpsert(t *testing.T) {
	t.Parallel()

	if len(userLocationAllColumns) == len(userLocationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserLocation{}
	if err = randomize.Struct(seed, &o, userLocationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserLocation: %s", err)
	}

	count, err := UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userLocationDBTypes, false, userLocationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserLocation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserLocation: %s", err)
	}

	count, err = UserLocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
