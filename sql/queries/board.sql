-- name: CreateBoard :one
INSERT INTO boards (data)
VALUES ($1)
RETURNING id;

-- name: GetBoard :one
SELECT id, data FROM boards
WHERE id = $1;
