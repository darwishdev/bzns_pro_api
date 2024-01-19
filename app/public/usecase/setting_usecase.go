package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

func (s *PublicUsecase) SettingsUpdate(ctx context.Context, req *rmsv1.SettingsUpdateRequest) error {
	if err := s.validator.Validate(req); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.SettingsUpdateSqlFromGrpc(req)
	err := s.repo.SettingsUpdate(ctx, params)
	if err != nil {
		return err
	}
	return nil

}

func (u *PublicUsecase) SettingsFindForUpdate(ctx context.Context, req *rmsv1.SettingsFindForUpdateRequest) (*rmsv1.SettingsFindForUpdateResponse, error) {
	settings, err := u.repo.SettingsFindForUpdate(ctx)

	if err != nil {
		return nil, err
	}
	resp := u.adapter.SettingsFindForUpdateGrpcFromSql(settings)

	return resp, nil
}
