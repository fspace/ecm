package app

import "github.com/spf13/viper"

type Module interface {
	Configure(v *viper.Viper)

	Init()
	// Stop() error
}
