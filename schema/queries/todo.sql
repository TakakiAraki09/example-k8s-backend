
-- name: GetTodosGetByUserId :many
SELECT id, text, done, user_id
FROM todos
WHERE user_id = ?;

-- name: GetTodosGetAll :many
SELECT id, text, done, user_id
FROM todos;

-- name: GetTodoById :one
SELECT id, text, done, user_id
FROM todos
WHERE id = ?
LIMIT 1;

-- name: CreateTodo :execresult
INSERT INTO todos(text, user_id) VALUES (?, ?);