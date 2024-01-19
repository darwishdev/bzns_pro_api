package service

import (
	"context"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/meloneg/mln_rms_core/app/accounts/adapter"
	"github.com/meloneg/mln_rms_core/app/accounts/repo"
	"github.com/meloneg/mln_rms_core/common/auth"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
<<<<<<< HEAD
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
=======
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
	"github.com/meloneg/mln_rms_core/common/redisclient"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
)

type AccountsUsecaseInterface interface {
	RoleCreate(ctx context.Context, req *rmsv1.RoleCreateRequest) (*rmsv1.RoleCreateResponse, error)
	PermissionsList(ctx context.Context, req *rmsv1.PermissionsListRequest) (*rmsv1.PermissionsListResponse, error)
	RoleUpdate(ctx context.Context, req *rmsv1.RoleUpdateRequest) (*rmsv1.RoleUpdateResponse, error)
	RolesList(ctx context.Context, req *rmsv1.RolesListRequest) (*rmsv1.RolesListResponse, error)
	RoleFindForUpdate(ctx context.Context, req *rmsv1.RoleFindForUpdateRequest) (*rmsv1.RoleUpdateRequest, error)
	RoleDeleteRestore(ctx context.Context, req *rmsv1.RoleDeleteRestoreRequest) (*rmsv1.RoleDeleteRestoreResponse, error)
	RolesInputList(ctx context.Context, req *rmsv1.RolesInputListRequest) (*rmsv1.RolesInputListResponse, error)
	// user
	UserLogin(ctx context.Context, req *rmsv1.UserLoginRequest) (*rmsv1.UserLoginResponse, error)
	UserAuthorize(ctx context.Context, req *auth.Payload) (*rmsv1.UserLoginResponse, error)
	UserResetPassword(ctx context.Context, req *rmsv1.UserResetPasswordRequest) error
	UsersList(ctx context.Context, req *rmsv1.UsersListRequest) (*rmsv1.UsersListResponse, error)
	UserDeleteRestore(ctx context.Context, req *rmsv1.UserDeleteRestoreRequest) (*rmsv1.UserDeleteRestoreResponse, error)
	UserCreate(ctx context.Context, req *rmsv1.UserCreateRequest) (*rmsv1.UserCreateResponse, error)
	UserFindForUpdate(ctx context.Context, req *rmsv1.UserFindForUpdateRequest) (*rmsv1.UserUpdateRequest, error)
	UserUpdate(ctx context.Context, req *rmsv1.UserUpdateRequest) (*rmsv1.UserUpdateResponse, error)
<<<<<<< HEAD
=======

	// customer

	AccountsList(ctx context.Context, req *rmsv1.AccountsListRequest) (*rmsv1.AccountsListResponse, error)
	AccountDeleteRestore(ctx context.Context, req *rmsv1.AccountDeleteRestoreRequest) (*rmsv1.AccountDeleteRestoreResponse, error)
	AccountCreate(ctx context.Context, req *rmsv1.AccountCreateRequest) (*rmsv1.AccountCreateResponse, error)
	AccountFindForUpdate(ctx context.Context, req *rmsv1.AccountFindForUpdateRequest) (*rmsv1.AccountUpdateRequest, error)
	AccountUpdate(ctx context.Context, req *rmsv1.AccountUpdateRequest) (*rmsv1.AccountUpdateResponse, error)
	AccountFind(ctx context.Context, req *rmsv1.AccountFindRequest) (*rmsv1.AccountFindResponse, error)
	AccountsInputList(ctx context.Context, req *rmsv1.AccountsInputListRequest) (*rmsv1.AccountsInputListResponse, error)
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
}

type AccountsUsecase struct {
	repo          repo.AccountsRepoInterface
	validator     *protovalidate.Validator
	tokenMaker    auth.Maker
	tokenDuration time.Duration
	adapter       adapter.AccountsAdapterInterface
	redisClient   redisclient.RedisClientInterface
}

func NewAccountsUsecase(store db.Store, validator *protovalidate.Validator, tokenMaker auth.Maker, tokenDuration time.Duration, redisClient redisclient.RedisClientInterface) *AccountsUsecase {
	repo := repo.NewAccountsRepo(store)
	adapter := adapter.NewAccountsAdapter()
	// cache := cache.NewAuthCache(redisClient)
	return &AccountsUsecase{
		repo:          repo,
		tokenMaker:    tokenMaker,
		validator:     validator,
		tokenDuration: tokenDuration,
		adapter:       adapter,
		redisClient:   redisClient,
	}
}
