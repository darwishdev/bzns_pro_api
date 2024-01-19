package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (api *Api) AccountsList(ctx context.Context, req *connect.Request[rmsv1.AccountsListRequest]) (*connect.Response[rmsv1.AccountsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountsList(ctx, req.Msg)
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

func (api *Api) AccountsInputList(ctx context.Context, req *connect.Request[rmsv1.AccountsInputListRequest]) (*connect.Response[rmsv1.AccountsInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountsInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) AccountDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.AccountDeleteRestoreRequest]) (*connect.Response[rmsv1.AccountDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) AccountCreate(ctx context.Context, req *connect.Request[rmsv1.AccountCreateRequest]) (*connect.Response[rmsv1.AccountCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) AccountFind(ctx context.Context, req *connect.Request[rmsv1.AccountFindRequest]) (*connect.Response[rmsv1.AccountFindResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountFind(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) AccountFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.AccountFindForUpdateRequest]) (*connect.Response[rmsv1.AccountUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) AccountUpdate(ctx context.Context, req *connect.Request[rmsv1.AccountUpdateRequest]) (*connect.Response[rmsv1.AccountUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.AccountUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
