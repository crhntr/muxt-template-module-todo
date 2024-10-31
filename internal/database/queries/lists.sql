-- name: InsertList :one
INSERT INTO lists (name, description) VALUES ($1, $2) RETURNING *;

-- name: SelectLists :many
SELECT * FROM lists;

-- name: UpdateList :one
UPDATE lists SET name = $2, description = $3 WHERE id = $1 RETURNING *;

-- name: SelectList :one
SELECT * FROM lists WHERE id = $1;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1;
