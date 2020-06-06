package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yongwoon/gin-blog/config/initializer"
	"github.com/yongwoon/gin-blog/db"
	"github.com/yongwoon/gin-blog/middleware"
	"github.com/yongwoon/gin-blog/router"
)

func main() {
	// 環境変数ロード
	initializer.InitDotenv()

	// DB 初期化
	dbConn := db.Init()

	engine := gin.Default()

	// set logger
	engine.Use(middleware.LoggerToFile())

	// Set cors
	// @see https://github.com/gin-contrib/cors
	engine.Use(cors.Default())

	// Router 初期化
	router.Router(engine, dbConn)
}
