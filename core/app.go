package core

import "go.uber.org/fx"

type App struct {
	Config    *Config
	Modules   []fx.Option
	Servers   []any
	Providers []any
	Invokers  []any
}

func NewApp(cfg *Config) *App {
	return &App{
		Config: cfg,
	}
}

func (a *App) AddModules(modules ...fx.Option) {
	a.Modules = append(a.Modules, modules...)
}

func (a *App) AddServers(servers ...any) {
	a.Servers = append(a.Servers, servers...)
}

func (a *App) AddProviders(providers ...any) {
	a.Providers = append(a.Providers, providers...)
}

func (a *App) AddInvokers(invokers ...any) {
	a.Invokers = append(a.Invokers, invokers...)
}

func (a *App) Start() {
	// default providers
	opts := []fx.Option{
		fx.Provide(func() *Config {
			return a.Config
		}),
		fx.Provide(NewLogger),
	}

	// custom
	opts = append(opts, a.Modules...)
	opts = append(opts, fx.Provide(a.Providers...))
	opts = append(opts, fx.Provide(a.Servers...))
	opts = append(opts, fx.Invoke(a.Invokers...))

	fx.New(opts...).Run()
}
