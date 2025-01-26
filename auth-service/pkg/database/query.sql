-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY first_name;

-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name, password, user_active
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET email = ?,
    first_name = ?,
    last_name = ?,
    password = ?,
    user_active = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
