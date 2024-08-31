-- name: CreateSession :one
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
) RETURNING *;

-- name: GetAllSessions :many
SELECT * FROM sessions
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: GetSessionByUser :many
SELECT * FROM sessions
WHERE user_id = sqlc.arg('user_id')
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: GetSessionsByActiveness :many
SELECT * FROM sessions
WHERE is_active = sqlc.arg('is_active')
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: GetUserSessionsByActiveness :many
SELECT * FROM sessions
WHERE user_id = sqlc.arg('user_id')
AND is_active = sqlc.arg('is_active')
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: LoggedOutSession :one
UPDATE users
SET
    logged_out = current_timestamp,
    updated_at = current_timestamp
WHERE
    id = sqlc.arg('id')
RETURNING *;

-- name: DeleteSessions :exec
DELETE FROM sessions
WHERE id = sqlc.arg('id');