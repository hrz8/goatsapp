package project

import (
	"net/http"

	"github.com/hrz8/goatsapp/web/template/component"
	"github.com/hrz8/gofx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	log *gofx.Logger
}

func NewHandler(log *gofx.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

type CreateProjectDto struct {
	Name        string `form:"name" validate:"required,min=49,max=50"`
	Alias       string `form:"alias" validate:"required,min=5,max=20"`
	WebhookURL  string `form:"webhook_url" validate:"max=255"`
	Description string `form:"description" validate:"max=140"`
}

func (h *Handler) CreateProjectForm(c echo.Context) error {
	var err error

	cc, ok := c.(*gofx.Context)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	p := new(CreateProjectDto)
	if err = cc.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if !cc.IsHtmx() {
		return echo.NewHTTPError(http.StatusMethodNotAllowed)
	}

	if err = c.Validate(p); err != nil {
		return cc.RenderView(
			http.StatusUnprocessableEntity,
			component.Toast(component.ToastProps{
				Message: err.Error(),
				Hidden:  false,
				Type:    "failed",
			}),
		)
	}
	return cc.RenderView(
		http.StatusOK,
		component.Toast(component.ToastProps{
			Message: "Project created successfully!",
			Hidden:  false,
			Type:    "success",
		}),
	)
}
