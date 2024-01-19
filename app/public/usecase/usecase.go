package service

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/public/adapter"
	"github.com/darwishdev/bzns_pro_api/app/public/repo"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type PublicUsecaseInterface interface {
	SettingsUpdate(ctx context.Context, req *rmsv1.SettingsUpdateRequest) error
	SettingsFindForUpdate(ctx context.Context, req *rmsv1.SettingsFindForUpdateRequest) (*rmsv1.SettingsFindForUpdateResponse, error)
}

type PublicUsecase struct {
	repo      repo.PublicRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.PublicAdapterInterface
}

func NewPublicUsecase(store db.Store, validator *protovalidate.Validator) *PublicUsecase {
	repo := repo.NewPublicRepo(store)
	adapter := adapter.NewPublicAdapter()

	return &PublicUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
