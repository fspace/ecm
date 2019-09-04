package backend

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func New() *Bundle {
	return &Bundle{}
}

type Bundle struct {
}

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
