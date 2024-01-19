package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

func (api *Api) CityCreate(ctx context.Context, req *connect.Request[rmsv1.CityCreateRequest]) (*connect.Response[rmsv1.CityCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.CityCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CityFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.CityFindForUpdateRequest]) (*connect.Response[rmsv1.CityUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.CityFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) CityUpdate(ctx context.Context, req *connect.Request[rmsv1.CityUpdateRequest]) (*connect.Response[rmsv1.CityUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.CityUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CitiesList(ctx context.Context, req *connect.Request[rmsv1.CitiesListRequest]) (*connect.Response[rmsv1.CitiesListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.CitiesList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "accounts")
	if err != nil {
		return nil, err
	}
	resp.Options = opts

	return connect.NewResponse(resp), nil
}

func (api *Api) CitiesInputList(ctx context.Context, req *connect.Request[rmsv1.CitiesInputListRequest]) (*connect.Response[rmsv1.CitiesInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.CitiesInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) CityDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.CityDeleteRestoreRequest]) (*connect.Response[rmsv1.CityDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.CityDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
