package viewutil

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func RenderView(ctx context.Context, w http.ResponseWriter, component templ.Component) error {
	return component.Render(ctx, w)
}

func IsHtmxRequest(ctx context.Context) bool {
	htmxRequest := ctx.Value("hx-request")
	return htmxRequest != nil && htmxRequest == true
}
