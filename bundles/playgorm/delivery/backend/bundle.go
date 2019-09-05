package backend

import (
	"fmt"
	"github.com/fspace/ecm/core/app"
	"github.com/jinzhu/gorm"
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
func (b *Bundle) Configure(viper *viper.Viper) {
	log.Println("playgorm-bundle::configure")
	//v := viper.Sub("hello-bundle")
	//
	//confVal := v.Get("someKey")
	//fmt.Println("someKey is : ", confVal)
	//fmt.Println(viper.AllKeys())
}

func (b *Bundle) Init() {
	log.Println("playgorm-bundle::init")
}

func (b *Bundle) Start() error {
	log.Println("playgorm-bundle::start|run")

	// 依赖注入 需要什么就声明什么 ,不够了 继续添加参数
	_run(b.App)

	return nil
}

//====================================================================
// ## 下面是真正有用的运行代码

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func _run(appInst *app.Application) {
	db := appInst.DB
	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	fmt.Printf("the product is : %#v \n", product)

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&product)
}
