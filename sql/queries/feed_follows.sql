-- name: FollowFeeds :one

INSERT INTO feed_follows(user_id, feed_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetFollowedFeeds :many
SELECT * FROM feed_follows
WHERE user_id = $1
ORDER BY created_at;

-- name: DeleteFollowedFeeds :exec
DELETE FROM feed_follows
WHERE id = $1 AND user_id = $2;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
