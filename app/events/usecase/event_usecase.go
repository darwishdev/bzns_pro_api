package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/darwishdev/bzns_pro_api/common/auth"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *EventsUsecase) EventCreate(ctx context.Context, req *rmsv1.EventCreateRequest) (*rmsv1.EventCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.EventCreateSqlFromGrpc(req)
	_, err := u.repo.EventCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventCreateResponse{}, nil

}

func (s *EventsUsecase) EventUpdate(ctx context.Context, req *rmsv1.EventUpdateRequest) (*rmsv1.EventUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.EventUpdateSqlFromGrpc(req)
	_, err := s.repo.EventUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventUpdateResponse{}, nil

}

func (s *EventsUsecase) EventsList(ctx context.Context, req *rmsv1.EventsListRequest, authorizedUser *auth.Payload) (*rmsv1.EventsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.EventsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.EventsListGrpcFromSql(record)

	resp.Options = authorizedUser.GetAccessableActionsForGroup("roles")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *EventsUsecase) EventDeleteRestore(ctx context.Context, req *rmsv1.EventDeleteRestoreRequest) (*rmsv1.EventDeleteRestoreResponse, error) {
	err := s.repo.EventDeleteRestore(ctx, req.EventIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventDeleteRestoreResponse{}, nil
}

func (s *EventsUsecase) EventsInputList(ctx context.Context, req *rmsv1.EventsInputListRequest) (*rmsv1.EventsInputListResponse, error) {
	roles, err := s.repo.EventsInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.EventsInputListGrpcFromSql(roles)

	return res, nil
}
