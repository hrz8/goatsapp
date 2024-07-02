-- name: CreateNewDevice :one
INSERT INTO
    devices (
        project_id,
        name,
        description,
        client_device_id,
        custom_attributes
    )
VALUES ($1, $2, $3, $4, $5) RETURNING id;