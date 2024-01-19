package gapi

import (
	"github.com/bufbuild/protovalidate-go"
	accountsUsecase "github.com/darwishdev/bzns_pro_api/app/accounts/usecase"
	eventsUsecase "github.com/darwishdev/bzns_pro_api/app/events/usecase"

	publicUsecase "github.com/darwishdev/bzns_pro_api/app/public/usecase"
	"github.com/darwishdev/bzns_pro_api/common/auth"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
<<<<<<< HEAD
	"github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1/bznsprov1connect"

=======
	"github.com/darwishdev/bzns_pro_api/common/pb/rms/v1/rmsv1connect"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	"github.com/darwishdev/bzns_pro_api/config"
)

// Server serves gRPC requests for our banking usecase.
type Api struct {
	bznsprov1connect.UnimplementedBznsProServiceHandler
	config          config.Config
	tokenMaker      auth.Maker
	accountsUsecase accountsUsecase.AccountsUsecaseInterface
<<<<<<< HEAD
	eventsUsecase   eventsUsecase.EventsUsecaseInterface

	publicUsecase publicUsecase.PublicUsecaseInterface
	store         db.Store
}

// NewServer creates a new gRPC server.
func NewApi(config config.Config, store db.Store) bznsprov1connect.BznsProServiceHandler {
=======
	ordersUsecase   ordersUsecase.OrdersUsecaseInterface
	entitiesUsecase entitiesUsecase.EntitiesUsecaseInterface
	publicUsecase   publicUsecase.PublicUsecaseInterface
	sessionUsecase  sessionUsecase.SessionsUsecaseInterface
	placesUsecase   placesUsecase.PlacesUsecaseInterface
	stockUsecase    stockUsecase.StockUsecaseInterface
	productUsecase  productUsecase.ProductsUsecaseInterface
	store           db.Store
	redisClient     redisclient.RedisClientInterface
}

// NewServer creates a new gRPC server.
func NewApi(config config.Config, store db.Store, redisClient redisclient.RedisClientInterface) rmsv1connect.RmsCoreServiceHandler {
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	validator, err := protovalidate.New()

	if err != nil {
		panic("cann't create validator in gapi/api.go")
	}
	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		panic("cann't create paset maker in gapi/api.go")
	}
<<<<<<< HEAD
	accountsUsecase := accountsUsecase.NewAccountsUsecase(store, validator, tokenMaker, config.AccessTokenDuration)
	eventsUsecase := eventsUsecase.NewEventsUsecase(store, validator)

	publicUsecase := publicUsecase.NewPublicUsecase(store, validator)
=======
	accountsUsecase := accountsUsecase.NewAccountsUsecase(store, validator, tokenMaker, config.AccessTokenDuration, redisClient)
	placesUsecase := placesUsecase.NewPlacesUsecase(store, validator)
	productUsecase := productUsecase.NewProductsUsecase(store, validator)
	entitiesUsecase := entitiesUsecase.NewEntitiesUsecase(store, validator)
	ordersUsecase := ordersUsecase.NewOrdersUsecase(store, validator)
	publicUsecase := publicUsecase.NewPublicUsecase(store, validator)
	sessionUsecase := sessionUsecase.NewSessionsUsecase(store, validator, tokenMaker, redisClient)
	stockUsecase := stockUsecase.NewStockUsecase(store, validator)
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	return &Api{
		config:          config,
		tokenMaker:      tokenMaker,
		store:           store,
		accountsUsecase: accountsUsecase,
		eventsUsecase:   eventsUsecase,
		publicUsecase:   publicUsecase,
<<<<<<< HEAD
=======
		redisClient:     redisClient,
		entitiesUsecase: entitiesUsecase,
		placesUsecase:   placesUsecase,
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
	}
}
