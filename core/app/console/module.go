package console

import "github.com/fspace/ecm/core/app"

// 让模块依赖 控制台应用程序这个根对象？
type BaseModule struct {
	App *Application
	app.ProtoModule
}
