-- name: GetAuthors :many
SELECT * FROM authors;

-- name: GetAuthorByID :one
SELECT * FROM authors WHERE id = ? ORDER BY id LIMIT 1;

-- name: GetAuthorByName :one
SELECT * FROM authors WHERE name = ? ORDER BY id LIMIT 1;

-- name: AuthorExists :one
SELECT EXISTS(SELECT 1 FROM authors WHERE name = ?);

-- name: CreateAuthor :exec
INSERT INTO authors (name, bio) VALUES (?, ?);

-- name: UpdateAuthor :exec
UPDATE authors SET name = ?, bio = ? WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;