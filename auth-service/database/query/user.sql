-- name: CreateUser :one
INSERT INTO users (
    email,
    hashed_password,
    email_verified
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE
    id = sqlc.arg('id');

-- name: CheckForExistingUser :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = sqlc.arg('email'));

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE
    email = sqlc.arg('email')
LIMIT 1;

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


