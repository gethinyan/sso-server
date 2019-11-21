package models

import (
	"fmt"
	"log"

	"github.com/gethinyan/sso-server/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

var dbConn *gorm.DB

// init 初始化数据库连接
func init() {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		setting.Database.Host, setting.Database.Port, setting.Database.User,
		setting.Database.Dbname, setting.Database.Sslmode, setting.Database.Password)
	dbConn, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	dbConn.AutoMigrate(&User{})
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(100)
}
