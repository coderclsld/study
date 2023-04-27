package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewMysqlConnect(host string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(host), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect error:", err.Error())
		return nil
	}
	return db
}
