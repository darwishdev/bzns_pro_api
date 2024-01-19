package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
<<<<<<< HEAD
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
=======
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
)

type AccountsAdapterInterface interface {
	PermissionsListGrpcFromSql(resp *[][]byte) (*rmsv1.PermissionsListResponse, error)
	rolesListRowGrpcFromSql(resp *db.RolesListRow) *rmsv1.RolesListRow
	RolesListGrpcFromSql(resp *[]db.RolesListRow) *rmsv1.RolesListResponse
	RoleCreateSqlFromGrpc(req *rmsv1.RoleCreateRequest) *db.RoleCreateTXParams
	RoleUpdateSqlFromGrpc(req *rmsv1.RoleUpdateRequest) *db.RoleUpdateTXParams
	RoleUpdateGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.RoleUpdateResponse
	RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.Role
	RoleFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.RoleUpdateRequest, error)
	RoleCreateGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.RoleCreateResponse
	RolesInputListGrpcFromSql(resp *[]db.RolesInputListRow) *rmsv1.RolesInputListResponse

	//userUserFindByEmailOrCode(ctx context.Context, req string) (*db.UserFindByEmailOrCodeRow, error)
	UserLoginGrpcFromSql(resp *db.UserFindByEmailOrCodeRow) (*rmsv1.UserLoginResponse, error)
	UserLoginGrpcFromCache(authSession *redisclient.AuthSession) (*rmsv1.UserLoginResponse, error)
	UserResetPasswordSqlFromGrpc(req *rmsv1.UserResetPasswordRequest) *db.UserResetPasswordParams
	UsersListGrpcFromSql(resp [][]byte) (*rmsv1.UsersListResponse, error)
	UserCreateSqlFromGrpc(req *rmsv1.UserCreateRequest) *db.UserCreateTXParams
	UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.User
	UserCreateGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.UserCreateResponse
	UserUpdateSqlFromGrpc(req *rmsv1.UserUpdateRequest) *db.UserUpdateTXParams
	UserUpdateGrpcFromSql(resp *db.AccountsSchemaUser) *rmsv1.UserUpdateResponse
	UserFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.UserUpdateRequest, error)
<<<<<<< HEAD
	UsersPermissionsMapFromSql(resp []db.UserPermissionsListRow) (*map[string]map[string]bool, error)
=======
	UsersPermissionsMapFromSql(resp []byte) (*map[string]map[string]bool, error)

	// customer

	accountsListRowGrpcFromSql(resp *db.AccountsListRow) *rmsv1.AccountsListRow
	AccountsListGrpcFromSql(resp *[]db.AccountsListRow) *rmsv1.AccountsListResponse
	AccountCreateSqlFromGrpc(req *rmsv1.AccountCreateRequest) *db.AccountCreateParams
	AccountEntityGrpcFromSql(resp *db.AccountsSchemaAccount) *rmsv1.Account
	AccountCreateGrpcFromSql(resp *db.AccountsSchemaAccount) *rmsv1.AccountCreateResponse
	AccountUpdateSqlFromGrpc(req *rmsv1.AccountUpdateRequest) *db.AccountUpdateParams
	AccountUpdateGrpcFromSql(resp *db.AccountsSchemaAccount) *rmsv1.AccountUpdateResponse
	AccountFindGrpcFromSql(resp *db.AccountFindRow) *rmsv1.AccountFindResponse
	AccountFindForUpdateGrpcFromSql(resp *db.AccountFindForUpdateRow) *rmsv1.AccountUpdateRequest

	// input
	AccountsInputListGrpcFromSql(resp *[]db.AccountsInputListRow) *rmsv1.AccountsInputListResponse
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
}

type AccountsAdapter struct {
	dateFormat string
}

func NewAccountsAdapter() *AccountsAdapter {
	return &AccountsAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
