package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

func (api *Api) UserLogin(ctx context.Context, req *connect.Request[rmsv1.UserLoginRequest]) (*connect.Response[rmsv1.UserLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserLogin(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UsersList(ctx context.Context, req *connect.Request[rmsv1.UsersListRequest]) (*connect.Response[rmsv1.UsersListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UsersList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "users")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) UserAuthorize(ctx context.Context, req *connect.Request[rmsv1.UserAuthorizeRequest]) (*connect.Response[rmsv1.UserLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	authorizedUser, _, err := api.authorizeUser(req.Header())
	if err != nil {
		return nil, err
	}
	resp, err := api.accountsUsecase.UserAuthorize(ctx, authorizedUser)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UserResetPassword(ctx context.Context, req *connect.Request[rmsv1.UserResetPasswordRequest]) (*connect.Response[rmsv1.UserResetPasswordResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	err := api.accountsUsecase.UserResetPassword(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	resp := &rmsv1.UserResetPasswordResponse{}
	return connect.NewResponse(resp), nil
}

func (api *Api) UserDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.UserDeleteRestoreRequest]) (*connect.Response[rmsv1.UserDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UserCreate(ctx context.Context, req *connect.Request[rmsv1.UserCreateRequest]) (*connect.Response[rmsv1.UserCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UserFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.UserFindForUpdateRequest]) (*connect.Response[rmsv1.UserUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) UserUpdate(ctx context.Context, req *connect.Request[rmsv1.UserUpdateRequest]) (*connect.Response[rmsv1.UserUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
