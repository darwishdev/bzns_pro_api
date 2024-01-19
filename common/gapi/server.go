package gapi

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
<<<<<<< HEAD
	"github.com/meloneg/mln_rms_core/common/pb/bznspro/v1/bznsprov1connect"
=======
	"github.com/meloneg/mln_rms_core/common/pb/rms/v1/rmsv1connect"
	"github.com/meloneg/mln_rms_core/common/redisclient"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	"github.com/meloneg/mln_rms_core/config"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	config config.Config
	store  db.Store
	api    bznsprov1connect.BznsProServiceHandler
}

func NewServer(config config.Config, store db.Store, redisClient redisclient.RedisClientInterface) *Server {
	return &Server{
		config: config,
		store:  store,
		api:    NewApi(config, store, redisClient),
	}
}

func (s Server) Start(addr string) {
	httpServer, err := s.NewGrpcHttpServer(addr)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initialize grpc ,HTTP api server:")
	}
	log.Debug().Str("running on %s", s.config.GRPCServerAddress).Msg("successfully running")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP listen and serve")
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("HTTP shutdown")
	}
}
func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

func (s Server) NewGrpcHttpServer(addr string) (*http.Server, error) {

	mux := http.NewServeMux()
	mux.Handle(
		"/",
		http.RedirectHandler("https://exploremelon.com", http.StatusFound),
	)
	compress1KB := connect.WithCompressMinBytes(1024)

	interceptors := connect.WithInterceptors(GrpcLogger())

	mux.Handle(bznsprov1connect.NewBznsProServiceHandler(
		s.api,
		interceptors,
		compress1KB,
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(bznsprov1connect.BznsProServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(bznsprov1connect.BznsProServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(bznsprov1connect.BznsProServiceName),
		compress1KB,
	))

	srv := &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	return srv, nil
}
