-- name: InsertTask :one
INSERT INTO tasks (list_id, description) VALUES ($1, $2) RETURNING *;

-- name: SelectTasksForList :many
SELECT * FROM tasks WHERE list_id = $1;

-- name: UpdateTaskSetCompletedAtNow :one
UPDATE tasks SET completed_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;
