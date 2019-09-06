package backend

import (
	"fmt"
	"github.com/fspace/ecm/bundles/hello"
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
	Config *hello.Config
}

// TODO 校验配置是否正确！
func (b *Bundle) Configure(viper *viper.Viper) error {
	log.Println("hello-bundle::configure")
	v := viper.Sub("hello-bundle")

	confVal := v.Get("someKey")
	fmt.Println("someKey is : ", confVal)
	//fmt.Println(viper.AllKeys())
	conf := hello.Config{}
	v.Unmarshal(&conf)
	fmt.Printf("config is : %#v \n", conf)
	b.Config = &conf

	return nil
}

func (b *Bundle) Init() error {
	log.Println("hello-bundle::init")
	return nil
}
