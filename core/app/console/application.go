package console

import (
	"fmt"
	"github.com/fspace/ecm/core/app"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

// 控制台 应用程序 和web api 程序略有不同 所以单独设计一个

func New(conf app.Config) *Application {
	inst := &Application{}
	inst.Config = conf
	return inst
}

// Application
type Application struct {
	app.Application
	modules map[string]app.Module // todo 改为 map[string]*Module 类型
}

func (ap *Application) Init() {
	log.Info("app::init")

	ap.modules = make(map[string]app.Module)
	// -------------------------------------------------------------------------------------- +|
	// ## 实例化必要的应用程序组件  事件注册等任务

	//  == 应用程序组件存储在Context里面

	// 从配置获取数据库路径 演示组件实例化需要的配置获取   从配置类获取 或者原始方式获取  最好用前者  以后好改为依赖注入！
	//dbPath := ap.Config.Raw.Get("sqlite_db")
	//db, err := gorm.Open("sqlite3", dbPath)
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	ap.DB = db

}

func (ap *Application) Run() {
	log.Info("app::run")
	defer ap.DB.Close()

	// 顺序运行各模块？
	for mid, mod := range ap.modules {
		fmt.Println("run module: ", mid)
		mod.Start() // TODO 错误处理
		// TODO Burrow中 反向停止开启的模块
	}

}

func (ap *Application) RegisterModule(mid string, m app.Module) error {
	err := m.Configure(ap.Config.Raw)
	if err != nil {
		return err
	}
	if err = m.Init(); err != nil {
		return err
	}

	ap.modules[mid] = m

	return nil
}
