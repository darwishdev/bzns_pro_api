package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

func (api *Api) RoleCreate(ctx context.Context, req *connect.Request[rmsv1.RoleCreateRequest]) (*connect.Response[rmsv1.RoleCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RoleCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) RoleFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.RoleFindForUpdateRequest]) (*connect.Response[rmsv1.RoleUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RoleFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) RoleUpdate(ctx context.Context, req *connect.Request[rmsv1.RoleUpdateRequest]) (*connect.Response[rmsv1.RoleUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RoleUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) PermissionsList(ctx context.Context, req *connect.Request[rmsv1.PermissionsListRequest]) (*connect.Response[rmsv1.PermissionsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.PermissionsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) RolesList(ctx context.Context, req *connect.Request[rmsv1.RolesListRequest]) (*connect.Response[rmsv1.RolesListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RolesList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "roles")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) RolesInputList(ctx context.Context, req *connect.Request[rmsv1.RolesInputListRequest]) (*connect.Response[rmsv1.RolesInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RolesInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) RoleDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.RoleDeleteRestoreRequest]) (*connect.Response[rmsv1.RoleDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RoleDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
