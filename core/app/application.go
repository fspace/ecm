package app

import "github.com/prometheus/common/log"

type Application struct {
	// Config appConfig
}

func New( /* conf  AppConfig */ ) *Application {
	return &Application{}
}

func (app *Application) Init() {
	log.Info("app::init")

}

func (app *Application) Run() {
	log.Info("app::run")
}
