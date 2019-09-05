package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Context represents a global-vars namespace , store shared vars here . alias is AppContext ,dry(dont repeat yourself) for removing the prefix App
type Context struct {
	DB *gorm.DB
}
