-- name: CreateVerificationCode :one
INSERT INTO verification_codes (
    user_id,
    code,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateVerificationCodeStatus :one
UPDATE verification_codes
SET
    is_used = sqlc.arg('is_used'),
    updated_at = current_timestamp
WHERE id = $1
RETURNING * ;