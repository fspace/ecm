package app

import "github.com/spf13/viper"

// modular define methods a module should have
type Module interface {
	// globalViper instance or SubViper?  当前才用拉配置  后期可以用依赖注入 就不需要从viper实例中拉配置信息了
	Configure(v *viper.Viper)

	// Init for initializing some components this module depend , eg: db redisClient mongoClient etc...
	Init()
	// Stop() error
}
