-- name: CreateUser :one
INSERT INTO users (
    email,
    hashed_password,
    email_verified
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE(sqlc.narg(email), email),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    email_verified = COALESCE(sqlc.narg(email_verified), email_verified),
    updated_at = current_timestamp
WHERE
    id = sqlc.arg(id)
RETURNING *;