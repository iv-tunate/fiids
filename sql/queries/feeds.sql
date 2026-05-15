-- name: CreateFeed :one
INSERT INTO feeds(name, url, user_id)
VALUES($1, $2, $3)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds
ORDER BY created_at DESC
LIMIT $1 OFFSET $2 ;

-- name: Getfeedsv2 :many
EXPLAIN (ANALYZE, BUFFERS)
SELECT * FROM feeds
WHERE created_at < $1
ORDER BY created_at
Limit $2;

-- CREATE INDEX idx_created_at ON feeds (created_at DESC)