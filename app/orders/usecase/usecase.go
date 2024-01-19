package service

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/orders/adapter"
	"github.com/darwishdev/bzns_pro_api/app/orders/repo"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
)

type OrdersUsecaseInterface interface {
	OrderCreate(ctx context.Context, req *rmsv1.OrderCreateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderCreateResponse, error)
	OrderUpdate(ctx context.Context, req *rmsv1.OrderUpdateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderUpdateResponse, error)
	OrderClose(ctx context.Context, req *rmsv1.OrderCloseRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderCloseResponse, error)
	OrderCommit(ctx context.Context, req *rmsv1.OrderCommitRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderCommitResponse, error)
	OrderProductCreate(ctx context.Context, req *rmsv1.OrderProductCreateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductCreateResponse, error)
	OrderProductUpdate(ctx context.Context, req *rmsv1.OrderProductUpdateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductUpdateResponse, error)
	OrderProductExtraUpdate(ctx context.Context, req *rmsv1.OrderProductExtraUpdateRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductExtraUpdateResponse, error)
	OrderProductExtraAttachDetach(ctx context.Context, req *rmsv1.OrderProductExtraAttachDetachRequest, authSession *redisclient.AuthSession) (*rmsv1.OrderProductExtraAttachDetachResponse, error)
}

type OrdersUsecase struct {
	repo      repo.OrdersRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.OrdersAdapterInterface
}

func NewOrdersUsecase(store db.Store, validator *protovalidate.Validator) OrdersUsecaseInterface {
	repo := repo.NewOrdersRepo(store)
	adapter := adapter.NewOrdersAdapter()

	return &OrdersUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
