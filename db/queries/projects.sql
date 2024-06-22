-- name: GetProjects :many
SELECT
    id,
    encode(digest(id::text || alias, 'sha256'), 'hex') as encoded_id,
    name,
    alias,
    description,
    settings
FROM projects
ORDER BY created_at DESC;

-- name: GetDefaultProject :one
SELECT
    id,
    encode(digest(id::text || alias, 'sha256'), 'hex') as encoded_id,
    name,
    alias,
    description,
    settings
FROM projects
ORDER BY created_at DESC
LIMIT 1;

-- name: GetProjectByEncodedID :one
SELECT
    id,
    name,
    alias,
    description,
    settings
FROM projects
WHERE encode(digest(id::text || alias, 'sha256'), 'hex') = $1;

-- name: GetProjectByAlias :one
SELECT id, alias FROM projects WHERE alias = $1;

-- name: CreateNewProjects :one
INSERT INTO
    projects (
        name,
        alias,
        description,
        settings
    )
VALUES ($1, $2, $3, $4) RETURNING id;