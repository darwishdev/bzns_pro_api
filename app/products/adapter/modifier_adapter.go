package adapter

import (
	"encoding/json"

	"github.com/meloneg/mln_rms_core/common/convertor"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

func (a *ProductsAdapter) modifiersListRowGrpcFromSql(resp *db.ModifiersListRow) (*rmsv1.ModifiersListRow, error) {
	serverResp := &rmsv1.ModifiersListRow{
		ModifierId:    resp.ModifierID,
		ModifierName:  resp.ModifierName,
		ModifierImage: resp.ModifierImage.String,
		MinChoices:    resp.MinChoices,
		MaxChoices:    resp.MaxChoices,
		CreatedAt:     resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:     resp.DeletedAt.Time.Format(a.dateFormat),
	}
	return serverResp, nil
}

func (a *ProductsAdapter) ModifiersListGrpcFromSql(resp *[]db.ModifiersListRow) (*rmsv1.ModifiersListResponse, error) {
	var modifiers []*rmsv1.ModifiersListRow
	var deletedModifiers []*rmsv1.ModifiersListRow
	for _, role := range *resp {
		c, err := a.modifiersListRowGrpcFromSql(&role)
		if err != nil {
			return nil, err
		}

		if role.DeletedAt.Valid {
			deletedModifiers = append(deletedModifiers, c)
		} else {
			modifiers = append(modifiers, c)
		}
	}
	return &rmsv1.ModifiersListResponse{
		Records:        modifiers,
		DeletedRecords: deletedModifiers,
	}, nil
}

func (a *ProductsAdapter) ModifierCreateSqlFromGrpc(req *rmsv1.ModifierCreateRequest) *db.ModifierCreateParams {
	resp := &db.ModifierCreateParams{
		InModifierName:         req.ModifierName,
		InModifierImage:        req.ModifierImage,
		InMinChoices:           req.MinChoices,
		InMaxChoices:           req.MaxChoices,
		InProductIds:           req.ProductIds,
		InModifierOptionNames:  []string{},
		InModifierOptionPrices: []float32{},
	}
	for _, v := range req.Options {
		resp.InModifierOptionNames = append(resp.InModifierOptionNames, v.ModifierOptionName)
		resp.InModifierOptionPrices = append(resp.InModifierOptionPrices, v.ModifierOptionPrice)
	}
	return resp
}

func (a *ProductsAdapter) ModifierUpdateSqlFromGrpc(req *rmsv1.ModifierUpdateRequest) *db.ModifierUpdateParams {
	resp := &db.ModifierUpdateParams{
		InModifierID:                   req.ModifierId,
		InModifierName:                 req.ModifierName,
		InModifierImage:                req.ModifierImage,
		InMinChoices:                   req.MinChoices,
		InMaxChoices:                   req.MaxChoices,
		InModifierOptionIdsToDetach:    req.OptionsToDetach,
		InModifierOptionIdsToUpdate:    []int32{},
		InModifierOptionNamesToUpdate:  []string{},
		InModifierOptionPricesToUpdate: []float32{},
		InModifierOptionNamesToAttach:  []string{},
		InModifierOptionPricesToAttach: []float32{},
	}
	for _, v := range req.OptionsToUpdate {
		resp.InModifierOptionIdsToUpdate = append(resp.InModifierOptionIdsToUpdate, v.ModifierOptionId)
		resp.InModifierOptionNamesToUpdate = append(resp.InModifierOptionNamesToUpdate, v.ModifierOptionName)
		resp.InModifierOptionPricesToUpdate = append(resp.InModifierOptionPricesToUpdate, v.ModifierOptionPrice)
	}
	for _, v := range req.OptionsToAttach {
		resp.InModifierOptionNamesToAttach = append(resp.InModifierOptionNamesToAttach, v.ModifierOptionName)
		resp.InModifierOptionPricesToAttach = append(resp.InModifierOptionPricesToAttach, v.ModifierOptionPrice)
	}
	return resp
}

func (a *ProductsAdapter) ModifierFindGrpcFromSql(resp *db.ProductsSchemaModifiersView) (*rmsv1.ModifierFindResponse, error) {
	var (
		options  []*rmsv1.ModifierOption
		products []*rmsv1.ModifierProduct
	)
	if resp.Options != nil {
		if len(resp.Options) > 0 {
			err := json.Unmarshal(resp.Options, &options)
			if err != nil {
				return nil, err
			}
		}
	}
	if resp.Products != nil {
		if len(resp.Products) > 0 {
			err := json.Unmarshal(resp.Products, &products)
			if err != nil {
				return nil, err
			}
		}
	}
	return &rmsv1.ModifierFindResponse{
		ModifierId:    int32(resp.ModifierID),
		ModifierName:  resp.ModifierName,
		ModifierImage: resp.ModifierImage.String,
		MinChoices:    resp.MinChoices,
		MaxChoices:    resp.MaxChoices,
		Options:       options,
		Products:      products,
		OptionsCount:  int32(resp.OptionsCount),
		CreatedAt:     resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:     resp.DeletedAt.Time.Format(a.dateFormat),
	}, nil

}

func (a *ProductsAdapter) ModifiersInputListGrpcFromSql(resp *[]db.ModifiersInputListRow) *rmsv1.ModifiersInputListResponse {
	// ModifiersInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.ModifierID, v.ModifierName)
		records = append(records, record)
	}
	return &rmsv1.ModifiersInputListResponse{
		Options: records,
	}
}

// func (a *ProductsAdapter) ModifierFindForUpdateGrpcFromSql(resp *db.ModifierFindForUpdateRow) *rmsv1.ModifierUpdateRequest {
// 	return &rmsv1.ModifierUpdateRequest{
// 		ModifierName:  resp.ModifierName,
// 		ModifierImage: resp.ModifierImage.String,
// 		WasteRatio:    resp.WasteRatio,
// 		CategoryId:    resp.CategoryID,
// 		UnitId:        resp.UnitID,
// 	}

// }
