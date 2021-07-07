package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("DATABASE_SCHEMA")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	return db
}

// func InitTest() *gorm.DB {
// 	DBMS := "mysql"
// 	USER := os.Getenv("DATABASE_USER")
// 	PASS := os.Getenv("DATABASE_PASSWORD")
// 	PROTOCOL := "tcp(db:3306)"
// 	DBNAME := os.Getenv("DATABASE_SCHEMA")
// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
// 	db, err := gorm.Open(DBMS, CONNECT)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	db.Set("gorm:table_options", "ENGINE=InnoDB")
// 	db.LogMode(true)
// 	db.AutoMigrate(&entity.Post{})

// 	return db
// }
