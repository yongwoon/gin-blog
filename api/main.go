package main

import (
	"github.com/yongwoon/gin-blog/db"
	"github.com/yongwoon/gin-blog/router"
)

func main() {
	dbConn := db.Init()
	router.Router(dbConn)
}
