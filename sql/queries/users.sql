-- name: CreateUser :one
INSERT INTO users(id, name, email)
VALUES($1, $2, $3)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET name = COALESCE($1, name), 
    email = COALESCE($2, email)
WHERE id = $3
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING email;
