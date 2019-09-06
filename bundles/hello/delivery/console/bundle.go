package console

import (
	"github.com/fspace/ecm/bundles/hello"
	"github.com/fspace/ecm/bundles/hello/delivery/console/cmds"
	"github.com/fspace/ecm/core/app/console"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

func New(appInst *console.Application) *Bundle {
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
	console.BaseModule
	Config *hello.Config
}

// TODO 校验配置是否正确！
func (b *Bundle) Configure(viper *viper.Viper) error {
	log.Println("hello-bundle::configure")
	//log.Println(viper.AllKeys())
	//v := viper.Sub("hello-bundle")
	//
	//confVal := v.Get("someKey")
	//fmt.Println("someKey is : ", confVal)
	////fmt.Println(viper.AllKeys())
	//conf := hello.Config{}
	//v.Unmarshal(&conf)
	//fmt.Printf("config is : %#v \n", conf)
	//b.Config = &conf

	return nil
}

func (b *Bundle) Init() error {
	log.Println("hello-bundle::init")

	initCommands(b.App.KingpinApp)

	return nil
}

func initCommands(app *kingpin.Application) {
	cmd_addUser := &cmds.AddUserCommand{}
	addUser := app.Command("user-add", "create users.").Action(cmd_addUser.Run)
	// 配置命令入参flag
	addUser.Flag("data", "user data.").Short('d').StringVar(&cmd_addUser.Data)
	_ = addUser
}
