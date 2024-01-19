package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	DbErrorParser(err error, errorHandler map[string]string) *connect.Error
	Querier

	// roles
	RoleCreateTX(ctx context.Context, arg RoleCreateTXParams) (RoleCreateTXResult, error)
	RoleUpdateTX(ctx context.Context, arg RoleUpdateTXParams) (RoleUpdateTXResult, error)

	//users
	UserCreateTX(ctx context.Context, arg UserCreateTXParams) (UserCreateTXResult, error)
	UserUpdateTX(ctx context.Context, arg UserUpdateTXParams) (UserUpdateTXResult, error)

	// transactions
}

// Store provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
