-- name: GenerateApiKey :one
INSERT INTO api_keys(name, api_key, user_id)
VALUES($1, $2, $3)
RETURNING *;

