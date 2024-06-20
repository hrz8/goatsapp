package viewutil

import (
	"context"
)

func IsHtmxRequest(ctx context.Context) bool {
	htmxRequest := ctx.Value("hx-request")
	return htmxRequest != nil && htmxRequest == true
}
