package repo

import (
	"context"

	db "github.com/meloneg/mln_rms_core/common/db/gen"
)

func (repo *AccountsRepo) UserFindByEmailOrCode(ctx context.Context, req string) (*db.UserFindByEmailOrCodeRow, error) {
	resp, err := repo.store.UserFindByEmailOrCode(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) UserPermissionsList(ctx context.Context, req int32) (*[]db.UserPermissionsListRow, error) {
	resp, err := repo.store.UserPermissionsList(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) UserFind(ctx context.Context, req int32) (*[]byte, error) {
	resp, err := repo.store.UserFind(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) UsersList(ctx context.Context) (*[][]byte, error) {
	resp, err := repo.store.UsersList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) UserResetPassword(ctx context.Context, req *db.UserResetPasswordParams) error {
	err := repo.store.UserResetPassword(context.Background(), *req)

	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) UserDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.UserDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) UserCreate(ctx context.Context, req *db.UserCreateTXParams) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserCreateTX(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp.User, nil
}

func (repo *AccountsRepo) UserFindForUpdate(ctx context.Context, req *int32) (*[]byte, error) {
	resp, err := repo.store.UserFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) UserUpdate(ctx context.Context, req *db.UserUpdateTXParams) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserUpdateTX(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp.User, nil
}
