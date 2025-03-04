-- name: GetTodosGetByUserId :many
SELECT id, text, done, user_id
FROM todos
WHERE user_id = ?;

-- name: GetUserById :one
SELECT id, name
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUsersById :many
SELECT id, name
FROM users
WHERE id IN (sqlc.slice(userIDs));

-- name: CreateTodo :execresult
INSERT INTO todos(text, user_id) VALUES (?, ?);
