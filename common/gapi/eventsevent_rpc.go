package gapi

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

func (api *Api) EventCreate(ctx context.Context, req *connect.Request[rmsv1.EventCreateRequest]) (*connect.Response[rmsv1.EventCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) EventUpdate(ctx context.Context, req *connect.Request[rmsv1.EventUpdateRequest]) (*connect.Response[rmsv1.EventUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) EventsList(ctx context.Context, req *connect.Request[rmsv1.EventsListRequest]) (*connect.Response[rmsv1.EventsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	authorizedUser, err := api.authorizeUser(req.Header())
	if err != nil {
		return nil, err
	}
	if !authorizedUser.Can("events", "EventsList") {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("don't_have_permission"))
	}

	resp, err := api.eventsUsecase.EventsList(ctx, req.Msg, authorizedUser)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) EventsInputList(ctx context.Context, req *connect.Request[rmsv1.EventsInputListRequest]) (*connect.Response[rmsv1.EventsInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventsInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) EventDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.EventDeleteRestoreRequest]) (*connect.Response[rmsv1.EventDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
