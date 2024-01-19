package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

func (u *ProductsUsecase) UnitCreate(ctx context.Context, req *rmsv1.UnitCreateRequest) (*rmsv1.UnitCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.UnitCreateSqlFromGrpc(req)
	record, err := u.repo.UnitCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.UnitCreateGrpcFromSql(record), nil

}
func (u *ProductsUsecase) UnitFindForUpdate(ctx context.Context, req *rmsv1.UnitFindForUpdateRequest) (*rmsv1.UnitUpdateRequest, error) {
	category, err := u.repo.UnitFindForUpdate(ctx, &req.UnitId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.UnitFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *ProductsUsecase) UnitUpdate(ctx context.Context, req *rmsv1.UnitUpdateRequest) (*rmsv1.UnitUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.UnitUpdateSqlFromGrpc(req)
	record, err := s.repo.UnitUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.UnitUpdateGrpcFromSql(record), nil

}

func (s *ProductsUsecase) UnitsList(ctx context.Context, req *rmsv1.UnitsListRequest) (*rmsv1.UnitsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.UnitsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.UnitsListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("units")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *ProductsUsecase) UnitDeleteRestore(ctx context.Context, req *rmsv1.UnitDeleteRestoreRequest) (*rmsv1.UnitDeleteRestoreResponse, error) {
	err := s.repo.UnitDeleteRestore(ctx, req.UnitIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.UnitDeleteRestoreResponse{}, nil
}

func (s *ProductsUsecase) UnitsInputList(ctx context.Context, req *rmsv1.UnitsInputListRequest) (*rmsv1.UnitsInputListResponse, error) {
	cities, err := s.repo.UnitsInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.UnitsInputListGrpcFromSql(cities)

	return res, nil
}
