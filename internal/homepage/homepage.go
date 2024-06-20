package homepage

import "go.uber.org/fx"

var Module = fx.Module("homepage", fx.Provide(NewHandler))
