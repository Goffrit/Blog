-- name: CreateUser :exec
INSERT INTO
    users (
        username,
        password,
        email,
        full_name,
        date_of_birth
    )
VALUES (?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT
    user_id,
    username,
    email,
    full_name,
    date_of_birth,
    created_at,
    updated_at
FROM users
WHERE
    user_id = ?;

-- name: ListUsers :many
SELECT
    user_id,
    username,
    email,
    full_name,
    date_of_birth,
    created_at,
    updated_at
FROM users
LIMIT ?
OFFSET
    ?;

-- name: GetUserByEmail :one
SELECT
    user_id,
    username,
    password,
    email,
    full_name,
FROM users
WHERE
    email = ?