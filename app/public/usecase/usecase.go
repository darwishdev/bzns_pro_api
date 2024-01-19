package service

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/meloneg/mln_rms_core/app/public/adapter"
	"github.com/meloneg/mln_rms_core/app/public/repo"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
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
