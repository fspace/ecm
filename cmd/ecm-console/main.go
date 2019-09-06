package main

import (
	"fmt"
	helloConsole "github.com/fspace/ecm/bundles/hello/delivery/console"
	"github.com/fspace/ecm/core/app"
	"github.com/fspace/ecm/core/app/console"
)

func main() {

	var cfg *app.Config
	// load appInst configurations
	// 多文件路径 可以实现配置覆盖 通用配置出现在后面 特化的出现在前面
	cfg, err := app.LoadConfig("./config", "./../../config")
	if err != nil {
		// panic(fmt.Errorf("Invalid appInst configuration: %s", err))
		fmt.Println("LoadConfig", err)
	}

	// appInstance :=
	// 包被同名变量覆盖？
	// TODO 把结构体cfg改为引用类型 避免复制！
	appInst := console.New(*cfg) // 创建应用  可以用上一步加载出来的应用配置作为应用程序对象的依赖
	appInst.Init()

	// 加载各个bundle
	// 注册的时候初始化呢 还是运行时Run 再遍历做初始化？
	appInst.RegisterModule("hello", helloConsole.New(
		appInst))

	appInst.Run()
}
