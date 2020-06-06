package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	v1 "github.com/yongwoon/gin-blog/router/v1"
)

// Router router list
func Router(dbConn *gorm.DB) {
	r := gin.Default()

	// Set cors
	// @see https://github.com/gin-contrib/cors
	r.Use(cors.Default())
	v1.InitRouter(r, dbConn)

	r.Run(":3001")
}
