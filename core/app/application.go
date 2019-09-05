package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

type Application struct {
	Context
	Config  Config            // todo 改为引用？
	modules map[string]Module // todo 改为 map[string]*Module 类型
}

func New(conf Config) *Application {
	return &Application{
		Config: conf,
	}
}

func (app *Application) Init() {
	log.Info("app::init")

	app.modules = make(map[string]Module)
	// 实例化必要的应用程序组件  事件注册等任务
	//  == 应用程序组件存储在Context里面
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	app.DB = db

}

func (app *Application) Run() {
	log.Info("app::run")
	defer app.DB.Close()

	// 运行其他模块？
	for mid, mod := range app.modules {
		fmt.Println("run module: ", mid)
		mod.Start() // TODO 错误处理
	}

}

func (app *Application) RegisterModule(mid string, m Module) {
	m.Configure(app.Config.Raw)
	m.Init()

	app.modules[mid] = m
}
