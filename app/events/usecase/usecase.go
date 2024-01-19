package service

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/meloneg/mln_rms_core/app/events/adapter"
	"github.com/meloneg/mln_rms_core/app/events/repo"
	"github.com/meloneg/mln_rms_core/common/auth"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

type EventsUsecaseInterface interface {
	EventCreate(ctx context.Context, req *rmsv1.EventCreateRequest) (*rmsv1.EventCreateResponse, error)
	EventUpdate(ctx context.Context, req *rmsv1.EventUpdateRequest) (*rmsv1.EventUpdateResponse, error)
	EventsList(ctx context.Context, req *rmsv1.EventsListRequest, authorizedUser *auth.Payload) (*rmsv1.EventsListResponse, error)
	EventDeleteRestore(ctx context.Context, req *rmsv1.EventDeleteRestoreRequest) (*rmsv1.EventDeleteRestoreResponse, error)
	EventsInputList(ctx context.Context, req *rmsv1.EventsInputListRequest) (*rmsv1.EventsInputListResponse, error)
}

type EventsUsecase struct {
	repo      repo.EventsRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.EventsAdapterInterface
}

func NewEventsUsecase(store db.Store, validator *protovalidate.Validator) EventsUsecaseInterface {
	repo := repo.NewEventsRepo(store)
	adapter := adapter.NewEventsAdapter()

	return &EventsUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
