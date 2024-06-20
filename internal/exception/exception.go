package exception

import "go.uber.org/fx"

var Module = fx.Module("exception", fx.Provide(NewHandler))
