package adapter

import (
	"encoding/json"

	"github.com/meloneg/mln_rms_core/common/convertor"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

//list

func (a *EventsAdapter) eventsListRowGrpcFromSql(resp *db.EventsListRow) *rmsv1.EventsListRow {
	return &rmsv1.EventsListRow{
		EventId:          resp.EventID,
		EventName:        resp.EventName,
		ConstructorTitle: resp.ConstructorTitle,
		EventDate:        resp.EventDate.Time.Format(a.dateFormat),
		EventImage:       resp.EventImage.String,
		CreatedAt:        resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:        resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:        resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *EventsAdapter) EventsListGrpcFromSql(resp *[]db.EventsListRow) *rmsv1.EventsListResponse {
	records := make([]*rmsv1.EventsListRow, 0)
	deletedRecords := make([]*rmsv1.EventsListRow, 0)
	for _, v := range *resp {
		record := a.eventsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.EventsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *EventsAdapter) EventsInputListGrpcFromSql(resp *[]db.EventsInputListRow) *rmsv1.EventsInputListResponse {
	// EventsInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.EventID, v.EventName)
		records = append(records, record)
	}
	return &rmsv1.EventsInputListResponse{
		Options: records,
	}
}

func (a *EventsAdapter) EventCreateSqlFromGrpc(req *rmsv1.EventCreateRequest) *db.EventCreateParams {

	return &db.EventCreateParams{

		EventName:        req.EventName,
		EventDescription: convertor.ToPgType(req.EventDescription),
	}
}

func (a *EventsAdapter) EventFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.EventUpdateRequest, error) {
	var response rmsv1.EventUpdateRequest
	err := json.Unmarshal([]byte(*resp), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *EventsAdapter) EventUpdateSqlFromGrpc(req *rmsv1.EventUpdateRequest) *db.EventUpdateParams {

	return &db.EventUpdateParams{

		EventID:          req.EventId,
		EventName:        req.EventName,
		EventDescription: convertor.ToPgType(req.EventDescription),
	}
}
