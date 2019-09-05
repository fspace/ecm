package backend

import (
	"fmt"
	"github.com/fspace/ecm/core/app"
	"github.com/spf13/viper"
	"log"
)

func New(appInst *app.Application) *Bundle {
	// fmt.Println(appInst)
	//return &Bundle{
	//	   //app.ModuleContext{App:app},
	//		app.BaseModule{App:appInst},
	//}
	b := &Bundle{}
	b.App = appInst
	return b
}

type Bundle struct {
	app.BaseModule
}

// TODO 校验配置是否正确！
func (b *Bundle) Configure(viper *viper.Viper) {
	log.Println("hello-bundle::configure")
	v := viper.Sub("hello-bundle")

	confVal := v.Get("someKey")
	fmt.Println("someKey is : ", confVal)
	//fmt.Println(viper.AllKeys())
}

func (b *Bundle) Init() {
	log.Println("hello-bundle::init")
}
