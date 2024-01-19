package adapter

import (
	"encoding/json"

	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (a *ProductsAdapter) productsListRowGrpcFromSql(resp *db.ProductsListRow) *rmsv1.ProductsListRow {
	return &rmsv1.ProductsListRow{
		ProductId:          resp.ProductID,
		ProductName:        resp.ProductName,
		ProductCode:        resp.ProductCode.String,
		ProductDescription: resp.ProductDescription.String,
		ProductImage:       resp.ProductImage.String,
		IsFinal:            resp.IsFinal.Bool,
		UnitId:             resp.UnitID,
		UnitName:           resp.UnitName,
		CategoryId:         resp.CategoryID,
		CategoryName:       resp.CategoryName,
		ProductCost:        resp.ProductCost,
		ProductPrice:       resp.ProductPrice,
		CreatedAt:          resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:          resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:          resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *ProductsAdapter) ProductsListGrpcFromSql(resp *[]db.ProductsListRow) *rmsv1.ProductsListResponse {
	records := make([]*rmsv1.ProductsListRow, 0)
	deletedRecords := make([]*rmsv1.ProductsListRow, 0)
	for _, v := range *resp {

		record := a.productsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.ProductsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *ProductsAdapter) ProductCreateSqlFromGrpc(req *rmsv1.ProductCreateRequest) *db.ProductCreateParams {
	resp := &db.ProductCreateParams{
		InProductName:              req.ProductName,
		InProductImage:             req.ProductImage,
		InProductDescription:       req.ProductDescription,
		InIsFinal:                  req.IsFinal,
		InCategoryID:               req.CategoryId,
		InUnitID:                   req.UnitId,
		InProductCode:              req.ProductCode,
		InProductPrice:             req.ProductPrice,
		InIngredientIds:            []int32{},
		InIngredientQuantities:     []float32{},
		InIngredientRemovable:      []bool{},
		InIngredientIncreaseLimits: []int32{},
		InIngredientPrices:         []float32{},
		InModifierIds:              req.ProductModifiers,
	}

	if req.ProductIngredients != nil {
		if len(req.ProductIngredients) > 0 {
			for _, ingredient := range req.ProductIngredients {
				resp.InIngredientIds = append(resp.InIngredientIds, ingredient.IngredientId)
				resp.InIngredientQuantities = append(resp.InIngredientQuantities, ingredient.Quantity)
				resp.InIngredientIncreaseLimits = append(resp.InIngredientIncreaseLimits, ingredient.IncreaseLimit)
				resp.InIngredientRemovable = append(resp.InIngredientRemovable, ingredient.IngredientRemovable)
				resp.InIngredientPrices = append(resp.InIngredientPrices, ingredient.IngredientPrice)

			}
		}
	}

	return resp
}
func (a *ProductsAdapter) ProductEntityGrpcFromSql(resp *db.ProductsSchemaProduct) *rmsv1.Product {
	return &rmsv1.Product{
		ProductId:          resp.ProductID,
		ProductName:        resp.ProductName,
		ProductCode:        resp.ProductCode.String,
		ProductDescription: resp.ProductDescription.String,
		ProductImage:       resp.ProductImage.String,
		IsFinal:            resp.IsFinal.Bool,
		CategoryId:         resp.CategoryID,
		UnitId:             resp.UnitID,
		ProductCost:        resp.ProductCost,
		ProductPrice:       resp.ProductPrice,
		CreatedAt:          resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:          resp.DeletedAt.Time.Format(a.dateFormat),
		UpdatedAt:          resp.UpdatedAt.Time.Format(a.dateFormat),
	}

}
func (a *ProductsAdapter) ProductCreateGrpcFromSql(resp *db.ProductsSchemaProduct) *rmsv1.ProductCreateResponse {
	return &rmsv1.ProductCreateResponse{
		// Product: a.ProductEntityGrpcFromSql(resp),
	}
}

func (a *ProductsAdapter) ProductUpdateSqlFromGrpc(req *rmsv1.ProductUpdateRequest) *db.ProductUpdateParams {
	resp := &db.ProductUpdateParams{
		InProductID:                        req.ProductId,
		InProductName:                      req.ProductName,
		InProductImage:                     req.ProductImage,
		InProductDescription:               req.ProductDescription,
		InIsFinal:                          req.IsFinal,
		InCategoryID:                       req.CategoryId,
		InUnitID:                           req.UnitId,
		InProductCode:                      req.ProductCode,
		InProductPrice:                     req.ProductPrice,
		InIngredientIdsToDetach:            req.IngredientsToDetach,
		InIngredientIdsToAttach:            make([]int32, len(req.IngredientsToAttach)),
		InIngredientIncreaseLimitsToAttach: make([]int32, len(req.IngredientsToAttach)),
		InIngredientQuantitiesToAttach:     make([]float32, len(req.IngredientsToAttach)),
		InIngredientPricesToAttach:         make([]float32, len(req.IngredientsToAttach)),
		InIngredientRemovableToAttach:      make([]bool, len(req.IngredientsToAttach)),
		InIngredientIdsToUpdate:            make([]int32, len(req.IngredientsToUpdate)),
		InIngredientIncreaseLimitsToUpdate: make([]int32, len(req.IngredientsToUpdate)),
		InIngredientQuantitiesToUpdate:     make([]float32, len(req.IngredientsToUpdate)),
		InIngredientPricesToUpdate:         make([]float32, len(req.IngredientsToUpdate)),
		InIngredientRemovableToUpdate:      make([]bool, len(req.IngredientsToUpdate)),
		InModifierIdsToDetach:              req.IngredientsToDetach,
		InModifierIdsToAttach:              req.ModifiersToAttach,
	}

	if len(req.IngredientsToAttach) > 0 {
		for index, ingredient := range req.IngredientsToAttach {
			resp.InIngredientIdsToAttach[index] = ingredient.IngredientId
			resp.InIngredientIncreaseLimitsToAttach[index] = ingredient.IncreaseLimit
			resp.InIngredientQuantitiesToAttach[index] = ingredient.Quantity
			resp.InIngredientPricesToAttach[index] = ingredient.IngredientPrice
			resp.InIngredientPricesToAttach[index] = ingredient.IngredientPrice
		}
	}
	if len(req.IngredientsToUpdate) > 0 {
		for index, ingredient := range req.IngredientsToUpdate {
			resp.InIngredientIdsToUpdate[index] = ingredient.IngredientId
			resp.InIngredientIncreaseLimitsToUpdate[index] = ingredient.IncreaseLimit
			resp.InIngredientQuantitiesToUpdate[index] = ingredient.Quantity
			resp.InIngredientPricesToUpdate[index] = ingredient.IngredientPrice
			resp.InIngredientPricesToUpdate[index] = ingredient.IngredientPrice
		}
	}

	return resp
}
func (a *ProductsAdapter) ProductsInputListGrpcFromSql(resp *[]db.ProductsInputListRow) *rmsv1.ProductsInputListResponse {
	// ProductsInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.ProductID, v.ProductName)
		records = append(records, record)
	}
	return &rmsv1.ProductsInputListResponse{
		Options: records,
	}
}
func (a *ProductsAdapter) ProductFindForUpdateGrpcFromSql(resp *db.ProductFindForUpdateRow) *rmsv1.ProductUpdateRequest {
	return &rmsv1.ProductUpdateRequest{
		ProductId:          resp.ProductID,
		ProductName:        resp.ProductName,
		ProductCode:        resp.ProductCode.String,
		ProductDescription: resp.ProductDescription.String,
		ProductImage:       resp.ProductImage.String,
		IsFinal:            resp.IsFinal.Bool,
		CategoryId:         resp.CategoryID,
		UnitId:             resp.UnitID,
		ProductPrice:       resp.ProductPrice,
	}

}

func (a *ProductsAdapter) productsListForTransactionRowGrpcFromSql(resp *db.ProductsListForTransactionRow) *rmsv1.ProductsListForTransactionRow {
	return &rmsv1.ProductsListForTransactionRow{
		ProductId:   resp.ProductID,
		ProductName: resp.ProductName,
		ProductCode: resp.ProductCode.String,
		UnitId:      resp.UnitID,
		UnitBuy:     resp.UnitBuy,
		UnitSell:    resp.UnitSell,
		UnitRatio:   resp.UnitRatio,
		ProductCost: resp.ProductCost,
	}
}

func (a *ProductsAdapter) ProductsListForTransactionGrpcFromSql(resp *[]db.ProductsListForTransactionRow) *rmsv1.ProductsListForTransactionResponse {
	records := make([]*rmsv1.ProductsListForTransactionRow, 0)
	for _, v := range *resp {
		record := a.productsListForTransactionRowGrpcFromSql(&v)

		records = append(records, record)

	}
	return &rmsv1.ProductsListForTransactionResponse{
		Records: records,
	}
}

func (a *ProductsAdapter) stockItemsListRowGrpcFromSql(resp *db.ProductsSchemaStockItemsView) *rmsv1.StockItemsListRow {
	return &rmsv1.StockItemsListRow{
		ItemId:           resp.ItemID,
		ItemCost:         resp.ItemCost,
		TransactionPrice: resp.ItemCost,
		ItemCode:         resp.ItemCode.String,
		ItemName:         resp.ItemName,
		ItemTypeId:       resp.ItemTypeID,
		ItemType:         resp.ItemType,
		CategoryId:       resp.CategoryID,
		CategoryName:     resp.CategoryName,
		UnitId:           resp.UnitID,
		UnitBuy:          resp.UnitBuy,
		UnitSell:         resp.UnitSell,
		UnitRatio:        resp.UnitRatio,
		Quantity:         0,
		Valuation:        0,
	}
}

func (a *ProductsAdapter) StockItemsListGrpcFromSql(resp *[]db.ProductsSchemaStockItemsView) *rmsv1.StockItemsListResponse {
	records := make([]*rmsv1.StockItemsListRow, 0)
	for _, v := range *resp {
		record := a.stockItemsListRowGrpcFromSql(&v)

		records = append(records, record)

	}
	return &rmsv1.StockItemsListResponse{
		Records: records,
	}
}

func (a *ProductsAdapter) ProductFindGrpcFromSql(resp db.ProductsSchemaProductsView) (*rmsv1.ProductsViewRow, error) {
	var modifiers []*rmsv1.ProductModifiersView

	if resp.ProductModifiers != nil {
		if err := json.Unmarshal(resp.ProductModifiers, &modifiers); err != nil {
			return nil, err
		}
	}

	var ingredients []*rmsv1.ProductIngredeintsView
	if resp.ProductIngredients != nil {
		if err := json.Unmarshal(resp.ProductIngredients, &ingredients); err != nil {
			return nil, err
		}
	}

	response := rmsv1.ProductsViewRow{
		ProductId:          uint32(resp.ProductID),
		ProductName:        resp.ProductName,
		ProductCode:        resp.ProductCode.String,
		ProductDescription: resp.ProductDescription.String,
		ProductImage:       resp.ProductImage.String,
		IsFinal:            resp.IsFinal.Bool,
		ProductCost:        resp.ProductCost,
		ProductPrice:       resp.ProductPrice,
		CreatedAt:          resp.CreatedAt.Time.Format(a.dateFormat),
		ItemTypeId:         uint32(resp.ItemTypeID),
		ItemType:           resp.ItemType,
		CategoryId:         uint32(resp.CategoryID),
		CategoryName:       resp.CategoryName,
		UnitId:             uint32(resp.UnitID),
		UnitBuy:            resp.UnitBuy,
		UnitSell:           resp.UnitSell,
		UnitRatio:          resp.UnitRatio,
		Modfieirs:          modifiers,
		Ingredients:        ingredients,
	}
	return &response, nil
}
