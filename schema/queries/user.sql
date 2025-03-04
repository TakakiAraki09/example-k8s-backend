-- name: GetUserById :one
SELECT id, name
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUsersById :many
SELECT id, name
FROM users
WHERE id IN (sqlc.slice(userIDs));
