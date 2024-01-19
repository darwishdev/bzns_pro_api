package service

import (
	"context"

	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (u *ProductsUsecase) IngredientsList(ctx context.Context, req *rmsv1.IngredientsListRequest) (*rmsv1.IngredientsListResponse, error) {
	record, err := u.repo.IngredientsList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := u.adapter.IngredientsListGrpcFromSql(&record)
	if err != nil {
		return nil, err
	}

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("ingredients")

	return resp, nil
}
func (u *ProductsUsecase) IngredientsInputList(ctx context.Context, req *rmsv1.IngredientsInputListRequest) (*rmsv1.IngredientsInputListResponse, error) {
	record, err := u.repo.IngredientsInputList(ctx)
	if err != nil {
		return nil, err
	}
	resp := u.adapter.IngredientsInputListGrpcFromSql(&record)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *ProductsUsecase) IngredientCreate(ctx context.Context, req *rmsv1.IngredientCreateRequest) (*rmsv1.IngredientCreateResponse, error) {
	param := u.adapter.IngredientCreateSqlFromGrpc(req)

	err := u.repo.IngredientCreate(ctx, *param)
	if err != nil {
		return nil, err
	}

	return &rmsv1.IngredientCreateResponse{}, nil
}

func (u *ProductsUsecase) IngredientUpdate(ctx context.Context, req *rmsv1.IngredientUpdateRequest) (*rmsv1.IngredientUpdateResponse, error) {
	param := u.adapter.IngredientUpdateSqlFromGrpc(req)

	err := u.repo.IngredientUpdate(ctx, *param)
	if err != nil {
		return nil, err
	}

	return &rmsv1.IngredientUpdateResponse{}, nil
}
func (u *ProductsUsecase) IngredientDeleteRestore(ctx context.Context, req *rmsv1.IngredientDeleteRestoreRequest) (*rmsv1.IngredientDeleteRestoreResponse, error) {

	err := u.repo.IngredientDeleteRestore(ctx, req.IngredientIds)
	if err != nil {
		return nil, err
	}

	return &rmsv1.IngredientDeleteRestoreResponse{}, nil
}

func (u *ProductsUsecase) IngredientFind(ctx context.Context, req *rmsv1.IngredientFindRequest) (*rmsv1.IngredientFindResponse, error) {

	record, err := u.repo.IngredientFind(ctx, req.IngredientId)
	if err != nil {
		return nil, err
	}
	resp, err := u.adapter.IngredientFindGrpcFromSql(record)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *ProductsUsecase) IngredientFindForUpdate(ctx context.Context, req *rmsv1.IngredientFindForUpdateRequest) (*rmsv1.IngredientUpdateRequest, error) {
	record, err := u.repo.IngredientFindForUpdate(ctx, req.IngredientId)
	if err != nil {
		return nil, err
	}
	resp := u.adapter.IngredientFindForUpdateGrpcFromSql(record)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
