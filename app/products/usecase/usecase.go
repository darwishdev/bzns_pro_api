package service

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/products/adapter"
	"github.com/darwishdev/bzns_pro_api/app/products/repo"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

type ProductsUsecaseInterface interface {
	//categories
	CategoryCreate(ctx context.Context, req *rmsv1.CategoryCreateRequest) (*rmsv1.CategoryCreateResponse, error)
	CategoryFindForUpdate(ctx context.Context, req *rmsv1.CategoryFindForUpdateRequest) (*rmsv1.CategoryUpdateRequest, error)
	CategoryUpdate(ctx context.Context, req *rmsv1.CategoryUpdateRequest) (*rmsv1.CategoryUpdateResponse, error)
	CategoriesList(ctx context.Context, req *rmsv1.CategoriesListRequest) (*rmsv1.CategoriesListResponse, error)
	CategoryDeleteRestore(ctx context.Context, req *rmsv1.CategoryDeleteRestoreRequest) (*rmsv1.CategoryDeleteRestoreResponse, error)
	CategoriesInputList(ctx context.Context, req *rmsv1.CategoriesInputListRequest) (*rmsv1.CategoriesInputListResponse, error)

	//units
	UnitCreate(ctx context.Context, req *rmsv1.UnitCreateRequest) (*rmsv1.UnitCreateResponse, error)
	UnitFindForUpdate(ctx context.Context, req *rmsv1.UnitFindForUpdateRequest) (*rmsv1.UnitUpdateRequest, error)
	UnitUpdate(ctx context.Context, req *rmsv1.UnitUpdateRequest) (*rmsv1.UnitUpdateResponse, error)
	UnitsList(ctx context.Context, req *rmsv1.UnitsListRequest) (*rmsv1.UnitsListResponse, error)
	UnitDeleteRestore(ctx context.Context, req *rmsv1.UnitDeleteRestoreRequest) (*rmsv1.UnitDeleteRestoreResponse, error)
	UnitsInputList(ctx context.Context, req *rmsv1.UnitsInputListRequest) (*rmsv1.UnitsInputListResponse, error)

	// products
	ProductCreate(ctx context.Context, req *rmsv1.ProductCreateRequest) (*rmsv1.ProductCreateResponse, error)
	ProductFindForUpdate(ctx context.Context, req *rmsv1.ProductFindForUpdateRequest) (*rmsv1.ProductUpdateRequest, error)
	ProductUpdate(ctx context.Context, req *rmsv1.ProductUpdateRequest) (*rmsv1.ProductUpdateResponse, error)
	ProductsList(ctx context.Context, req *rmsv1.ProductsListRequest) (*rmsv1.ProductsListResponse, error)
	ProductDeleteRestore(ctx context.Context, req *rmsv1.ProductDeleteRestoreRequest) (*rmsv1.ProductDeleteRestoreResponse, error)
	ProductsInputList(ctx context.Context, req *rmsv1.ProductsInputListRequest) (*rmsv1.ProductsInputListResponse, error)
	ProductsListForTransaction(ctx context.Context, req *rmsv1.ProductsListForTransactionRequest) (*rmsv1.ProductsListForTransactionResponse, error)
	StockItemsList(ctx context.Context, req *rmsv1.StockItemsListRequest) (*rmsv1.StockItemsListResponse, error)
	ProductFind(ctx context.Context, req *rmsv1.ProductFindRequest) (*rmsv1.ProductFindResponse, error)

	//ingredients
	IngredientsList(ctx context.Context, req *rmsv1.IngredientsListRequest) (*rmsv1.IngredientsListResponse, error)
	IngredientCreate(ctx context.Context, req *rmsv1.IngredientCreateRequest) (*rmsv1.IngredientCreateResponse, error)
	IngredientUpdate(ctx context.Context, req *rmsv1.IngredientUpdateRequest) (*rmsv1.IngredientUpdateResponse, error)
	IngredientDeleteRestore(ctx context.Context, req *rmsv1.IngredientDeleteRestoreRequest) (*rmsv1.IngredientDeleteRestoreResponse, error)
	IngredientFind(ctx context.Context, req *rmsv1.IngredientFindRequest) (*rmsv1.IngredientFindResponse, error)
	IngredientFindForUpdate(ctx context.Context, req *rmsv1.IngredientFindForUpdateRequest) (*rmsv1.IngredientUpdateRequest, error)
	IngredientsInputList(ctx context.Context, req *rmsv1.IngredientsInputListRequest) (*rmsv1.IngredientsInputListResponse, error)
	// menu
	MenuFind(ctx context.Context, req *rmsv1.MenuFindRequest) (*rmsv1.MenuFindResponse, error)

	// modifier
	ModifiersList(ctx context.Context, req *rmsv1.ModifiersListRequest) (*rmsv1.ModifiersListResponse, error)
	ModifiersInputList(ctx context.Context, req *rmsv1.ModifiersInputListRequest) (*rmsv1.ModifiersInputListResponse, error)
	ModifierCreate(ctx context.Context, req *rmsv1.ModifierCreateRequest) (*rmsv1.ModifierCreateResponse, error)
	ModifierUpdate(ctx context.Context, req *rmsv1.ModifierUpdateRequest) (*rmsv1.ModifierUpdateResponse, error)
	ModifierDeleteRestore(ctx context.Context, req *rmsv1.ModifierDeleteRestoreRequest) (*rmsv1.ModifierDeleteRestoreResponse, error)
	ModifierFind(ctx context.Context, req *rmsv1.ModifierFindRequest) (*rmsv1.ModifierFindResponse, error)
}

type ProductsUsecase struct {
	repo      repo.ProductsRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.ProductsAdapterInterface
}

func NewProductsUsecase(store db.Store, validator *protovalidate.Validator) ProductsUsecaseInterface {
	repo := repo.NewProductsRepo(store)
	adapter := adapter.NewProductsAdapter()

	return &ProductsUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
