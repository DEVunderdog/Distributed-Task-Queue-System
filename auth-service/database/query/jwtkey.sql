-- name: CreateJWTKey :one
INSERT INTO jwtkeys (
    public_key,
    private_key,
    algorithm,
    is_active,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateJWTKeysActiveness :one
UPDATE jwtkeys
SET
    is_active = sqlc.arg('is_active'),
    updated_at = current_timestamp
WHERE
    public_key = sqlc.arg('public_key')
RETURNING *;

-- name: DeleteJWTKey :exec
DELETE FROM jwtkeys
WHERE public_key = sqlc.arg('public_key')
RETURNING *;