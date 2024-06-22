package project

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/gofx"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	log  *gofx.Logger
	repo *dbrepo.Queries
}

func NewService(log *gofx.Logger, repo *dbrepo.Queries) *Service {
	return &Service{log, repo}
}

type Settings struct {
	WebhookURL string `json:"webhook_url"`
}

func (s *Service) IsAliasExist(c *gofx.Context, alias string) bool {
	_, err := s.repo.GetProjectByAlias(c.Request().Context(), alias)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return false
	}

	return true
}

func (s *Service) CreateProject(c *gofx.Context, payload *CreateProjectDto) error {
	evt := "CreateProject"
	var projectSettings []byte
	var err error

	settings := &Settings{WebhookURL: payload.WebhookURL}
	projectSettings, err = json.Marshal(settings)
	if err != nil {
		s.log.JSON.Error(
			"failed when marshaling json",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return err
	}

	data := new(dbrepo.CreateNewProjectsParams)
	data.Name = payload.Name
	data.Alias = payload.Alias
	if payload.Description != "" {
		data.Description = &payload.Description
	}
	data.Settings = projectSettings
	if _, err = s.repo.CreateNewProjects(c.Request().Context(), data); err != nil {
		s.log.JSON.Error(
			"failed when store to db",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return err
	}

	return nil
}

func (s *Service) ListProjects(c *gofx.Context) ([]*dbrepo.GetProjectsRow, error) {
	evt := "ListProjects"
	var result []*dbrepo.GetProjectsRow
	var err error

	result, err = s.repo.GetProjects(c.Request().Context())
	if err != nil {
		s.log.JSON.Error(
			"failed when list projects from db",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return result, err
	}

	return result, nil
}
