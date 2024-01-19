package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

func (u *ProductsUsecase) CategoryCreate(ctx context.Context, req *rmsv1.CategoryCreateRequest) (*rmsv1.CategoryCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.CategoryCreateSqlFromGrpc(req)
	record, err := u.repo.CategoryCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.CategoryCreateGrpcFromSql(record), nil

}
func (u *ProductsUsecase) CategoryFindForUpdate(ctx context.Context, req *rmsv1.CategoryFindForUpdateRequest) (*rmsv1.CategoryUpdateRequest, error) {
	category, err := u.repo.CategoryFindForUpdate(ctx, &req.CategoryId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.CategoryFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *ProductsUsecase) CategoryUpdate(ctx context.Context, req *rmsv1.CategoryUpdateRequest) (*rmsv1.CategoryUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.CategoryUpdateSqlFromGrpc(req)
	record, err := s.repo.CategoryUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.CategoryUpdateGrpcFromSql(record), nil

}

func (s *ProductsUsecase) CategoriesList(ctx context.Context, req *rmsv1.CategoriesListRequest) (*rmsv1.CategoriesListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.CategoriesList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.CategoriesListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("categories")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *ProductsUsecase) CategoryDeleteRestore(ctx context.Context, req *rmsv1.CategoryDeleteRestoreRequest) (*rmsv1.CategoryDeleteRestoreResponse, error) {
	err := s.repo.CategoryDeleteRestore(ctx, req.CategoryIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.CategoryDeleteRestoreResponse{}, nil
}

func (s *ProductsUsecase) CategoriesInputList(ctx context.Context, req *rmsv1.CategoriesInputListRequest) (*rmsv1.CategoriesInputListResponse, error) {
	categories, err := s.repo.CategoriesInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.CategoriesInputListGrpcFromSql(categories)

	return res, nil
}
