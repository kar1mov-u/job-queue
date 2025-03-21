-- name: CreateTask :one
INSERT INTO tasks(
    status) VALUES($1)
    RETURNING id;


-- name: GetTasks :many
SELECT id, created_at, completed_at, status, link
FROM tasks
WHERE created_at BETWEEN $1 AND $2
ORDER BY created_at ASC;

-- name: UpdateStatus :exec
UPDATE tasks SET status = $1 WHERE id=$2;

-- name: CompleteTask :exec
UPDATE tasks SET status = $1, completed_at = NOW(), link = $2 WHERE  id=$3;

-- name: GetTask :one 
SELECT * FROM tasks WHERE id=$1;