package repo

import (
	"context"

	db "github.com/meloneg/mln_rms_core/common/db/gen"
)

type EventsRepoInterface interface {
	EventsInputList(ctx context.Context) (*[]db.EventsInputListRow, error)
	EventsList(ctx context.Context) (*[]db.EventsListRow, error)
	EventCreate(ctx context.Context, req *db.EventCreateParams) (*db.EventsSchemaEvent, error)
	EventUpdate(ctx context.Context, req *db.EventUpdateParams) (*db.EventsSchemaEvent, error)
	EventDeleteRestore(ctx context.Context, req []int32) error
}

type EventsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewEventsRepo(store db.Store) EventsRepoInterface {
	errorHandler := map[string]string{
		"events_event_name_key": "eventName",
	}
	return &EventsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
