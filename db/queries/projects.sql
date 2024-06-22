-- name: GetProjects :many
SELECT
    id,
    name,
    alias,
    description,
    settings
FROM projects
ORDER BY created_at ASC;

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