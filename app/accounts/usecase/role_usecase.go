package service

import (
	"context"

	"github.com/bufbuild/connect-go"
<<<<<<< HEAD
	"github.com/darwishdev/bzns_pro_api/common/auth"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
=======
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
)

func (u *AccountsUsecase) RoleCreate(ctx context.Context, req *rmsv1.RoleCreateRequest) (*rmsv1.RoleCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.RoleCreateSqlFromGrpc(req)
	record, err := u.repo.RoleCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.RoleCreateGrpcFromSql(record), nil

}
func (u *AccountsUsecase) RoleFindForUpdate(ctx context.Context, req *rmsv1.RoleFindForUpdateRequest) (*rmsv1.RoleUpdateRequest, error) {
	category, err := u.repo.RoleFindForUpdate(ctx, &req.RoleId)

	if err != nil {
		return nil, err
	}
	res, err := u.adapter.RoleFindForUpdateGrpcFromSql(category)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AccountsUsecase) RoleUpdate(ctx context.Context, req *rmsv1.RoleUpdateRequest) (*rmsv1.RoleUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.RoleUpdateSqlFromGrpc(req)
	record, err := s.repo.RoleUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.RoleUpdateGrpcFromSql(record), nil

}

func (s *AccountsUsecase) PermissionsList(ctx context.Context, req *rmsv1.PermissionsListRequest) (*rmsv1.PermissionsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.PermissionsList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.adapter.PermissionsListGrpcFromSql(record)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return resp, nil

}

func (s *AccountsUsecase) RolesList(ctx context.Context, req *rmsv1.RolesListRequest) (*rmsv1.RolesListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.RolesList(ctx)
	if err != nil {
		return nil, err
	}
	resp := s.adapter.RolesListGrpcFromSql(record)

	return resp, nil
}

func (s *AccountsUsecase) RoleDeleteRestore(ctx context.Context, req *rmsv1.RoleDeleteRestoreRequest) (*rmsv1.RoleDeleteRestoreResponse, error) {
	err := s.repo.RoleDeleteRestore(ctx, req.RoleIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.RoleDeleteRestoreResponse{}, nil
}

func (s *AccountsUsecase) RolesInputList(ctx context.Context, req *rmsv1.RolesInputListRequest) (*rmsv1.RolesInputListResponse, error) {
	roles, err := s.repo.RolesInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.RolesInputListGrpcFromSql(roles)

	return res, nil
}
