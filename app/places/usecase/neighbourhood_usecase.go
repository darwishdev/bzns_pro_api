package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (u *PlacesUsecase) NeighbourhoodCreate(ctx context.Context, req *rmsv1.NeighbourhoodCreateRequest) (*rmsv1.NeighbourhoodCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.NeighbourhoodCreateSqlFromGrpc(req)
	record, err := u.repo.NeighbourhoodCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.NeighbourhoodCreateGrpcFromSql(record), nil

}
func (u *PlacesUsecase) NeighbourhoodFindForUpdate(ctx context.Context, req *rmsv1.NeighbourhoodFindForUpdateRequest) (*rmsv1.NeighbourhoodUpdateRequest, error) {
	category, err := u.repo.NeighbourhoodFindForUpdate(ctx, &req.NeighbourhoodId)

	if err != nil {
		return nil, err
	}
	res, err := u.adapter.NeighbourhoodFindForUpdateGrpcFromSql(category)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PlacesUsecase) NeighbourhoodUpdate(ctx context.Context, req *rmsv1.NeighbourhoodUpdateRequest) (*rmsv1.NeighbourhoodUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.NeighbourhoodUpdateSqlFromGrpc(req)
	record, err := s.repo.NeighbourhoodUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.NeighbourhoodUpdateGrpcFromSql(record), nil

}

func (s *PlacesUsecase) NeighbourhoodsList(ctx context.Context, req *rmsv1.NeighbourhoodsListRequest) (*rmsv1.NeighbourhoodsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.NeighbourhoodsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.NeighbourhoodsListGrpcFromSql(record)

	// resp.Options = &rmsv1.ListDataOptions{
	// 	Title:       "neighbourhoods_list",
	// 	Description: "neighbourhoods_description",
	// 	CreateHandler: &rmsv1.CreateHandler{
	// 		Endpoint:  "neighbourhoodCreate",
	// 		RouteName: "neighbourhood_create",
	// 	},
	// 	UpdateHandler: &rmsv1.UpdateHandler{
	// 		RouteName:           "neighbourhood_update",
	// 		FindRequestProperty: "neighbourhoodId",
	// 		Endpoint:            "neighbourhoodUpdate",
	// 		FindEndpoint:        "neighbourhoodFindForUpdate",
	// 	},
	// 	DeleteRestoreHandler: &rmsv1.DeleteRestoreHandler{
	// 		Endpoint:        "neighbourhoodDeleteRestore",
	// 		RequestProperty: "neighbourhoodId",
	// 	},
	// 	ImportHandler: &rmsv1.ImportHandler{
	// 		Endpoint:           "users/neighbourhoods/file",
	// 		ImportTemplateLink: "templates/users/Neighbourhoods.xlsx",
	// 	},
	// }

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PlacesUsecase) NeighbourhoodDeleteRestore(ctx context.Context, req *rmsv1.NeighbourhoodDeleteRestoreRequest) (*rmsv1.NeighbourhoodDeleteRestoreResponse, error) {
	err := s.repo.NeighbourhoodDeleteRestore(ctx, req.NeighbourhoodIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.NeighbourhoodDeleteRestoreResponse{}, nil
}

func (s *PlacesUsecase) NeighbourhoodsInputList(ctx context.Context, req *rmsv1.NeighbourhoodsInputListRequest) (*rmsv1.NeighbourhoodsInputListResponse, error) {
	neighbourhoods, err := s.repo.NeighbourhoodsInputList(ctx, req.DistrictId)
	if err != nil {
		return nil, err
	}
	res := s.adapter.NeighbourhoodsInputListGrpcFromSql(neighbourhoods)

	return res, nil
}
