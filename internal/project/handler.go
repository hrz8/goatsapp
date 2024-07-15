package project

import (
	"errors"
	"net/http"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/goatsapp/views/component"
	"github.com/hrz8/goatsapp/views/page"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	log *gofx.Logger
	svc *Service
}

func NewHandler(log *gofx.Logger, svc *Service) *Handler {
	return &Handler{log, svc}
}

func (h *Handler) Settings(e echo.Context) error {
	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.RenderView(http.StatusOK, page.ProjectSettings())
}

func (h *Handler) CreateProjectForm(e echo.Context) error {
	var err error

	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !c.IsHtmx() {
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

	p := new(CreateProjectDto)
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if err = c.Validate(p); err != nil {
		return c.RenderView(
			http.StatusUnprocessableEntity,
			component.Toast(component.ToastProps{
				Message: err.Error(),
				Type:    "failed",
			}),
		)
	}

	err = h.svc.CreateProject(c, p)
	if err != nil && errors.Is(err, ErrAliasAlreadyExist) {
		return c.RenderView(
			http.StatusUnprocessableEntity,
			component.Toast(component.ToastProps{
				Message: "Project with alias: " + p.Alias + " already exist",
				Type:    "failed",
			}),
		)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.RenderView(
		http.StatusOK,
		component.Toast(component.ToastProps{
			Message: "Project created, reloading...",
			Type:    "success",
		}),
	)
}

func (h *Handler) ListProjectSelector(e echo.Context) error {
	var err error

	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !c.IsHtmx() {
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

	p := new(ListProjectDto)
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var ck *http.Cookie
	ck, err = c.Cookie("project_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	curr := ck.Value

	var projects []*dbrepo.GetProjectsRow
	projects, err = h.svc.ListProjects(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.RenderView(
		http.StatusOK,
		component.ProjectSelectorList(p.Name, curr, projects),
	)
}

func (h *Handler) SwitchProject(e echo.Context) error {
	var err error

	c, ok := e.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !c.IsHtmx() {
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

	p := new(SwitchProjectDto)
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "project_id"
	cookie.Value = p.ProjectID
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.RenderView(
		http.StatusOK,
		component.Toast(component.ToastProps{
			Message: "Reloading, please wait...",
			Type:    "success",
		}),
	)
}
