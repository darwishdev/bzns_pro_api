package redisclient

import (
	"context"
	"encoding/json"

	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/rms/v1"
)

type AuthSession struct {
	UserName    string
	AccountCode string
	SideBar     []*rmsv1.SideBarItem
	Permissions map[string]map[string]bool
	UserID      int32
	EntityID    int32
	SessionID   int32
	AccountID   int32
	DeviceID    int32
}

func (r *RedisClient) AuthSessionCreate(ctx context.Context, req *AuthSession) error {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, req.UserName, jsonBytes, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) AuthSessionSetSessionId(ctx context.Context, req *AuthSession) (*AuthSession, error) {
	authSession, err := r.AuthSessionFind(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	authSession.SessionID = req.SessionID

	err = r.AuthSessionCreate(ctx, authSession)
	if err != nil {
		return nil, err
	}
	return authSession, nil
}

func (r *RedisClient) AuthSessionFind(ctx context.Context, username string) (*AuthSession, error) {

	var parsedStruct AuthSession
	jsonBytes, err := r.client.Get(ctx, username).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &parsedStruct)
	if err != nil {
		return nil, err
	}

	return &parsedStruct, nil
}
