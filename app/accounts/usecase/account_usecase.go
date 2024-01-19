package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
)

func (s *AccountsUsecase) AccountsList(ctx context.Context, req *rmsv1.AccountsListRequest) (*rmsv1.AccountsListResponse, error) {
	record, err := s.repo.AccountsList(ctx, &req.AccountTypeId)
	if err != nil {
		return nil, err
	}
	resp := s.adapter.AccountsListGrpcFromSql(record)

	// resp.Options = authorizedAccount.GetAccessableActionsForGroup("accounts")
	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *AccountsUsecase) AccountsInputList(ctx context.Context, req *rmsv1.AccountsInputListRequest) (*rmsv1.AccountsInputListResponse, error) {
	record, err := s.repo.AccountsInputList(ctx, &req.AccountTypeId)
	if err != nil {
		return nil, err
	}
	resp := s.adapter.AccountsInputListGrpcFromSql(record)

	return resp, nil
}

func (s *AccountsUsecase) AccountDeleteRestore(ctx context.Context, req *rmsv1.AccountDeleteRestoreRequest) (*rmsv1.AccountDeleteRestoreResponse, error) {
	err := s.repo.AccountDeleteRestore(ctx, req.AccountIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.AccountDeleteRestoreResponse{}, nil
}

func (u *AccountsUsecase) AccountCreate(ctx context.Context, req *rmsv1.AccountCreateRequest) (*rmsv1.AccountCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.AccountCreateSqlFromGrpc(req)
	record, err := u.repo.AccountCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.AccountCreateGrpcFromSql(record), nil

}
func (u *AccountsUsecase) AccountFindForUpdate(ctx context.Context, req *rmsv1.AccountFindForUpdateRequest) (*rmsv1.AccountUpdateRequest, error) {
	user, err := u.repo.AccountFindForUpdate(ctx, &req.AccountId)

	if err != nil {
		return nil, err
	}
	resp := u.adapter.AccountFindForUpdateGrpcFromSql(user)

	return resp, nil
}

func (u *AccountsUsecase) AccountFind(ctx context.Context, req *rmsv1.AccountFindRequest) (*rmsv1.AccountFindResponse, error) {
	user, err := u.repo.AccountFind(ctx, &req.AccountId)

	if err != nil {
		return nil, err
	}
	resp := u.adapter.AccountFindGrpcFromSql(user)

	return resp, nil
}

func (s *AccountsUsecase) AccountUpdate(ctx context.Context, req *rmsv1.AccountUpdateRequest) (*rmsv1.AccountUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.AccountUpdateSqlFromGrpc(req)
	record, err := s.repo.AccountUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.AccountUpdateGrpcFromSql(record), nil

}
