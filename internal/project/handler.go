package project

import (
	"net/http"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/goatsapp/web/template/component"
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
				Hidden:  false,
				Type:    "failed",
			}),
		)
	}

	exist := h.svc.IsAliasExist(c, p.Alias)
	if exist {
		return c.RenderView(
			http.StatusUnprocessableEntity,
			component.Toast(component.ToastProps{
				Message: "Project with alias: " + p.Alias + " already exist",
				Hidden:  false,
				Type:    "failed",
			}),
		)
	}

	if err = h.svc.CreateProject(c, p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.RenderView(
		http.StatusOK,
		component.Toast(component.ToastProps{
			Message: "Project created successfully!",
			Hidden:  false,
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
	ck, err = c.Cookie("app_id")
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
