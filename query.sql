-- name: GetLog :one
SELECT * FROM logs
WHERE client_ip = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM logs
ORDER BY created_at;