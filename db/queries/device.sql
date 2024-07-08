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

-- name: GetDevicesByProjectEncodedID :many
SELECT
    id,
    client_device_id,
    name,
    phone_number,
    jid,
    description,
    is_active
FROM devices
WHERE
    project_id = $1
ORDER BY created_at DESC;