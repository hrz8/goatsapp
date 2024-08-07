// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: projects.sql

package dbrepo

import (
	"context"
)

const createNewProjects = `-- name: CreateNewProjects :one
INSERT INTO
    projects (
        name,
        alias,
        description,
        settings
    )
VALUES ($1, $2, $3, $4) RETURNING id
`

type CreateNewProjectsParams struct {
	Name        string  `db:"name" json:"name"`
	Alias       string  `db:"alias" json:"alias"`
	Description *string `db:"description" json:"description"`
	Settings    []byte  `db:"settings" json:"settings"`
}

func (q *Queries) CreateNewProjects(ctx context.Context, arg *CreateNewProjectsParams) (int32, error) {
	row := q.db.QueryRow(ctx, createNewProjects,
		arg.Name,
		arg.Alias,
		arg.Description,
		arg.Settings,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getDefaultProject = `-- name: GetDefaultProject :one
SELECT
    id,
    encode(digest(id::text || alias, 'sha256'), 'hex') as encoded_id,
    name,
    alias,
    description,
    settings
FROM projects
ORDER BY created_at DESC
LIMIT 1
`

type GetDefaultProjectRow struct {
	ID          int32   `db:"id" json:"id"`
	EncodedID   string  `db:"encoded_id" json:"encoded_id"`
	Name        string  `db:"name" json:"name"`
	Alias       string  `db:"alias" json:"alias"`
	Description *string `db:"description" json:"description"`
	Settings    []byte  `db:"settings" json:"settings"`
}

func (q *Queries) GetDefaultProject(ctx context.Context) (*GetDefaultProjectRow, error) {
	row := q.db.QueryRow(ctx, getDefaultProject)
	var i GetDefaultProjectRow
	err := row.Scan(
		&i.ID,
		&i.EncodedID,
		&i.Name,
		&i.Alias,
		&i.Description,
		&i.Settings,
	)
	return &i, err
}

const getProjectByAlias = `-- name: GetProjectByAlias :one
SELECT id, alias FROM projects WHERE alias = $1
`

type GetProjectByAliasRow struct {
	ID    int32  `db:"id" json:"id"`
	Alias string `db:"alias" json:"alias"`
}

func (q *Queries) GetProjectByAlias(ctx context.Context, alias string) (*GetProjectByAliasRow, error) {
	row := q.db.QueryRow(ctx, getProjectByAlias, alias)
	var i GetProjectByAliasRow
	err := row.Scan(&i.ID, &i.Alias)
	return &i, err
}

const getProjectByEncodedID = `-- name: GetProjectByEncodedID :one
SELECT
    id,
    name,
    alias,
    description,
    settings
FROM projects
WHERE encode(digest(id::text || alias, 'sha256'), 'hex') = $1::text
`

type GetProjectByEncodedIDRow struct {
	ID          int32   `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Alias       string  `db:"alias" json:"alias"`
	Description *string `db:"description" json:"description"`
	Settings    []byte  `db:"settings" json:"settings"`
}

func (q *Queries) GetProjectByEncodedID(ctx context.Context, dollar_1 string) (*GetProjectByEncodedIDRow, error) {
	row := q.db.QueryRow(ctx, getProjectByEncodedID, dollar_1)
	var i GetProjectByEncodedIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Alias,
		&i.Description,
		&i.Settings,
	)
	return &i, err
}

const getProjects = `-- name: GetProjects :many
SELECT
    id,
    encode(digest(id::text || alias, 'sha256'), 'hex') as encoded_id,
    name,
    alias,
    description,
    settings
FROM projects
ORDER BY created_at DESC
`

type GetProjectsRow struct {
	ID          int32   `db:"id" json:"id"`
	EncodedID   string  `db:"encoded_id" json:"encoded_id"`
	Name        string  `db:"name" json:"name"`
	Alias       string  `db:"alias" json:"alias"`
	Description *string `db:"description" json:"description"`
	Settings    []byte  `db:"settings" json:"settings"`
}

func (q *Queries) GetProjects(ctx context.Context) ([]*GetProjectsRow, error) {
	rows, err := q.db.Query(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetProjectsRow
	for rows.Next() {
		var i GetProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.EncodedID,
			&i.Name,
			&i.Alias,
			&i.Description,
			&i.Settings,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
