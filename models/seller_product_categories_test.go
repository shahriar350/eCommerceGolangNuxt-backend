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

func testSellerProductCategories(t *testing.T) {
	t.Parallel()

	query := SellerProductCategories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSellerProductCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
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

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSellerProductCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SellerProductCategories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSellerProductCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SellerProductCategorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSellerProductCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SellerProductCategoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SellerProductCategory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SellerProductCategoryExists to return true, but got false.")
	}
}

func testSellerProductCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	sellerProductCategoryFound, err := FindSellerProductCategory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if sellerProductCategoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSellerProductCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SellerProductCategories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSellerProductCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SellerProductCategories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSellerProductCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	sellerProductCategoryOne := &SellerProductCategory{}
	sellerProductCategoryTwo := &SellerProductCategory{}
	if err = randomize.Struct(seed, sellerProductCategoryOne, sellerProductCategoryDBTypes, false, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, sellerProductCategoryTwo, sellerProductCategoryDBTypes, false, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sellerProductCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sellerProductCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SellerProductCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSellerProductCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	sellerProductCategoryOne := &SellerProductCategory{}
	sellerProductCategoryTwo := &SellerProductCategory{}
	if err = randomize.Struct(seed, sellerProductCategoryOne, sellerProductCategoryDBTypes, false, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, sellerProductCategoryTwo, sellerProductCategoryDBTypes, false, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sellerProductCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sellerProductCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func sellerProductCategoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func sellerProductCategoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SellerProductCategory) error {
	*o = SellerProductCategory{}
	return nil
}

func testSellerProductCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SellerProductCategory{}
	o := &SellerProductCategory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory object: %s", err)
	}

	AddSellerProductCategoryHook(boil.BeforeInsertHook, sellerProductCategoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryBeforeInsertHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.AfterInsertHook, sellerProductCategoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryAfterInsertHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.AfterSelectHook, sellerProductCategoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryAfterSelectHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.BeforeUpdateHook, sellerProductCategoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryBeforeUpdateHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.AfterUpdateHook, sellerProductCategoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryAfterUpdateHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.BeforeDeleteHook, sellerProductCategoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryBeforeDeleteHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.AfterDeleteHook, sellerProductCategoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryAfterDeleteHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.BeforeUpsertHook, sellerProductCategoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryBeforeUpsertHooks = []SellerProductCategoryHook{}

	AddSellerProductCategoryHook(boil.AfterUpsertHook, sellerProductCategoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	sellerProductCategoryAfterUpsertHooks = []SellerProductCategoryHook{}
}

func testSellerProductCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSellerProductCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(sellerProductCategoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSellerProductCategoryToOneSellerProductUsingSellerProduct(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SellerProductCategory
	var foreign SellerProduct

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, sellerProductDBTypes, false, sellerProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProduct struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.SellerProductID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SellerProduct().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SellerProductCategorySlice{&local}
	if err = local.L.LoadSellerProduct(ctx, tx, false, (*[]*SellerProductCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SellerProduct == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SellerProduct = nil
	if err = local.L.LoadSellerProduct(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SellerProduct == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSellerProductCategoryToOneSetOpSellerProductUsingSellerProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SellerProductCategory
	var b, c SellerProduct

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sellerProductCategoryDBTypes, false, strmangle.SetComplement(sellerProductCategoryPrimaryKeyColumns, sellerProductCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, sellerProductDBTypes, false, strmangle.SetComplement(sellerProductPrimaryKeyColumns, sellerProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sellerProductDBTypes, false, strmangle.SetComplement(sellerProductPrimaryKeyColumns, sellerProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SellerProduct{&b, &c} {
		err = a.SetSellerProduct(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SellerProduct != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SellerProductCategories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.SellerProductID, x.ID) {
			t.Error("foreign key was wrong value", a.SellerProductID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SellerProductID))
		reflect.Indirect(reflect.ValueOf(&a.SellerProductID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.SellerProductID, x.ID) {
			t.Error("foreign key was wrong value", a.SellerProductID, x.ID)
		}
	}
}

func testSellerProductCategoryToOneRemoveOpSellerProductUsingSellerProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SellerProductCategory
	var b SellerProduct

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sellerProductCategoryDBTypes, false, strmangle.SetComplement(sellerProductCategoryPrimaryKeyColumns, sellerProductCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, sellerProductDBTypes, false, strmangle.SetComplement(sellerProductPrimaryKeyColumns, sellerProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetSellerProduct(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveSellerProduct(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.SellerProduct().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.SellerProduct != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.SellerProductID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.SellerProductCategories) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testSellerProductCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
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

func testSellerProductCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SellerProductCategorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSellerProductCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SellerProductCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	sellerProductCategoryDBTypes = map[string]string{`ID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `SellerProductID`: `bigint`, `CategoryID`: `bigint`}
	_                            = bytes.MinRead
)

func testSellerProductCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(sellerProductCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(sellerProductCategoryAllColumns) == len(sellerProductCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSellerProductCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(sellerProductCategoryAllColumns) == len(sellerProductCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SellerProductCategory{}
	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sellerProductCategoryDBTypes, true, sellerProductCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(sellerProductCategoryAllColumns, sellerProductCategoryPrimaryKeyColumns) {
		fields = sellerProductCategoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			sellerProductCategoryAllColumns,
			sellerProductCategoryPrimaryKeyColumns,
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

	slice := SellerProductCategorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSellerProductCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(sellerProductCategoryAllColumns) == len(sellerProductCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SellerProductCategory{}
	if err = randomize.Struct(seed, &o, sellerProductCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SellerProductCategory: %s", err)
	}

	count, err := SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, sellerProductCategoryDBTypes, false, sellerProductCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SellerProductCategory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SellerProductCategory: %s", err)
	}

	count, err = SellerProductCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
