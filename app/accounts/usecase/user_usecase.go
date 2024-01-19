package service

import (
	"context"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/meloneg/mln_rms_core/common/auth"
<<<<<<< HEAD
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
=======
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
	"github.com/meloneg/mln_rms_core/common/redisclient"
	"github.com/rs/zerolog/log"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
)

func (u *AccountsUsecase) UserLogin(ctx context.Context, req *rmsv1.UserLoginRequest) (*rmsv1.UserLoginResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	// params := u.adapter.UserLoginSqlFromGrpc(req)
	record, err := u.repo.UserFindByEmailOrCode(ctx, strings.TrimSpace(req.UserEmailOrCode))
	if err != nil {
		return nil, err
	}

	resp, err := u.adapter.UserLoginGrpcFromSql(record)
	if err != nil {
		return nil, err
	}

	// permissions, err := u.repo.UserPermissionsList(ctx, resp.UserId)
	// if err != nil {
	// 	return nil, err
	// }
	permissionsMap, err := u.adapter.UsersPermissionsMapFromSql(record.PermissionGroups)
	if err != nil {
		return nil, err
	}

	userPyaload := auth.UserPayload{
<<<<<<< HEAD
		Username:    resp.UserEmail,
		UserId:      resp.UserId,
		EntityId:    resp.EntityId,
		DeviceId:    req.DeviceId,
		SessionId:   resp.SessionId,
		Permissions: *permissionsMap,
		Duration:    u.tokenDuration,
=======
		Username: resp.UserName,
		UserId:   resp.UserId,
		Duration: u.tokenDuration,
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	}
	accessToken, accessPayload, err := u.tokenMaker.CreateToken(userPyaload)
	if err != nil {
		return nil, err
	}

	authSession := &redisclient.AuthSession{
		UserName:    resp.UserName,
		AccountCode: record.AccountCode,
		SideBar:     resp.SideBar,
		Permissions: *permissionsMap,
		UserID:      resp.UserId,
		EntityID:    record.EntityID,
		SessionID:   int32(record.SessionID),
		AccountID:   record.AccountID,
		DeviceID:    req.DeviceId,
	}
	log.Debug().Interface("authSessionLogin", authSession).Msg("authSessionLogin")

	err = u.redisClient.AuthSessionCreate(ctx, authSession)
	if err != nil {
		log.Debug().Interface("error  from          redos", err).Msg("ASd")
		return nil, err
	}
	loginInfo := &rmsv1.LoginInfo{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt.Format("2006-01-02 15:04:05"),
	}
	resp.LoginInfo = loginInfo
	return resp, nil
}

func (u *AccountsUsecase) UserAuthorize(ctx context.Context, req *auth.Payload) (*rmsv1.UserLoginResponse, error) {
	// record, err := u.repo.UserFind(ctx, req.UserId)
	// if err != nil {
	// 	return nil, err
	// }

	// resp, err := u.adapter.UserLoginGrpcFromSql(record)
	// if err != nil {
	// 	return nil, err
	// }
	return &rmsv1.UserLoginResponse{}, nil
}

func (u *AccountsUsecase) UserResetPassword(ctx context.Context, req *rmsv1.UserResetPasswordRequest) error {
	params := u.adapter.UserResetPasswordSqlFromGrpc(req)
	err := u.repo.UserResetPassword(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountsUsecase) UsersList(ctx context.Context, req *rmsv1.UsersListRequest) (*rmsv1.UsersListResponse, error) {
	record, err := s.repo.UsersList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.adapter.UsersListGrpcFromSql(*record)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AccountsUsecase) UserDeleteRestore(ctx context.Context, req *rmsv1.UserDeleteRestoreRequest) (*rmsv1.UserDeleteRestoreResponse, error) {
	err := s.repo.UserDeleteRestore(ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.UserDeleteRestoreResponse{}, nil
}

func (u *AccountsUsecase) UserCreate(ctx context.Context, req *rmsv1.UserCreateRequest) (*rmsv1.UserCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.UserCreateSqlFromGrpc(req)
	record, err := u.repo.UserCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.UserCreateGrpcFromSql(record), nil

}
func (u *AccountsUsecase) UserFindForUpdate(ctx context.Context, req *rmsv1.UserFindForUpdateRequest) (*rmsv1.UserUpdateRequest, error) {
	user, err := u.repo.UserFindForUpdate(ctx, &req.UserId)

	if err != nil {
		return nil, err
	}
	res, err := u.adapter.UserFindForUpdateGrpcFromSql(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AccountsUsecase) UserUpdate(ctx context.Context, req *rmsv1.UserUpdateRequest) (*rmsv1.UserUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.UserUpdateSqlFromGrpc(req)
	record, err := s.repo.UserUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.UserUpdateGrpcFromSql(record), nil

}
