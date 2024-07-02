package device

import "go.uber.org/fx"

var Module = fx.Module("device", fx.Provide(NewHandler), fx.Provide(NewService))
