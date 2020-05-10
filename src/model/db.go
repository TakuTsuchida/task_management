package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db = dbConnect()
	db.AutoMigrate(&Auth{})
}

func dbConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "docker_user"
	PASS := "docker_user_pwd"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "docker_db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	var err error
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}
