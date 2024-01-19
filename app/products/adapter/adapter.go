package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

type ProductsAdapterInterface interface {
	//categories
	categoriesListRowGrpcFromSql(resp *db.CategoriesListRow) *rmsv1.CategoriesListRow
	CategoriesListGrpcFromSql(resp *[]db.CategoriesListRow) *rmsv1.CategoriesListResponse
	CategoryCreateSqlFromGrpc(req *rmsv1.CategoryCreateRequest) *db.CategoryCreateParams
	CategoryEntityGrpcFromSql(resp *db.ProductsSchemaCategory) *rmsv1.Category
	CategoryCreateGrpcFromSql(resp *db.ProductsSchemaCategory) *rmsv1.CategoryCreateResponse
	CategoryUpdateSqlFromGrpc(req *rmsv1.CategoryUpdateRequest) *db.CategoryUpdateParams
	CategoryUpdateGrpcFromSql(resp *db.ProductsSchemaCategory) *rmsv1.CategoryUpdateResponse
	CategoriesInputListGrpcFromSql(resp *[]db.CategoriesInputListRow) *rmsv1.CategoriesInputListResponse
	CategoryFindForUpdateGrpcFromSql(resp *db.CategoryFindForUpdateRow) *rmsv1.CategoryUpdateRequest

	// units
	unitsListRowGrpcFromSql(resp *db.UnitsListRow) *rmsv1.UnitsListRow
	UnitsListGrpcFromSql(resp *[]db.UnitsListRow) *rmsv1.UnitsListResponse
	UnitCreateSqlFromGrpc(req *rmsv1.UnitCreateRequest) *db.UnitCreateParams
	UnitEntityGrpcFromSql(resp *db.ProductsSchemaUnit) *rmsv1.Unit
	UnitCreateGrpcFromSql(resp *db.ProductsSchemaUnit) *rmsv1.UnitCreateResponse
	UnitUpdateSqlFromGrpc(req *rmsv1.UnitUpdateRequest) *db.UnitUpdateParams
	UnitUpdateGrpcFromSql(resp *db.ProductsSchemaUnit) *rmsv1.UnitUpdateResponse
	UnitsInputListGrpcFromSql(resp *[]db.UnitsInputListRow) *rmsv1.UnitsInputListResponse
	UnitFindForUpdateGrpcFromSql(resp *db.UnitFindForUpdateRow) *rmsv1.UnitUpdateRequest

	// products
	productsListRowGrpcFromSql(resp *db.ProductsListRow) *rmsv1.ProductsListRow
	ProductsListGrpcFromSql(resp *[]db.ProductsListRow) *rmsv1.ProductsListResponse
	ProductFindGrpcFromSql(resp db.ProductsSchemaProductsView) (*rmsv1.ProductsViewRow, error)
	ProductCreateSqlFromGrpc(req *rmsv1.ProductCreateRequest) *db.ProductCreateParams
	ProductEntityGrpcFromSql(resp *db.ProductsSchemaProduct) *rmsv1.Product
	ProductCreateGrpcFromSql(resp *db.ProductsSchemaProduct) *rmsv1.ProductCreateResponse
	ProductUpdateSqlFromGrpc(req *rmsv1.ProductUpdateRequest) *db.ProductUpdateParams
	ProductsInputListGrpcFromSql(resp *[]db.ProductsInputListRow) *rmsv1.ProductsInputListResponse
	ProductFindForUpdateGrpcFromSql(resp *db.ProductFindForUpdateRow) *rmsv1.ProductUpdateRequest
	ProductsListForTransactionGrpcFromSql(resp *[]db.ProductsListForTransactionRow) *rmsv1.ProductsListForTransactionResponse
	StockItemsListGrpcFromSql(resp *[]db.ProductsSchemaStockItemsView) *rmsv1.StockItemsListResponse

	// menu
	MenuFindGrpcFromSql(resp [][]byte) (*rmsv1.MenuFindResponse, error)

	//
	ingredientsListRowGrpcFromSql(resp *db.IngredientsListRow) (*rmsv1.IngredientsListRow, error)
	IngredientsListGrpcFromSql(resp *[]db.IngredientsListRow) (*rmsv1.IngredientsListResponse, error)
	IngredientCreateSqlFromGrpc(req *rmsv1.IngredientCreateRequest) *db.IngredientCreateParams
	IngredientUpdateSqlFromGrpc(req *rmsv1.IngredientUpdateRequest) *db.IngredientUpdateParams
	IngredientFindGrpcFromSql(resp *db.IngredientFindRow) (*rmsv1.IngredientFindResponse, error)
	IngredientsInputListGrpcFromSql(resp *[]db.IngredientsInputListRow) *rmsv1.IngredientsInputListResponse
	IngredientFindForUpdateGrpcFromSql(resp *db.IngredientFindForUpdateRow) *rmsv1.IngredientUpdateRequest

	// modifiers

	modifiersListRowGrpcFromSql(resp *db.ModifiersListRow) (*rmsv1.ModifiersListRow, error)
	ModifiersListGrpcFromSql(resp *[]db.ModifiersListRow) (*rmsv1.ModifiersListResponse, error)
	ModifierCreateSqlFromGrpc(req *rmsv1.ModifierCreateRequest) *db.ModifierCreateParams
	ModifierUpdateSqlFromGrpc(req *rmsv1.ModifierUpdateRequest) *db.ModifierUpdateParams
	ModifierFindGrpcFromSql(resp *db.ProductsSchemaModifiersView) (*rmsv1.ModifierFindResponse, error)
	ModifiersInputListGrpcFromSql(resp *[]db.ModifiersInputListRow) *rmsv1.ModifiersInputListResponse
}

type ProductsAdapter struct {
	dateFormat string
}

func NewProductsAdapter() ProductsAdapterInterface {
	return &ProductsAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
