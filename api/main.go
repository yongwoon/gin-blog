package main

import (
	"github.com/yongwoon/gin-blog/config/initializer"
	"github.com/yongwoon/gin-blog/db"
	"github.com/yongwoon/gin-blog/router"
)

func main() {
	// 環境変数ロード
	initializer.InitDotenv()
	// DB 初期化
	dbConn := db.Init()
	// Router 初期化
	router.Router(dbConn)
}
