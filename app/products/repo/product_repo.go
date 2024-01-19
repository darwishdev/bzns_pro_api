package repo

import (
	"context"

	db "github.com/meloneg/mln_rms_core/common/db/gen"
)

func (repo *ProductsRepo) ProductFind(ctx context.Context, id int32) (*db.ProductsSchemaProductsView, error) {
	resp, err := repo.store.ProductFind(context.Background(), id)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *ProductsRepo) ProductsList(ctx context.Context) (*[]db.ProductsListRow, error) {
	resp, err := repo.store.ProductsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *ProductsRepo) ProductsListForTransaction(ctx context.Context) (*[]db.ProductsListForTransactionRow, error) {
	resp, err := repo.store.ProductsListForTransaction(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *ProductsRepo) StockItemsList(ctx context.Context) (*[]db.ProductsSchemaStockItemsView, error) {
	resp, err := repo.store.StockItemsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *ProductsRepo) ProductsInputList(ctx context.Context) (*[]db.ProductsInputListRow, error) {
	resp, err := repo.store.ProductsInputList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *ProductsRepo) ProductDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.ProductDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}
func (repo *ProductsRepo) ProductCreate(ctx context.Context, req *db.ProductCreateParams) error {
	err := repo.store.ProductCreate(context.Background(), *req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *ProductsRepo) ProductFindForUpdate(ctx context.Context, req *int32) (*db.ProductFindForUpdateRow, error) {
	resp, err := repo.store.ProductFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *ProductsRepo) ProductUpdate(ctx context.Context, req *db.ProductUpdateParams) error {
	err := repo.store.ProductUpdate(context.Background(), *req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
