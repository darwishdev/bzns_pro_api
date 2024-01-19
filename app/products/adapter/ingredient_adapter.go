package adapter

import (
	"encoding/json"

	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (a *ProductsAdapter) ingredientsListRowGrpcFromSql(resp *db.IngredientsListRow) (*rmsv1.IngredientsListRow, error) {
	serverResp := &rmsv1.IngredientsListRow{
		IngredientId:       int32(resp.IngredientID),
		IngredientName:     resp.IngredientName,
		IngredientImage:    resp.IngredientImage.String,
		IngredientCalories: resp.IngredientCalories.Float32,
		IngredientCost:     resp.IngredientCost,
		WasteRatio:         resp.WasteRatio,
		UnitId:             resp.UnitID,
		UnitBuy:            resp.UnitBuy,
		UnitSell:           resp.UnitSell,
		UnitRatio:          float32(resp.UnitRatio),
		CategoryName:       resp.CategoryName,
		CategoryId:         int32(resp.CategoryID),
		ProductsCount:      int32(resp.ProductsCount),
		CreatedAt:          resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:          resp.DeletedAt.Time.Format(a.dateFormat),
	}
	return serverResp, nil
}

func (a *ProductsAdapter) IngredientsListGrpcFromSql(resp *[]db.IngredientsListRow) (*rmsv1.IngredientsListResponse, error) {
	var ingredients []*rmsv1.IngredientsListRow
	var Ingredients []*rmsv1.IngredientsListRow
	for _, role := range *resp {
		c, err := a.ingredientsListRowGrpcFromSql(&role)
		if err != nil {
			return nil, err
		}

		if role.DeletedAt.Valid {
			Ingredients = append(Ingredients, c)
		} else {
			ingredients = append(ingredients, c)
		}
	}
	return &rmsv1.IngredientsListResponse{
		Records:        ingredients,
		DeletedRecords: Ingredients,
	}, nil
}

func (a *ProductsAdapter) IngredientCreateSqlFromGrpc(req *rmsv1.IngredientCreateRequest) *db.IngredientCreateParams {
	return &db.IngredientCreateParams{
		IngredientName:  req.IngredientName,
		IngredientImage: convertor.ToPgType(req.IngredientImage),
		WasteRatio:      req.WasteRatio,
		CategoryID:      req.CategoryId,
		UnitID:          req.UnitId,
	}
}

func (a *ProductsAdapter) IngredientUpdateSqlFromGrpc(req *rmsv1.IngredientUpdateRequest) *db.IngredientUpdateParams {
	return &db.IngredientUpdateParams{
		IngredientID:    req.IngredientId,
		IngredientName:  req.IngredientName,
		IngredientImage: convertor.ToPgType(req.IngredientImage),
		WasteRatio:      req.WasteRatio,
		CategoryID:      req.CategoryId,
		UnitID:          req.UnitId,
	}
}

func (a *ProductsAdapter) IngredientFindGrpcFromSql(resp *db.IngredientFindRow) (*rmsv1.IngredientFindResponse, error) {
	var (
		ingredientProduts         []*rmsv1.IngredientPorudctRow
		ingredientModifierOptions []*rmsv1.IngredientModifierOptionRow
	)
	if resp.Products != nil {
		if len(resp.Products) > 0 {
			err := json.Unmarshal(resp.Products, &ingredientProduts)
			if err != nil {
				return nil, err
			}
		}
	}
	if resp.ModifierOptions != nil {
		if len(resp.ModifierOptions) > 0 {
			err := json.Unmarshal(resp.ModifierOptions, &ingredientModifierOptions)
			if err != nil {
				return nil, err
			}
		}
	}
	return &rmsv1.IngredientFindResponse{
		IngredientId:       int32(resp.IngredientID),
		IngredientName:     resp.IngredientName,
		IngredientImage:    resp.IngredientImage.String,
		IngredientCalories: resp.IngredientCalories.Float32,
		IngredientCost:     resp.IngredientCost,
		WasteRatio:         resp.WasteRatio,
		UnitId:             resp.UnitID,
		UnitBuy:            resp.UnitBuy,
		UnitSell:           resp.UnitSell,
		UnitRatio:          float32(resp.UnitRatio),
		CategoryName:       resp.CategoryName,
		CategoryId:         int32(resp.CategoryID),
		ProductsCount:      int32(resp.ProductsCount),
		CreatedAt:          resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:          resp.DeletedAt.Time.Format(a.dateFormat),
		Products:           ingredientProduts,
		ModifierOptions:    ingredientModifierOptions,
	}, nil

}

func (a *ProductsAdapter) IngredientsInputListGrpcFromSql(resp *[]db.IngredientsInputListRow) *rmsv1.IngredientsInputListResponse {
	// IngredientsInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.IngredientID, v.IngredientName)
		records = append(records, record)
	}
	return &rmsv1.IngredientsInputListResponse{
		Options: records,
	}
}
func (a *ProductsAdapter) IngredientFindForUpdateGrpcFromSql(resp *db.IngredientFindForUpdateRow) *rmsv1.IngredientUpdateRequest {
	return &rmsv1.IngredientUpdateRequest{
		IngredientName:  resp.IngredientName,
		IngredientImage: resp.IngredientImage.String,
		WasteRatio:      resp.WasteRatio,
		CategoryId:      resp.CategoryID,
		UnitId:          resp.UnitID,
	}

}
