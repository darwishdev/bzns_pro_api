package gapi

import (
	"testing"
	"time"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	"github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1/rmsv1connect"
	"github.com/darwishdev/bzns_pro_api/common/random"
	"github.com/darwishdev/bzns_pro_api/config"
)

func newTestApi(t *testing.T, store db.Store) rmsv1connect.RmsCoreServiceHandler {
	config := config.Config{
		TokenSymmetricKey:   random.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	api := NewApi(config, store)

	return api
}
