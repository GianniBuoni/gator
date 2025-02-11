-- name: CreatePost :one
INSERT INTO posts (
  id,
  created_at,
  updated_at,
  title,
  description,
  url,
  published_at,
  feed_id
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 )
ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.*
  FROM posts
  INNER JOIN feed_follows ff
  ON posts.feed_id = ff.feed_id
  INNER JOIN users ON ff.user_id = users.id
  WHERE users.id = $1
  ORDER BY posts.published_at DESC
  LIMIT $2;
