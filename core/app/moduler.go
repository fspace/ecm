package app

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

// 参考：Burrow
// modular define methods a module should have
type Module interface {
	// globalViper instance or SubViper?  当前才用拉配置  后期可以用依赖注入 就不需要从viper实例中拉配置信息了
	Configure(v *viper.Viper) error

	// Init for initializing some components this module depend , eg: db redisClient mongoClient etc...
	// TODO 是否引入依赖 AppContext ？  还是依赖App？              appContext可以作为app的内嵌对象

	// TODO 签名改为Init() error
	Init() error

	// TODO 引入依赖解析  depends   如果一个模块依赖另一个模块那么其他模块必须限于此模块被Init Run

	Start() error

	// Stop is called to stop operation of the coordinator. In this func, the coordinator should call the Stop func for
	// any of its modules, and stop any goroutines that it has started. While it can return an error if there is a
	// problem, the errors are mostly ignored.
	Stop() error
}

// 让模块依赖应用程序这个根对象？
type BaseModule struct {
	App *Application
	ProtoModule
}

type ProtoModule struct {
}

// TODO 全部补全么？

func (m *ProtoModule) Configure(v *viper.Viper) error {
	log.Info("default Module run Configure from BaseModule")
	return nil
}
func (m *ProtoModule) Init() error {
	log.Info("default Module run Init from BaseModule")
	return nil
}

func (m *ProtoModule) Start() error {
	log.Info("default Module run Start from BaseModule")
	return nil
}
func (m *ProtoModule) Stop() error {
	return nil
}
