package adapter

import (
	"encoding/json"

	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
	"github.com/meloneg/mln_rms_core/common/redisclient"
	"github.com/rs/zerolog/log"
)

//createa

func (a *OrdersAdapter) OrderCreateSqlFromGrpc(req *rmsv1.OrderCreateRequest, authUser *redisclient.AuthSession) *db.OrderCreateParams {
	var (
		IngredientTypeId     int32 = 2
		ModifierOptionTypeId int32 = 1
	)
	resp := &db.OrderCreateParams{
		InOrderNote:         req.OrderNote,
		InCreatedBy:         authUser.AccountID,
		InEntityID:          authUser.EntityID,
		InSessionID:         authUser.SessionID,
		InHallTableID:       req.HallTableId,
		InCustomerID:        req.CustomerId,
		InOrderTypeID:       req.OrderTypeId,
		InDiscount:          req.Discount,
		InProductIds:        []int32{},
		InProductSequances:  []int32{},
		InProductQuantities: []float32{},
		InNotes:             []string{},
		InExtrasIds:         []int32{},
		InExtrasQuantities:  []float32{},
		InExtrasTypeIds:     []int32{},
	}

	for _, product := range req.Products {

		if (len(product.Ingredients) == 0 && len(product.ModifierOptions) == 0) || (product.Ingredients == nil && product.ModifierOptions == nil) {
			resp.InProductIds = append(resp.InProductIds, product.ProductId)
			resp.InProductSequances = append(resp.InProductSequances, product.ProductSequance)
			resp.InNotes = append(resp.InNotes, product.Note)
			resp.InProductQuantities = append(resp.InProductQuantities, product.Quantity)
		}

		log.Debug().Interface("from adapter", product.Ingredients).Interface("mos", product.ModifierOptions).Int("length", len(product.Ingredients)).Msg("debug")
		for _, productIngredient := range product.Ingredients {
			resp.InProductIds = append(resp.InProductIds, product.ProductId)
			resp.InProductSequances = append(resp.InProductSequances, product.ProductSequance)
			resp.InNotes = append(resp.InNotes, product.Note)
			resp.InProductQuantities = append(resp.InProductQuantities, product.Quantity)
			resp.InExtrasIds = append(resp.InExtrasIds, productIngredient.ExtraId)
			resp.InExtrasQuantities = append(resp.InExtrasQuantities, productIngredient.Quantity)
			resp.InExtrasTypeIds = append(resp.InExtrasTypeIds, IngredientTypeId)
		}
		for _, productModifierOption := range product.ModifierOptions {
			resp.InProductIds = append(resp.InProductIds, product.ProductId)
			resp.InProductSequances = append(resp.InProductSequances, product.ProductSequance)
			resp.InNotes = append(resp.InNotes, product.Note)
			resp.InProductQuantities = append(resp.InProductQuantities, product.Quantity)
			resp.InExtrasIds = append(resp.InExtrasIds, productModifierOption.ExtraId)
			resp.InExtrasQuantities = append(resp.InExtrasQuantities, productModifierOption.Quantity)
			resp.InExtrasTypeIds = append(resp.InExtrasTypeIds, ModifierOptionTypeId)
		}
	}
	return resp
}

//close

func (a *OrdersAdapter) OrderCloseSqlFromGrpc(req *rmsv1.OrderCloseRequest, authUser *redisclient.AuthSession) *db.OrderCloseParams {
	resp := &db.OrderCloseParams{
		InClosedBy: authUser.AccountID,
		InOrderID:  req.OrderId,
		InTips:     req.Tips,
	}

	return resp
}

func (a *OrdersAdapter) OrderCommitSqlFromGrpc(req *rmsv1.OrderCommitRequest, authUser *redisclient.AuthSession) *db.OrderCommitParams {
	resp := &db.OrderCommitParams{
		InCommittedBy:   authUser.AccountID,
		InOrderID:       req.OrderId,
		InPaymentTypeID: req.PaymentTypeId,
		InTips:          req.Tips,
	}

	return resp
}

