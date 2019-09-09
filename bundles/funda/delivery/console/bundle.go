package console

import (
	"github.com/fspace/ecm/bundles/funda/delivery/console/cmds"
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
}

// TODO 校验配置是否正确！
func (b *Bundle) Configure(viper *viper.Viper) error {
	return nil
}

func (b *Bundle) Init() error {
	log.Println("funda-bundle::init")

	initCommands(b.App.KingpinApp)

	return nil
}

//==========================================================================================
func initCommands(app *kingpin.Application) {
	// 每一个命令都可以关联 flag和args
	addContactAgentCmd := &cmds.AddContactAgent{}
	// 调用方法： go  run main.go contact-agent-add  --email yiqing@qq.com
	addContactAgent := app.Command("contact-agent-add", "create ContactAgent.").Action(addContactAgentCmd.Run)
	// 配置命令入参flag
	// addContactAgent.Flag("data", "user data.").Short('d').StringVar(&addContactAgentCmd.Data)
	//	 email := addContactAgent.Arg("email","CustomerEmailAddress").Required().String()
	// addContactAgentCmd.CustomerEmailAddress = *email
	// &addContactAgentCmd.CustomerEmailAddress = email
	addContactAgent.Flag("email", "CustomerEmailAddress").Required().StringVar(&addContactAgentCmd.CustomerEmailAddress)
}
