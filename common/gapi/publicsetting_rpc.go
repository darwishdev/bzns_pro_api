package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

func (api *Api) SettingsUpdate(ctx context.Context, req *connect.Request[rmsv1.SettingsUpdateRequest]) (*connect.Response[rmsv1.SettingsUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	err := api.publicUsecase.SettingsUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&rmsv1.SettingsUpdateResponse{}), nil
}

func (api *Api) SettingsFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.SettingsFindForUpdateRequest]) (*connect.Response[rmsv1.SettingsFindForUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.SettingsFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
