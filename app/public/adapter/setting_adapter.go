package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (a *PublicAdapter) SettingsUpdateSqlFromGrpc(req *rmsv1.SettingsUpdateRequest) *db.SettingsUpdateParams {
	keys := make([]string, len(req.Settings))
	values := make([]string, len(req.Settings))
	for index, v := range req.Settings {
		keys[index] = v.SettingKey
		values[index] = v.SettingValue
	}
	return &db.SettingsUpdateParams{
		Keys:   keys,
		Values: values,
	}
}
func (a *PublicAdapter) SettingsEntityGrpcFromSql(resp []db.Setting) []*rmsv1.Setting {
	grpcResp := make([]*rmsv1.Setting, len(resp))
	for _, v := range resp {
		record := &rmsv1.Setting{
			SettingKey:   v.SettingKey,
			SettingValue: v.SettingValue,
		}
		grpcResp = append(grpcResp, record)
	}
	return grpcResp

}

func (a *PublicAdapter) SettingsFindForUpdateGrpcFromSql(resp *[]db.SettingsFindForUpdateRow) *rmsv1.SettingsFindForUpdateResponse {
	grpcRows := make([]*rmsv1.SettingsFindForUpdateRow, len(*resp))
	for index, v := range *resp {
		grpcRow := &rmsv1.SettingsFindForUpdateRow{
			SettingKey:   v.SettingKey,
			SettingValue: v.SettingValue,
			SettingType:  v.SettingType,
		}

		grpcRows[index] = grpcRow

	}

	return &rmsv1.SettingsFindForUpdateResponse{
		Settings: grpcRows,
	}

}
