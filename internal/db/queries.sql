-- name: GetUserByID :one
SELECT id, name, age
FROM users
WHERE id= $1;

-- name: CreateUser :one

INSERT INTO users (name, age)
VALUES ($1, $2)
RETURNING id, name, age;

-- name: UpdateUser :one
UPDATE users
SET name = $2, age = $3
WHERE id = $1
RETURNING id, name, age;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;