-- name: CreateUser :one
INSERT INTO users (id, name) VALUES ($1, $2) RETURNING id, name, created_at, updated_at;