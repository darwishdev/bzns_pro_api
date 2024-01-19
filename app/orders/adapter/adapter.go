package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
)

type OrdersAdapterInterface interface {
	OrderCreateSqlFromGrpc(req *rmsv1.OrderCreateRequest, authUser *redisclient.AuthSession) *db.OrderCreateParams
	OrderCloseSqlFromGrpc(req *rmsv1.OrderCloseRequest, authUser *redisclient.AuthSession) *db.OrderCloseParams
	OrderCommitSqlFromGrpc(req *rmsv1.OrderCommitRequest, authUser *redisclient.AuthSession) *db.OrderCommitParams
	OrderProductCreateSqlFromGrpc(req *rmsv1.OrderProductCreateRequest, authUser *redisclient.AuthSession) *db.OrderProductCreateParams
	OrderProductUpdateSqlFromGrpc(req *rmsv1.OrderProductUpdateRequest, authUser *redisclient.AuthSession) *db.OrderProductUpdateParams
	OrderProductExtraUpdateSqlFromGrpc(req *rmsv1.OrderProductExtraUpdateRequest, authUser *redisclient.AuthSession) *db.OrderProductExtraUpdateParams
	OrderUpdateSqlFromGrpc(req *rmsv1.OrderUpdateRequest, authUser *redisclient.AuthSession) *db.OrderUpdateParams
	OrderProductExtraAttachDetachSqlFromGrpc(req *rmsv1.OrderProductExtraAttachDetachRequest, authUser *redisclient.AuthSession) *db.OrderProductExtraAttachDetachParams
	OrderCreateGrpcFromSqlc(resp *db.OrderActiveFindRow) (*rmsv1.ActiveOrdersViewRow, error)
}

type OrdersAdapter struct {
	dateFormat string
}

func NewOrdersAdapter() OrdersAdapterInterface {
	return &OrdersAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
