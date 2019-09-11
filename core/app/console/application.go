package console

import (
	"fmt"
	"github.com/fspace/ecm/core/app"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// 控制台 应用程序 和web api 程序略有不同 所以单独设计一个

func New(conf app.Config) *Application {
	inst := &Application{}
	inst.Config = conf
	return inst
}

// Application
type Application struct {
	KingpinApp *kingpin.Application
	app.Application
	modules map[string]app.Module // todo 改为 map[string]*Module 类型
}

func (ap *Application) Init() {
	log.Info("app::init")

	ap.KingpinApp = kingpin.New("ecm-console", "My modular application.")
	ap.KingpinApp.HelpFlag.Short('h')

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

	kingpin.MustParse(ap.KingpinApp.Parse(os.Args[1:]))
}

// Fixme 这里感觉有冗余 是否可以提取到一个单独的结构去？
// 设计一个Modules| 或者 ModuleAgg 聚合对象 然后实现注册方法 运行方法 init方法  然后生命周期挂接到app上
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
