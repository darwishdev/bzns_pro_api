package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

type ProductsRepoInterface interface {
	//categories
	CategoriesList(ctx context.Context) (*[]db.CategoriesListRow, error)
	CategoryDeleteRestore(ctx context.Context, req []int32) error
	CategoryCreate(ctx context.Context, req *db.CategoryCreateParams) (*db.ProductsSchemaCategory, error)
	CategoryFindForUpdate(ctx context.Context, req *int32) (*db.CategoryFindForUpdateRow, error)
	CategoryUpdate(ctx context.Context, req *db.CategoryUpdateParams) (*db.ProductsSchemaCategory, error)
	CategoriesInputList(ctx context.Context) (*[]db.CategoriesInputListRow, error)

	// units
	UnitsList(ctx context.Context) (*[]db.UnitsListRow, error)
	UnitsInputList(ctx context.Context) (*[]db.UnitsInputListRow, error)
	UnitDeleteRestore(ctx context.Context, req []int32) error
	UnitCreate(ctx context.Context, req *db.UnitCreateParams) (*db.ProductsSchemaUnit, error)
	UnitFindForUpdate(ctx context.Context, req *int32) (*db.UnitFindForUpdateRow, error)
	UnitUpdate(ctx context.Context, req *db.UnitUpdateParams) (*db.ProductsSchemaUnit, error)

	//products
	ProductsList(ctx context.Context) (*[]db.ProductsListRow, error)
	ProductsInputList(ctx context.Context) (*[]db.ProductsInputListRow, error)
	ProductDeleteRestore(ctx context.Context, req []int32) error
	ProductCreate(ctx context.Context, req *db.ProductCreateParams) error
	ProductFindForUpdate(ctx context.Context, req *int32) (*db.ProductFindForUpdateRow, error)
	ProductUpdate(ctx context.Context, req *db.ProductUpdateParams) error
	ProductsListForTransaction(ctx context.Context) (*[]db.ProductsListForTransactionRow, error)
	ProductFind(ctx context.Context, id int32) (*db.ProductsSchemaProductsView, error)
	StockItemsList(ctx context.Context) (*[]db.ProductsSchemaStockItemsView, error)

	// menu
	MenuFindNested(ctx context.Context) ([][]byte, error)

	// ingredients
	IngredientsList(ctx context.Context) ([]db.IngredientsListRow, error)
	IngredientCreate(ctx context.Context, req db.IngredientCreateParams) error
	IngredientUpdate(ctx context.Context, req db.IngredientUpdateParams) error
	IngredientDeleteRestore(ctx context.Context, req []int32) error
	IngredientFind(ctx context.Context, req int32) (*db.IngredientFindRow, error)
	IngredientFindForUpdate(ctx context.Context, req int32) (*db.IngredientFindForUpdateRow, error)
	IngredientsInputList(ctx context.Context) ([]db.IngredientsInputListRow, error)

	// modifiers
	ModifiersList(ctx context.Context) ([]db.ModifiersListRow, error)
	ModifiersInputList(ctx context.Context) ([]db.ModifiersInputListRow, error)
	ModifierCreate(ctx context.Context, req db.ModifierCreateParams) error
	ModifierUpdate(ctx context.Context, req db.ModifierUpdateParams) error
	ModifierDeleteRestore(ctx context.Context, req []int32) error
	ModifierFind(ctx context.Context, req int32) (*db.ProductsSchemaModifiersView, error)
	// ModifierFindForUpdate(ctx context.Context, req int32) (*db.ModifierFindForUpdateRow, error)
}

type ProductsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewProductsRepo(store db.Store) *ProductsRepo {
	errorHandler := map[string]string{
		"categories_category_name_key":  "categoryName",
		"modifiers_modifier_name_key":   "modifierName",
		"branches_branch_name_key":      "branchName",
		"warehouses_warehouse_name_key": "warehouseName",
	}
	return &ProductsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