func (a *OrdersAdapter) OrderProductCreateSqlFromGrpc(req *rmsv1.OrderProductCreateRequest, authUser *redisclient.AuthSession) *db.OrderProductCreateParams {
	var (
		IngredientTypeId     int32 = 2
		ModifierOptionTypeId int32 = 1
	)
	resp := &db.OrderProductCreateParams{
		InOrderID:           req.OrderId,
		InProductIds:        []int32{},
		InProductSequances:  []int32{},
		InProductQuantities: []float32{},
		InNotes:             []string{},
		InExtrasIds:         []int32{},
		InExtrasQuantities:  []float32{},
		InExtrasTypeIds:     []int32{},
	}

	if (len(req.Ingredients) == 0 && len(req.ModifierOptions) == 0) || (req.Ingredients == nil && req.ModifierOptions == nil) {
		resp.InProductIds = append(resp.InProductIds, req.ProductId)
		resp.InProductSequances = append(resp.InProductSequances, req.ProductSequance)
		resp.InNotes = append(resp.InNotes, req.Note)
		resp.InProductQuantities = append(resp.InProductQuantities, req.Quantity)
	}

	log.Debug().Interface("from adapter", req.Ingredients).Interface("mos", req.ModifierOptions).Int("length", len(req.Ingredients)).Msg("debug")
	for _, productIngredient := range req.Ingredients {
		resp.InProductIds = append(resp.InProductIds, req.ProductId)
		resp.InProductSequances = append(resp.InProductSequances, req.ProductSequance)
		resp.InNotes = append(resp.InNotes, req.Note)
		resp.InProductQuantities = append(resp.InProductQuantities, req.Quantity)
		resp.InExtrasIds = append(resp.InExtrasIds, productIngredient.ExtraId)
		resp.InExtrasQuantities = append(resp.InExtrasQuantities, productIngredient.Quantity)
		resp.InExtrasTypeIds = append(resp.InExtrasTypeIds, IngredientTypeId)
	}
	for _, productModifierOption := range req.ModifierOptions {
		resp.InProductIds = append(resp.InProductIds, req.ProductId)
		resp.InProductSequances = append(resp.InProductSequances, req.ProductSequance)
		resp.InNotes = append(resp.InNotes, req.Note)
		resp.InProductQuantities = append(resp.InProductQuantities, req.Quantity)
		resp.InExtrasIds = append(resp.InExtrasIds, productModifierOption.ExtraId)
		resp.InExtrasQuantities = append(resp.InExtrasQuantities, productModifierOption.Quantity)
		resp.InExtrasTypeIds = append(resp.InExtrasTypeIds, ModifierOptionTypeId)
	}

	return resp
}
func (a *OrdersAdapter) OrderCreateGrpcFromSqlc(resp *db.OrderActiveFindRow) (*rmsv1.ActiveOrdersViewRow, error) {
	var products []*rmsv1.OrderProduct
	err := json.Unmarshal(resp.OrderProducts, &products)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ActiveOrdersViewRow{
		OrderId:         resp.OrderID,
		Subtotal:        float64(resp.Subtotal),
		ClosedAt:        resp.ClosedAt.Time.Format(a.dateFormat),
		EntityId:        resp.EntityID,
		TotalTax:        float64(resp.TotalTax),
		AddressId:       resp.AddressID.Int32,
		OrderCode:       resp.OrderCode,
		OrderNote:       resp.OrderNote.String,
		OrderType:       resp.OrderType,
		SessionId:       resp.SessionID,
		CustomerId:      resp.CustomerID,
		EntityName:      resp.EntityName,
		TaxPercent:      float64(resp.TaxPercent),
		TotalPrice:      float64(resp.TotalPrice),
		ClosedById:      resp.ClosedByID,
		CreatedById:     resp.CreatedByID,
		CustomerName:    resp.CustomerName,
		HallTableId:     resp.HallTableID,
		OrderTypeId:     resp.OrderTypeID,
		ClosedByName:    resp.ClosedByName,
		TotalDiscount:   float64(resp.TotalDiscount),
		CreatedByName:   resp.CreatedByName,
		ServicePercent:  float64(resp.ServicePercent),
		DiscountPercent: float64(resp.DiscountPercent),
		TotalServiceFee: float64(resp.TotalServiceFee),
		OrderProducts:   products,
	}, nil

}
func (a *OrdersAdapter) OrderProductUpdateSqlFromGrpc(req *rmsv1.OrderProductUpdateRequest, authUser *redisclient.AuthSession) *db.OrderProductUpdateParams {

	resp := &db.OrderProductUpdateParams{
		InOrderID:         req.OrderId,
		InProductID:       req.ProductId,
		InProductSequance: req.ProductSequance,
		InQuantity:        req.Quantity,
		InNote:            req.Note,
		InVoiedBy:         0,
	}
	if req.IsVoided {
		resp.InVoiedBy = authUser.AccountID
	}

	return resp
}

func (a *OrdersAdapter) OrderProductExtraUpdateSqlFromGrpc(req *rmsv1.OrderProductExtraUpdateRequest, authUser *redisclient.AuthSession) *db.OrderProductExtraUpdateParams {

	resp := &db.OrderProductExtraUpdateParams{
		InOrderID:         req.OrderId,
		InProductID:       req.ProductId,
		InProductSequance: req.ProductSequance,
		InExtraTypeID:     req.ExtraTypeId,
		InExtraID:         req.ExtraId,
		InExtraQuantity:   req.Quantity,
	}

	return resp
}

func (a *OrdersAdapter) OrderProductExtraAttachDetachSqlFromGrpc(req *rmsv1.OrderProductExtraAttachDetachRequest, authUser *redisclient.AuthSession) *db.OrderProductExtraAttachDetachParams {

	resp := &db.OrderProductExtraAttachDetachParams{
		InOrderID:             req.OrderId,
		InProductID:           req.ProductId,
		InProductSequance:     req.ProductSequance,
		InExtraTypeID:         req.ExtraTypeId,
		InAttachExtraID:       req.AttachExtraId,
		InAttachExtraQuantity: req.Quantity,
		InDetachExtraID:       req.DetachExtraId,
	}

	return resp
}

func (a *OrdersAdapter) OrderUpdateSqlFromGrpc(req *rmsv1.OrderUpdateRequest, authUser *redisclient.AuthSession) *db.OrderUpdateParams {

	return &db.OrderUpdateParams{
		InOrderID:   req.OrderId,
		InOrderNote: req.OrderNote,

		InHallTableID: req.HallTableId,
		InCustomerID:  req.CustomerId,

		InDiscount: req.Discount,
	}
}
