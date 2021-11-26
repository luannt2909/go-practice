package dbfx

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

func provideGormDB() (*gorm.DB, error) {
	uri := viper.GetString("MYSQL_URI")
	return gorm.Open(mysql.Open(uri))
}
