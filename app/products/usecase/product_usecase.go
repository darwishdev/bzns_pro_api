package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

func (u *ProductsUsecase) ProductCreate(ctx context.Context, req *rmsv1.ProductCreateRequest) (*rmsv1.ProductCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.ProductCreateSqlFromGrpc(req)
	err := u.repo.ProductCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ProductCreateResponse{}, nil

}
func (u *ProductsUsecase) ProductFindForUpdate(ctx context.Context, req *rmsv1.ProductFindForUpdateRequest) (*rmsv1.ProductUpdateRequest, error) {
	category, err := u.repo.ProductFindForUpdate(ctx, &req.ProductId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.ProductFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *ProductsUsecase) ProductUpdate(ctx context.Context, req *rmsv1.ProductUpdateRequest) (*rmsv1.ProductUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.ProductUpdateSqlFromGrpc(req)
	err := s.repo.ProductUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ProductUpdateResponse{}, nil

}

func (s *ProductsUsecase) ProductsList(ctx context.Context, req *rmsv1.ProductsListRequest) (*rmsv1.ProductsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.ProductsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.ProductsListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("products")

	// time.Sleep(4 * time.Second)
	return resp, nil
}
func (s *ProductsUsecase) ProductFind(ctx context.Context, req *rmsv1.ProductFindRequest) (*rmsv1.ProductFindResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.ProductFind(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	resp, err := s.adapter.ProductFindGrpcFromSql(*record)
	if err != nil {
		return nil, err
	}

	return &rmsv1.ProductFindResponse{Item: resp}, nil
}

func (s *ProductsUsecase) ProductsListForTransaction(ctx context.Context, req *rmsv1.ProductsListForTransactionRequest) (*rmsv1.ProductsListForTransactionResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.ProductsListForTransaction(ctx)
	if err != nil {
		return nil, err
	}
	resp := s.adapter.ProductsListForTransactionGrpcFromSql(record)
	return resp, nil
}

func (s *ProductsUsecase) ProductDeleteRestore(ctx context.Context, req *rmsv1.ProductDeleteRestoreRequest) (*rmsv1.ProductDeleteRestoreResponse, error) {
	err := s.repo.ProductDeleteRestore(ctx, req.ProductIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ProductDeleteRestoreResponse{}, nil
}

func (s *ProductsUsecase) ProductsInputList(ctx context.Context, req *rmsv1.ProductsInputListRequest) (*rmsv1.ProductsInputListResponse, error) {
	categories, err := s.repo.ProductsInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.ProductsInputListGrpcFromSql(categories)

	return res, nil
}

func (s *ProductsUsecase) StockItemsList(ctx context.Context, req *rmsv1.StockItemsListRequest) (*rmsv1.StockItemsListResponse, error) {
	categories, err := s.repo.StockItemsList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.StockItemsListGrpcFromSql(categories)

	return res, nil
}
