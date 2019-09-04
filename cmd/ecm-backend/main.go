package main

import (
	"fmt"
	"github.com/fspace/ecm/bundles/hello/delivery/backend"
	"github.com/fspace/ecm/core/app"
)

func main() {
	// ## 加载配置

	// ## 全局组件实例化 如 db log redis ...

	// ## App 应用程序根实例化            此处考虑跟上面组件实例化的关系 是否有依赖  app = { comp1 , comp2 ... }
	// 根据依赖注入思想   App 可看做 组件  和 bundle 的集合 那么实例化顺序 应该是 先部件 在整体

	// ## 各个bundles 实例化  并挂接|注册到 App上

	var cfg *app.Config
	// load application configurations
	cfg, err := app.LoadConfig("./config")
	if err != nil {
		// panic(fmt.Errorf("Invalid application configuration: %s", err))
		fmt.Println("LoadConfig", err)
	}

	// appInstance :=
	// 包被同名变量覆盖？
	// TODO 把结构体cfg改为引用类型 避免复制！
	app := app.New(*cfg) // 创建应用  可以用上一步加载出来的应用配置作为应用程序对象的依赖
	app.Init()

	// 加载各个bundle
	// 注册的时候初始化呢 还是运行时Run 再遍历做初始化？
	app.RegisterModule("hello", backend.New( /** 依赖注入 后期考虑 前期先用pull的方法 拉自己的依赖 */ ))

	app.Run()
}
