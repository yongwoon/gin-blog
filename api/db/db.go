package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "blog_development"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	return db
}
