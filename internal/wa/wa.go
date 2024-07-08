package wa

import "go.uber.org/fx"

var Module = fx.Module("wa", fx.Provide(NewEvent))
