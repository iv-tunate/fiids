-- name: CreatePost :one

INSERT INTO posts(title, url, description, feed_id, published_at)
VALUES($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetPostsForUser :many
SELECT posts.* from posts
JOIN feed_follows on posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
limit $2 OFFSET $3;