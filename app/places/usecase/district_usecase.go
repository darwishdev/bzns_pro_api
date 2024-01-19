package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (u *PlacesUsecase) DistrictCreate(ctx context.Context, req *rmsv1.DistrictCreateRequest) (*rmsv1.DistrictCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.DistrictCreateSqlFromGrpc(req)
	record, err := u.repo.DistrictCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.DistrictCreateGrpcFromSql(record), nil

}
func (u *PlacesUsecase) DistrictFindForUpdate(ctx context.Context, req *rmsv1.DistrictFindForUpdateRequest) (*rmsv1.DistrictUpdateRequest, error) {
	category, err := u.repo.DistrictFindForUpdate(ctx, &req.DistrictId)

	if err != nil {
		return nil, err
	}
	res, err := u.adapter.DistrictFindForUpdateGrpcFromSql(category)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PlacesUsecase) DistrictUpdate(ctx context.Context, req *rmsv1.DistrictUpdateRequest) (*rmsv1.DistrictUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.DistrictUpdateSqlFromGrpc(req)
	record, err := s.repo.DistrictUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.DistrictUpdateGrpcFromSql(record), nil

}

func (s *PlacesUsecase) DistrictsList(ctx context.Context, req *rmsv1.DistrictsListRequest) (*rmsv1.DistrictsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.DistrictsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.DistrictsListGrpcFromSql(record)

	// resp.Options = &rmsv1.ListDataOptions{
	// 	Title:       "districts_list",
	// 	Description: "districts_description",
	// 	CreateHandler: &rmsv1.CreateHandler{
	// 		Endpoint:  "districtCreate",
	// 		RouteName: "district_create",
	// 	},
	// 	UpdateHandler: &rmsv1.UpdateHandler{
	// 		RouteName:           "district_update",
	// 		FindRequestProperty: "districtId",
	// 		Endpoint:            "districtUpdate",
	// 		FindEndpoint:        "districtFindForUpdate",
	// 	},
	// 	DeleteRestoreHandler: &rmsv1.DeleteRestoreHandler{
	// 		Endpoint:        "districtDeleteRestore",
	// 		RequestProperty: "districtId",
	// 	},
	// 	ImportHandler: &rmsv1.ImportHandler{
	// 		Endpoint:           "users/districts/file",
	// 		ImportTemplateLink: "templates/users/Districts.xlsx",
	// 	},
	// }

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PlacesUsecase) DistrictDeleteRestore(ctx context.Context, req *rmsv1.DistrictDeleteRestoreRequest) (*rmsv1.DistrictDeleteRestoreResponse, error) {
	err := s.repo.DistrictDeleteRestore(ctx, req.DistrictIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.DistrictDeleteRestoreResponse{}, nil
}

func (s *PlacesUsecase) DistrictsInputList(ctx context.Context, req *rmsv1.DistrictsInputListRequest) (*rmsv1.DistrictsInputListResponse, error) {
	districts, err := s.repo.DistrictsInputList(ctx, req.CityId)
	if err != nil {
		return nil, err
	}
	res := s.adapter.DistrictsInputListGrpcFromSql(districts)

	return res, nil
}
