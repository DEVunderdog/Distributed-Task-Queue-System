// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CheckForExistingUser(ctx context.Context, email string) (bool, error)
	CountJWTKeys(ctx context.Context) (int64, error)
	CreateJWTKey(ctx context.Context, arg CreateJWTKeyParams) (Jwtkey, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerificationCode(ctx context.Context, arg CreateVerificationCodeParams) (VerificationCode, error)
	DeleteJWTKey(ctx context.Context, publicKey string) error
	DeleteSessions(ctx context.Context, id int64) error
	GetAllSessions(ctx context.Context, arg GetAllSessionsParams) ([]Session, error)
	GetLatestJWTKey(ctx context.Context, isActive pgtype.Bool) ([]Jwtkey, error)
	GetSessionByUser(ctx context.Context, arg GetSessionByUserParams) ([]Session, error)
	GetSessionsByActiveness(ctx context.Context, arg GetSessionsByActivenessParams) ([]Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserSessionsByActiveness(ctx context.Context, arg GetUserSessionsByActivenessParams) ([]Session, error)
	LoggedOutSession(ctx context.Context, id int64) (Session, error)
	UpdateJWTKeysActiveness(ctx context.Context, arg UpdateJWTKeysActivenessParams) (Jwtkey, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerificationCodeStatus(ctx context.Context, arg UpdateVerificationCodeStatusParams) (VerificationCode, error)
}

var _ Querier = (*Queries)(nil)
