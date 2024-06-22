// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbrepo

import (
	"context"
)

type Querier interface {
	CreateNewProjects(ctx context.Context, arg *CreateNewProjectsParams) (int32, error)
	GetProjectByAlias(ctx context.Context, alias string) (*GetProjectByAliasRow, error)
	GetProjects(ctx context.Context) ([]*GetProjectsRow, error)
}

var _ Querier = (*Queries)(nil)
