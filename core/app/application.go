package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

// Application
// FIXME 可以考虑也实现 Module接口 变为组合设计模式？
type Application struct {
	Context
	Config  Config            // todo 改为引用？
	modules map[string]Module // todo 改为 map[string]*Module 类型
	// GinRouter gin.IRouter
	GinRouter *gin.Engine
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

	// 实例化 gin-router  先绑定到gin框架  按理 这里可以用不同的框架的  可以做包级别适配  比如不同的框架里面都有这个类
	// 然后main入口 顶部切换不同的包即可 使用代码不变
	app.GinRouter = gin.Default()
	buildRouter(app)
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
	panic(app.GinRouter.Run(":8080"))
}

func (app *Application) RegisterModule(mid string, m Module) error {
	err := m.Configure(app.Config.Raw)
	if err != nil {
		return err
	}
	if err = m.Init(); err != nil {
		return err
	}

	app.modules[mid] = m

	return nil
}

// --------------------------------------------------------

func buildRouter(appInst *Application) {
	r := appInst.GinRouter
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
