package adapter

import (
	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

type EventsAdapterInterface interface {
	eventsListRowGrpcFromSql(resp *db.EventsListRow) *rmsv1.EventsListRow
	EventsListGrpcFromSql(resp *[]db.EventsListRow) *rmsv1.EventsListResponse
	EventsInputListGrpcFromSql(resp *[]db.EventsInputListRow) *rmsv1.EventsInputListResponse
	EventCreateSqlFromGrpc(req *rmsv1.EventCreateRequest) *db.EventCreateParams
	EventFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.EventUpdateRequest, error)
	EventUpdateSqlFromGrpc(req *rmsv1.EventUpdateRequest) *db.EventUpdateParams
}

type EventsAdapter struct {
	dateFormat string
}

func NewEventsAdapter() EventsAdapterInterface {
	return &EventsAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
