package gapi

import (
	"testing"
	"time"

	db "github.com/meloneg/mln_rms_core/common/db/gen"
	"github.com/meloneg/mln_rms_core/common/pb/bznspro/v1/rmsv1connect"
	"github.com/meloneg/mln_rms_core/common/random"
	"github.com/meloneg/mln_rms_core/config"
)

func newTestApi(t *testing.T, store db.Store) rmsv1connect.RmsCoreServiceHandler {
	config := config.Config{
		TokenSymmetricKey:   random.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	api := NewApi(config, store)

	return api
}
