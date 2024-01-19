package db

import (
	"context"

	"github.com/rs/zerolog/log"
)

type UserCreateTXParams struct {
	UserParams        UserCreateParams
	PermissionsParams []UserPermissionsBulkCreateParams
	RolesParams       []UserRolesBulkCreateParams
}

type UserCreateTXResult struct {
	User AccountsSchemaUser
}

func (store *SQLStore) UserCreateTX(ctx context.Context, arg UserCreateTXParams) (UserCreateTXResult, error) {
	var result UserCreateTXResult

	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		user, err := q.UserCreate(ctx, arg.UserParams)
		if err != nil {
			return err
		}

		for i := 0; i < len(arg.PermissionsParams); i++ {
			arg.PermissionsParams[i].UserID = user.UserID
		}
		_, err = q.UserPermissionsBulkCreate(ctx, arg.PermissionsParams)
		if err != nil {
			return err
		}

		for i := 0; i < len(arg.RolesParams); i++ {
			arg.RolesParams[i].UserID = user.UserID
		}
		_, err = q.UserRolesBulkCreate(ctx, arg.RolesParams)
		if err != nil {
			return err
		}

		result.User = user

		return err
	})

	return result, err
}

type UserUpdateTXParams struct {
	UserParams        UserUpdateParams
	PermissionsParams []UserPermissionsBulkCreateParams
	RolesParams       []UserRolesBulkCreateParams
}

type UserUpdateTXResult struct {
	User AccountsSchemaUser
}

func (store *SQLStore) UserUpdateTX(ctx context.Context, arg UserUpdateTXParams) (UserUpdateTXResult, error) {
	var result UserUpdateTXResult

	log.Debug().Interface("args", arg).Msg("from tx")
	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		user, err := q.UserUpdate(ctx, arg.UserParams)
		if err != nil {
			log.Debug().Interface("err", err).Msg("from UserUpdate")

			return err
		}
		err = q.UserPermissionsClear(ctx, arg.UserParams.UserID)
		if err != nil {
			return err
		}

		for i := 0; i < len(arg.PermissionsParams); i++ {
			arg.PermissionsParams[i].UserID = user.UserID
		}
		_, err = q.UserPermissionsBulkCreate(ctx, arg.PermissionsParams)
		if err != nil {
			return err
		}

		err = q.UserRolesClear(ctx, arg.UserParams.UserID)
		if err != nil {
			return err
		}
		for i := 0; i < len(arg.RolesParams); i++ {
			arg.RolesParams[i].UserID = user.UserID
		}
		_, err = q.UserRolesBulkCreate(ctx, arg.RolesParams)
		if err != nil {
			return err
		}

		result.User = user

		return err
	})

	return result, err
}
