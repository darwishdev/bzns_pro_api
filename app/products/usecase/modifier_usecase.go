package service

import (
	"context"

	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

func (u *ProductsUsecase) ModifiersList(ctx context.Context, req *rmsv1.ModifiersListRequest) (*rmsv1.ModifiersListResponse, error) {
	record, err := u.repo.ModifiersList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := u.adapter.ModifiersListGrpcFromSql(&record)
	if err != nil {
		return nil, err
	}

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("modifiers")

	return resp, nil
}
func (u *ProductsUsecase) ModifiersInputList(ctx context.Context, req *rmsv1.ModifiersInputListRequest) (*rmsv1.ModifiersInputListResponse, error) {
	record, err := u.repo.ModifiersInputList(ctx)
	if err != nil {
		return nil, err
	}
	resp := u.adapter.ModifiersInputListGrpcFromSql(&record)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *ProductsUsecase) ModifierCreate(ctx context.Context, req *rmsv1.ModifierCreateRequest) (*rmsv1.ModifierCreateResponse, error) {
	param := u.adapter.ModifierCreateSqlFromGrpc(req)

	err := u.repo.ModifierCreate(ctx, *param)
	if err != nil {
		return nil, err
	}

	return &rmsv1.ModifierCreateResponse{}, nil
}

func (u *ProductsUsecase) ModifierUpdate(ctx context.Context, req *rmsv1.ModifierUpdateRequest) (*rmsv1.ModifierUpdateResponse, error) {
	param := u.adapter.ModifierUpdateSqlFromGrpc(req)

	err := u.repo.ModifierUpdate(ctx, *param)
	if err != nil {
		return nil, err
	}

	return &rmsv1.ModifierUpdateResponse{}, nil
}
func (u *ProductsUsecase) ModifierDeleteRestore(ctx context.Context, req *rmsv1.ModifierDeleteRestoreRequest) (*rmsv1.ModifierDeleteRestoreResponse, error) {

	err := u.repo.ModifierDeleteRestore(ctx, req.ModifierIds)
	if err != nil {
		return nil, err
	}

	return &rmsv1.ModifierDeleteRestoreResponse{}, nil
}

func (u *ProductsUsecase) ModifierFind(ctx context.Context, req *rmsv1.ModifierFindRequest) (*rmsv1.ModifierFindResponse, error) {

	record, err := u.repo.ModifierFind(ctx, req.ModifierId)
	if err != nil {
		return nil, err
	}
	resp, err := u.adapter.ModifierFindGrpcFromSql(record)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// func (u *ProductsUsecase) ModifierFindForUpdate(ctx context.Context, req *rmsv1.ModifierFindForUpdateRequest) (*rmsv1.ModifierUpdateRequest, error) {
// 	record, err := u.repo.ModifierFindForUpdate(ctx, req.ModifierId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp := u.adapter.ModifierFindForUpdateGrpcFromSql(record)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }
