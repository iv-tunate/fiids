-- name: FollowFeeds :one

INSERT INTO feed_follows(user_id, feed_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetFollowedFeeds :many
SELECT * FROM feed_follows
WHERE user_id = $1
ORDER BY created_at;