-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBookByID :one
SELECT * FROM books WHERE id = ? ORDER BY id LIMIT 1;

-- name: CreateBook :exec
INSERT INTO books (title, author_id) VALUES (?, ?);

-- name: UpdateBook :exec
UPDATE books SET title = ?, author_id = ? WHERE id = ?;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = ?;