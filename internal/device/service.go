package device

import (
	"log/slog"
	"net/http"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/gofx"
)

type Service struct {
	log  *gofx.Logger
	repo *dbrepo.Queries
}

func NewService(log *gofx.Logger, repo *dbrepo.Queries) *Service {
	return &Service{log, repo}
}

func (s *Service) DecodeProjectID(c *gofx.Context, encodedID string) int32 {
	project, _ := s.repo.GetProjectByEncodedID(c.Request().Context(), encodedID)
	return project.ID
}

func (s *Service) CreateDevice(c *gofx.Context, payload *CreateDeviceDto) error {
	evt := "CreateDevice"
	var err error
	var cookie *http.Cookie

	cookie, err = c.Cookie("project_id")
	if err != nil {
		s.log.JSON.Error(
			"failed to get project_id cookie",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return ErrProjectIDNotFound
	}

	project, _ := s.repo.GetProjectByEncodedID(c.Request().Context(), cookie.Value)

	data := new(dbrepo.CreateNewDeviceParams)
	data.ProjectID = project.ID
	data.Name = payload.Name
	data.Description = &payload.Description
	data.ClientDeviceID = payload.ClientDeviceID
	data.CustomAttributes = []byte(payload.CustomAttributes)
	if _, err = s.repo.CreateNewDevice(c.Request().Context(), data); err != nil {
		s.log.JSON.Error(
			"failed when store to db",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return ErrStoreDevice
	}

	return nil
}
