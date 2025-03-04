-- name: ExampleAll :many
SELECT * FROM example_table;

-- name: ExampleOne :one
SELECT * FROM example_table LIMIT 1;

-- name: CreateExampleContent :execresult
INSERT INTO example_table (title, hoge) VALUES (?, ?);
