package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (api *Api) DistrictCreate(ctx context.Context, req *connect.Request[rmsv1.DistrictCreateRequest]) (*connect.Response[rmsv1.DistrictCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.DistrictCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) DistrictFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.DistrictFindForUpdateRequest]) (*connect.Response[rmsv1.DistrictUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.DistrictFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) DistrictUpdate(ctx context.Context, req *connect.Request[rmsv1.DistrictUpdateRequest]) (*connect.Response[rmsv1.DistrictUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.DistrictUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) DistrictsList(ctx context.Context, req *connect.Request[rmsv1.DistrictsListRequest]) (*connect.Response[rmsv1.DistrictsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.DistrictsList(ctx, req.Msg)
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

func (api *Api) DistrictsInputList(ctx context.Context, req *connect.Request[rmsv1.DistrictsInputListRequest]) (*connect.Response[rmsv1.DistrictsInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.DistrictsInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) DistrictDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.DistrictDeleteRestoreRequest]) (*connect.Response[rmsv1.DistrictDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.DistrictDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
