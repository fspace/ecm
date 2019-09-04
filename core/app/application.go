package app

import "github.com/prometheus/common/log"

type Application struct {
	Config  Config
	modules map[string]Module
}

func New(conf Config) *Application {
	return &Application{
		Config: conf,
	}
}

func (app *Application) Init() {
	log.Info("app::init")

	app.modules = make(map[string]Module)
}

func (app *Application) Run() {
	log.Info("app::run")
}

func (app *Application) RegisterModule(mid string, m Module) {
	m.Configure(app.Config.Raw)
	m.Init()

	app.modules[mid] = m
}
