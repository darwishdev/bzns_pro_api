package repo

import (
	"context"

	db "github.com/meloneg/mln_rms_core/common/db/gen"
)

type AccountsRepoInterface interface {
	// role
	PermissionsList(ctx context.Context) (*[][]byte, error)
	RoleCreate(ctx context.Context, req *db.RoleCreateTXParams) (*db.AccountsSchemaRole, error)
	RoleUpdate(ctx context.Context, req *db.RoleUpdateTXParams) (*db.AccountsSchemaRole, error)
	RoleFindForUpdate(ctx context.Context, req *int32) (*[]byte, error)
	RolesList(ctx context.Context) (*[]db.RolesListRow, error)
	RoleDeleteRestore(ctx context.Context, req []int32) error
	RolesInputList(ctx context.Context) (*[]db.RolesInputListRow, error)
	// user
	UserFind(ctx context.Context, req int32) (*[]byte, error)
	UserFindByEmailOrCode(ctx context.Context, req string) (*db.UserFindByEmailOrCodeRow, error)
	UserResetPassword(ctx context.Context, req *db.UserResetPasswordParams) error
	UsersList(ctx context.Context) (*[][]byte, error)
	UserDeleteRestore(ctx context.Context, req []int32) error
	UserCreate(ctx context.Context, req *db.UserCreateTXParams) (*db.AccountsSchemaUser, error)
	UserFindForUpdate(ctx context.Context, req *int32) (*[]byte, error)
	UserUpdate(ctx context.Context, req *db.UserUpdateTXParams) (*db.AccountsSchemaUser, error)
	UserPermissionsList(ctx context.Context, req int32) (*[]db.UserPermissionsListRow, error)
}

type AccountsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewAccountsRepo(store db.Store) AccountsRepoInterface {
	errorHandler := map[string]string{
		"roles_role_name_key":          "roleName",
		"users_user_name_key":          "userName",
		"users_user_email_key":         "userEmail",
		"customers_customer_email_key": "customerEmail",
		"customers_customer_phone_key": "customerPhone",
	}
	return &AccountsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
