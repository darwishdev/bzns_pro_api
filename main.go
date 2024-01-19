package main

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	"github.com/darwishdev/bzns_pro_api/common/gapi"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
	"github.com/darwishdev/bzns_pro_api/config"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	ctx := context.Background()
	state, err := config.LoadState("./config")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load the state config")
	}
	config, err := config.LoadConfig("./config", state.State)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load the config")
	}
	store, err := db.InitDB(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	rc := redisclient.NewRedisClient(config.RedisHost, config.RedisPort, config.RedisPassword, config.RedisDatabase)

	server := gapi.NewServer(config, store, rc)
	server.Start(config.GRPCServerAddress)
}
