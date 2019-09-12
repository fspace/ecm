package backend

import (
	"github.com/fspace/ecm/bundles/funda/delivery/backend/controllers"
	"github.com/fspace/ecm/bundles/funda/gateways/memory"
	"github.com/fspace/ecm/bundles/funda/usecases"
	"github.com/fspace/ecm/core/app"
	"github.com/spf13/viper"
	"log"
)

func New(appInst *app.Application) *Bundle {
	b := &Bundle{}
	b.App = appInst
	return b
}

type Bundle struct {
	app.BaseModule
}

// TODO 校验配置是否正确！
func (b *Bundle) Configure(viper *viper.Viper) error {
	log.Println("funda-bundle::configure")
	return nil
}

func (b *Bundle) Init() error {
	log.Println("funda-bundle::init")

	buildRouter(b.App)
	return nil
}

func (b *Bundle) Start() error {
	log.Println("funda-bundle::start|run")
	return nil
}

//====================================================================
// ## 下面是真正做事的代码
// --------------------------------------------------------
// 构建本模块下的路由
func buildRouter(appInst *app.Application) {
	r := appInst.GinRouter
	// Simple group: v1
	rg := r.Group("/funda")
	{
		//rg.POST("/agent-contact", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		//})
		interactor := usecases.NewContactAgentInteractor(memory.NewInMemoryHouseRepository())
		rg.POST("/agent-contact", controllers.NewAgentController(interactor).Contact)

	}
}
