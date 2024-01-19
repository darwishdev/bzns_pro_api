package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

type PublicRepoInterface interface {
	SettingsUpdate(ctx context.Context, req *db.SettingsUpdateParams) error
	SettingsFindForUpdate(ctx context.Context) (*[]db.SettingsFindForUpdateRow, error)
}

type PublicRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewPublicRepo(store db.Store) *PublicRepo {
	errorHandler := map[string]string{}
	return &PublicRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
