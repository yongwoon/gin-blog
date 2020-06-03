package main

import (
	"github.com/yongwoon/gin-blog/db"
	"github.com/yongwoon/gin-blog/routers"
)

func main() {
	dbConn := db.Init()
	routers.Router(dbConn)
}
