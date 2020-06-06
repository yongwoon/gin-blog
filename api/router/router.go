package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	v1 "github.com/yongwoon/gin-blog/router/v1"
)

// Router router list
func Router(engine *gin.Engine, dbConn *gorm.DB) {
	v1.InitRouter(engine, dbConn)

	engine.Run(":3001")
}
