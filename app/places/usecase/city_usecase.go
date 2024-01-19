package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (u *PlacesUsecase) CityCreate(ctx context.Context, req *rmsv1.CityCreateRequest) (*rmsv1.CityCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.CityCreateSqlFromGrpc(req)
	record, err := u.repo.CityCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.CityCreateGrpcFromSql(record), nil

}
func (u *PlacesUsecase) CityFindForUpdate(ctx context.Context, req *rmsv1.CityFindForUpdateRequest) (*rmsv1.CityUpdateRequest, error) {
	category, err := u.repo.CityFindForUpdate(ctx, &req.CityId)

	if err != nil {
		return nil, err
	}
	res, err := u.adapter.CityFindForUpdateGrpcFromSql(category)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PlacesUsecase) CityUpdate(ctx context.Context, req *rmsv1.CityUpdateRequest) (*rmsv1.CityUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.CityUpdateSqlFromGrpc(req)
	record, err := s.repo.CityUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.CityUpdateGrpcFromSql(record), nil

}

func (s *PlacesUsecase) CitiesList(ctx context.Context, req *rmsv1.CitiesListRequest) (*rmsv1.CitiesListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.CitiesList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.CitiesListGrpcFromSql(record)

	// resp.Options = &rmsv1.ListDataOptions{
	// 	Title:       "cities_list",
	// 	Description: "cities_description",
	// 	CreateHandler: &rmsv1.CreateHandler{
	// 		Endpoint:  "cityCreate",
	// 		RouteName: "city_create",
	// 	},
	// 	UpdateHandler: &rmsv1.UpdateHandler{
	// 		RouteName:           "city_update",
	// 		FindRequestProperty: "cityId",
	// 		Endpoint:            "cityUpdate",
	// 		FindEndpoint:        "cityFindForUpdate",
	// 	},
	// 	DeleteRestoreHandler: &rmsv1.DeleteRestoreHandler{
	// 		Endpoint:        "cityDeleteRestore",
	// 		RequestProperty: "cityId",
	// 	},
	// 	ImportHandler: &rmsv1.ImportHandler{
	// 		Endpoint:           "users/cities/file",
	// 		ImportTemplateLink: "templates/users/Cities.xlsx",
	// 	},
	// }

	// // time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PlacesUsecase) CityDeleteRestore(ctx context.Context, req *rmsv1.CityDeleteRestoreRequest) (*rmsv1.CityDeleteRestoreResponse, error) {
	err := s.repo.CityDeleteRestore(ctx, req.CityIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.CityDeleteRestoreResponse{}, nil
}

func (s *PlacesUsecase) CitiesInputList(ctx context.Context, req *rmsv1.CitiesInputListRequest) (*rmsv1.CitiesInputListResponse, error) {
	cities, err := s.repo.CitiesInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.CitiesInputListGrpcFromSql(cities)

	return res, nil
}
