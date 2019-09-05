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
	// -------------------------------------------------------------------------------------- +|
	// ## 实例化必要的应用程序组件  事件注册等任务

	//  == 应用程序组件存储在Context里面

	// 从配置获取数据库路径 演示组件实例化需要的配置获取   从配置类获取 或者原始方式获取  最好用前者  以后好改为依赖注入！
	dbPath := app.Config.Raw.Get("sqlite_db")
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("连接数据库失败")
	}
	app.DB = db

}

func (app *Application) Run() {
	log.Info("app::run")
	defer app.DB.Close()

	// 顺序运行各模块？
	for mid, mod := range app.modules {
		fmt.Println("run module: ", mid)
		mod.Start() // TODO 错误处理
		// TODO Burrow中 反向停止开启的模块
	}

}

func (app *Application) RegisterModule(mid string, m Module) {
	m.Configure(app.Config.Raw)
	m.Init()

	app.modules[mid] = m
}
