package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (api *Api) NeighbourhoodCreate(ctx context.Context, req *connect.Request[rmsv1.NeighbourhoodCreateRequest]) (*connect.Response[rmsv1.NeighbourhoodCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.NeighbourhoodCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) NeighbourhoodFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.NeighbourhoodFindForUpdateRequest]) (*connect.Response[rmsv1.NeighbourhoodUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.NeighbourhoodFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) NeighbourhoodUpdate(ctx context.Context, req *connect.Request[rmsv1.NeighbourhoodUpdateRequest]) (*connect.Response[rmsv1.NeighbourhoodUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.NeighbourhoodUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) NeighbourhoodsList(ctx context.Context, req *connect.Request[rmsv1.NeighbourhoodsListRequest]) (*connect.Response[rmsv1.NeighbourhoodsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.NeighbourhoodsList(ctx, req.Msg)
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

func (api *Api) NeighbourhoodsInputList(ctx context.Context, req *connect.Request[rmsv1.NeighbourhoodsInputListRequest]) (*connect.Response[rmsv1.NeighbourhoodsInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.NeighbourhoodsInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) NeighbourhoodDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.NeighbourhoodDeleteRestoreRequest]) (*connect.Response[rmsv1.NeighbourhoodDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.placesUsecase.NeighbourhoodDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
