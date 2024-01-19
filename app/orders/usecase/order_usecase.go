package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
	"github.com/meloneg/mln_rms_core/common/redisclient"
)

func (u *OrdersUsecase) OrderCreate(ctx context.Context, req *rmsv1.OrderCreateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	params := u.adapter.OrderCreateSqlFromGrpc(req, authSession)
	record, err := u.repo.OrderCreate(ctx, params)
	if err != nil {
		return nil, err
	}

	order, err := u.repo.OrderActiveFind(ctx, record)
	if err != nil {
		return nil, err
	}
	response, err := u.adapter.OrderCreateGrpcFromSqlc(order)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderCreateResponse{
		Order: response,
	}, nil
}

func (u *OrdersUsecase) OrderUpdate(ctx context.Context, req *rmsv1.OrderUpdateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderUpdateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderUpdateSqlFromGrpc(req, authSession)
	err := u.repo.OrderUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderUpdateResponse{}, nil
}

func (u *OrdersUsecase) OrderClose(ctx context.Context, req *rmsv1.OrderCloseRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderCloseResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderCloseSqlFromGrpc(req, authSession)
	err := u.repo.OrderClose(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderCloseResponse{
		OrderId: req.OrderId,
	}, nil
}
func (u *OrdersUsecase) OrderCommit(ctx context.Context, req *rmsv1.OrderCommitRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderCommitResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderCommitSqlFromGrpc(req, authSession)
	err := u.repo.OrderCommit(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderCommitResponse{
		OrderId: req.OrderId,
	}, nil
}

func (u *OrdersUsecase) OrderProductCreate(ctx context.Context, req *rmsv1.OrderProductCreateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderProductCreateSqlFromGrpc(req, authSession)
	err := u.repo.OrderProductCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderProductCreateResponse{}, nil
}

func (u *OrdersUsecase) OrderProductUpdate(ctx context.Context, req *rmsv1.OrderProductUpdateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductUpdateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderProductUpdateSqlFromGrpc(req, authSession)
	err := u.repo.OrderProductUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderProductUpdateResponse{}, nil
}

func (u *OrdersUsecase) OrderProductExtraUpdate(ctx context.Context, req *rmsv1.OrderProductExtraUpdateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductExtraUpdateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderProductExtraUpdateSqlFromGrpc(req, authSession)
	err := u.repo.OrderProductExtraUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderProductExtraUpdateResponse{}, nil
}

func (u *OrdersUsecase) OrderProductExtraAttachDetach(ctx context.Context, req *rmsv1.OrderProductExtraAttachDetachRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductExtraAttachDetachResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OrderProductExtraAttachDetachSqlFromGrpc(req, authSession)
	err := u.repo.OrderProductExtraAttachDetach(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.OrderProductExtraAttachDetachResponse{}, nil
}
