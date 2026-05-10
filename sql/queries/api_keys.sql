-- name: GenerateApiKey :one
INSERT INTO api_keys(name, api_key, user_id)
VALUES($1, $2, $3)
RETURNING *;

-- name: CheckApiKey :one
SELECT id, api_key, user_id, revoked_at 
FROM api_keys 
WHERE api_key = $1;

-- name: RevokePiKey :exec
UPDATE api_keys
SET revoked_at = NOW()
WHERE api_key = $1;