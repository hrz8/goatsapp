package static

import "go.uber.org/fx"

var Module = fx.Module("static", fx.Provide(NewHandler))
