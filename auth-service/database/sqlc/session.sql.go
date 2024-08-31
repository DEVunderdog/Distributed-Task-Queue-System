// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: session.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    token,
    refresh_token,
    token_expires_at,
    refresh_token_expires_at,
    is_active,
    ip,
    user_agent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, is_active, ip, user_agent, logged_out, created_at, updated_at
`

type CreateSessionParams struct {
	UserID                int64              `json:"user_id"`
	Token                 string             `json:"token"`
	RefreshToken          string             `json:"refresh_token"`
	TokenExpiresAt        pgtype.Timestamptz `json:"token_expires_at"`
	RefreshTokenExpiresAt pgtype.Timestamptz `json:"refresh_token_expires_at"`
	IsActive              pgtype.Bool        `json:"is_active"`
	Ip                    pgtype.Text        `json:"ip"`
	UserAgent             pgtype.Text        `json:"user_agent"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.UserID,
		arg.Token,
		arg.RefreshToken,
		arg.TokenExpiresAt,
		arg.RefreshTokenExpiresAt,
		arg.IsActive,
		arg.Ip,
		arg.UserAgent,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.RefreshToken,
		&i.TokenExpiresAt,
		&i.RefreshTokenExpiresAt,
		&i.IsActive,
		&i.Ip,
		&i.UserAgent,
		&i.LoggedOut,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSessions = `-- name: DeleteSessions :exec
DELETE FROM sessions
WHERE id = $1
`

func (q *Queries) DeleteSessions(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSessions, id)
	return err
}

const getAllSessions = `-- name: GetAllSessions :many
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, is_active, ip, user_agent, logged_out, created_at, updated_at FROM sessions
LIMIT $2
OFFSET $1
`

type GetAllSessionsParams struct {
	Offset pgtype.Int4 `json:"offset"`
	Limit  pgtype.Int4 `json:"limit"`
}

func (q *Queries) GetAllSessions(ctx context.Context, arg GetAllSessionsParams) ([]Session, error) {
	rows, err := q.db.Query(ctx, getAllSessions, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.RefreshToken,
			&i.TokenExpiresAt,
			&i.RefreshTokenExpiresAt,
			&i.IsActive,
			&i.Ip,
			&i.UserAgent,
			&i.LoggedOut,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSessionByUser = `-- name: GetSessionByUser :many
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, is_active, ip, user_agent, logged_out, created_at, updated_at FROM sessions
WHERE user_id = $1
LIMIT $3
OFFSET $2
`

type GetSessionByUserParams struct {
	UserID int64       `json:"user_id"`
	Offset pgtype.Int4 `json:"offset"`
	Limit  pgtype.Int4 `json:"limit"`
}

func (q *Queries) GetSessionByUser(ctx context.Context, arg GetSessionByUserParams) ([]Session, error) {
	rows, err := q.db.Query(ctx, getSessionByUser, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.RefreshToken,
			&i.TokenExpiresAt,
			&i.RefreshTokenExpiresAt,
			&i.IsActive,
			&i.Ip,
			&i.UserAgent,
			&i.LoggedOut,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSessionsByActiveness = `-- name: GetSessionsByActiveness :many
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, is_active, ip, user_agent, logged_out, created_at, updated_at FROM sessions
WHERE is_active = $1
LIMIT $3
OFFSET $2
`

type GetSessionsByActivenessParams struct {
	IsActive pgtype.Bool `json:"is_active"`
	Offset   pgtype.Int4 `json:"offset"`
	Limit    pgtype.Int4 `json:"limit"`
}

func (q *Queries) GetSessionsByActiveness(ctx context.Context, arg GetSessionsByActivenessParams) ([]Session, error) {
	rows, err := q.db.Query(ctx, getSessionsByActiveness, arg.IsActive, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.RefreshToken,
			&i.TokenExpiresAt,
			&i.RefreshTokenExpiresAt,
			&i.IsActive,
			&i.Ip,
			&i.UserAgent,
			&i.LoggedOut,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserSessionsByActiveness = `-- name: GetUserSessionsByActiveness :many
SELECT id, user_id, token, refresh_token, token_expires_at, refresh_token_expires_at, is_active, ip, user_agent, logged_out, created_at, updated_at FROM sessions
WHERE user_id = $1
AND is_active = $2
LIMIT $4
OFFSET $3
`

type GetUserSessionsByActivenessParams struct {
	UserID   int64       `json:"user_id"`
	IsActive pgtype.Bool `json:"is_active"`
	Offset   pgtype.Int4 `json:"offset"`
	Limit    pgtype.Int4 `json:"limit"`
}

func (q *Queries) GetUserSessionsByActiveness(ctx context.Context, arg GetUserSessionsByActivenessParams) ([]Session, error) {
	rows, err := q.db.Query(ctx, getUserSessionsByActiveness,
		arg.UserID,
		arg.IsActive,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.RefreshToken,
			&i.TokenExpiresAt,
			&i.RefreshTokenExpiresAt,
			&i.IsActive,
			&i.Ip,
			&i.UserAgent,
			&i.LoggedOut,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const loggedOutSession = `-- name: LoggedOutSession :one
UPDATE users
SET
    logged_out = current_timestamp,
    updated_at = current_timestamp
WHERE
    id = $1
RETURNING id, email, hashed_password, email_verified, created_at, updated_at
`

func (q *Queries) LoggedOutSession(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, loggedOutSession, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
