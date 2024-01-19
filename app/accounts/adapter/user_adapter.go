package adapter

import (
	"encoding/json"

	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
<<<<<<< HEAD
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
=======
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	"github.com/rs/zerolog/log"

	"golang.org/x/crypto/bcrypt"
)

func (a *AccountsAdapter) UserLoginGrpcFromCache(authSession *redisclient.AuthSession) (*rmsv1.UserLoginResponse, error) {
	response := rmsv1.UserLoginResponse{
		UserId:   authSession.UserID,
		UserName: authSession.UserName,
		SideBar:  authSession.SideBar,
	}

	return &response, nil
}

func (a *AccountsAdapter) UserLoginGrpcFromSql(resp *db.UserFindByEmailOrCodeRow) (*rmsv1.UserLoginResponse, error) {
	var sidebar []*rmsv1.SideBarItem
	err := json.Unmarshal(resp.SideBar, &sidebar)
	if err != nil {
		return nil, err
	}
	return &rmsv1.UserLoginResponse{
		UserId:    resp.UserID,
		UserName:  resp.UserName,
		AccountId: resp.AccountID,
		SessionId: int32(resp.SessionID),
		EntityId:  resp.EntityID,
		SideBar:   sidebar,
	}, nil
}

func (a *AccountsAdapter) UsersListGrpcFromSql(resp [][]byte) (*rmsv1.UsersListResponse, error) {
	records := make([]*rmsv1.UsersListRow, 0)
	deletedRecords := make([]*rmsv1.UsersListRow, 0)
	for _, v := range resp {
		var record rmsv1.UsersListRow
		err := json.Unmarshal([]byte(v), &record)

		if record.DeletedAt == "" {
			log.Debug().Interface("record", record.DeletedAt == "").Msg("parse")
			records = append(records, &record)
		} else {
			log.Debug().Interface("deleted record", record.DeletedAt == "").Msg("deleted  parse")
			deletedRecords = append(records, &record)
		}
		if err != nil {
			return nil, err
		}
	}
	response := &rmsv1.UsersListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}

	// log.Debug().Interface("respo", record.DeletedAt == "").Msg("deleted  parse")

	return response, nil
}

type permissionGroupRow struct {
	Authorized         bool   `json:"authorized"`
	PermissionFunction string `json:"permission_function"`
}
type permissionGroups struct {
	Permissions     []permissionGroupRow `json:"permissions"`
	PermissionGroup string               `json:"permission_group"`
}

func (a *AccountsAdapter) UsersPermissionsMapFromSql(respBytes []byte) (*map[string]map[string]bool, error) {
	var resp []permissionGroups
	err := json.Unmarshal(respBytes, &resp)
	if err != nil {
		return nil, err
	}
	response := make(map[string]map[string]bool, len(resp))
	// var parsedResp []*permissionGroupRow
	for _, group := range resp {
		// var permissions []permissionGroupRow
		// err := json.Unmarshal(group.Permissions, &permissions)
		// if err != nil {
		// 	return nil, err
		// }

		groupPermissions := make(map[string]bool, len(group.Permissions))
		for _, permission := range group.Permissions {
			groupPermissions[permission.PermissionFunction] = permission.Authorized
		}
		response[group.PermissionGroup] = groupPermissions

	}

	return &response, nil
}

func (a *AccountsAdapter) UserResetPasswordSqlFromGrpc(req *rmsv1.UserResetPasswordRequest) *db.UserResetPasswordParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)

	return &db.UserResetPasswordParams{
		UserEmail:    req.UserEmail,
		UserPassword: string(hashedPassword),
	}

}

func (a *AccountsAdapter) UserCreateSqlFromGrpc(req *rmsv1.UserCreateRequest) *db.UserCreateTXParams {
	permissionsParams := make([]db.UserPermissionsBulkCreateParams, 0)
	rolesParams := make([]db.UserRolesBulkCreateParams, 0)
	for _, v := range req.Permissions {
		userPermission := db.UserPermissionsBulkCreateParams{
			PermissionID: v,
		}
		permissionsParams = append(permissionsParams, userPermission)
	}
	for _, v := range req.Roles {
		userRole := db.UserRolesBulkCreateParams{
			RoleID: v,
		}
		rolesParams = append(rolesParams, userRole)
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)
	return &db.UserCreateTXParams{
		UserParams: db.UserCreateParams{
			UserName:     req.UserName,
			UserImage:    convertor.ToPgType(req.UserImage),
			UserEmail:    req.UserEmail,
			UserPhone:    convertor.ToPgType(req.UserPhone),
			UserPassword: string(hashedPassword),
		},
		PermissionsParams: permissionsParams,
		RolesParams:       rolesParams,
	}
}
func (a *AccountsAdapter) UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.User {
	formatLayout := "2006-01-02 15:04:05"
	formattedCreatedTime := resp.CreatedAt.Time.Format(formatLayout)
	formattedUpdatedTime := resp.UpdatedAt.Time.Format(formatLayout)
	return &rmsv1.User{
		UserId:    int32(resp.UserID),
		UserName:  resp.UserName,
		UserImage: resp.UserImage.String,
		UserEmail: resp.UserEmail,
		UserPhone: resp.UserPhone.String,
		CreatedAt: formattedCreatedTime,
		UpdatedAt: formattedUpdatedTime,
	}

}
func (a *AccountsAdapter) UserCreateGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.UserCreateResponse {
	return &rmsv1.UserCreateResponse{
		User: a.UserEntityGrpcFromSql(resp),
	}
}

func (a *AccountsAdapter) UserUpdateSqlFromGrpc(req *rmsv1.UserUpdateRequest) *db.UserUpdateTXParams {
	permissionsParams := make([]db.UserPermissionsBulkCreateParams, 0)
	rolesParams := make([]db.UserRolesBulkCreateParams, 0)
	for _, v := range req.Permissions {
		userPermission := db.UserPermissionsBulkCreateParams{
			PermissionID: v,
		}
		permissionsParams = append(permissionsParams, userPermission)
	}
	for _, v := range req.Roles {
		userRole := db.UserRolesBulkCreateParams{
			RoleID: v,
		}
		rolesParams = append(rolesParams, userRole)
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)

	return &db.UserUpdateTXParams{
		UserParams: db.UserUpdateParams{
			UserID:       req.UserId,
			UserName:     req.UserName,
			UserImage:    convertor.ToPgType(req.UserImage),
			UserEmail:    req.UserEmail,
			UserPhone:    convertor.ToPgType(req.UserPhone),
			UserPassword: string(hashedPassword),
		},
		PermissionsParams: permissionsParams,
		RolesParams:       rolesParams,
	}
}
func (a *AccountsAdapter) UserUpdateGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.UserUpdateResponse {
	return &rmsv1.UserUpdateResponse{
		User: a.UserEntityGrpcFromSql(resp),
	}
}

func (a *AccountsAdapter) UserFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.UserUpdateRequest, error) {
	var response rmsv1.UserUpdateRequest
	err := json.Unmarshal([]byte(*resp), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
