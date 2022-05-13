-- name: GetUser :one
SELECT *
FROM "user"
WHERE id = $1;

-- name: GetUsers :many
SELECT (id, username, email, role)
FROM "user";

-- name: CreateUser :one
INSERT INTO "user"(
  id, username, password, email, user_role
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET password = $1, email = $2
WHERE id =  $3;
