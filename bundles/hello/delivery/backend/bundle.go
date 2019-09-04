package backend

import (
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
}

func (b *Bundle) Init() {
	log.Println("hello-bundle::init")
}
