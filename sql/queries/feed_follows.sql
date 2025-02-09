-- name: CreateFeedFollow :one
WITH inserted_feeds_follows AS (
  INSERT INTO feed_follows (
    id,
    created_at,
    updated_at,
    user_id,
    feed_id
  ) 
  VALUES (
    $1, $2, $3, $4, $5
  ) RETURNING *
) 
SELECT 
  inserted_feeds_follows.*, 
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feeds_follows
INNER JOIN feeds ON inserted_feeds_follows.feed_id = feeds.id
INNER JOIN users ON inserted_feeds_follows.user_id = users.id;
