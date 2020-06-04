package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

// Router router list
func Router(dbConn *gorm.DB) {
	r := gin.Default()

	// Set cors
	// @see https://github.com/gin-contrib/cors
	r.Use(cors.Default())
	initV1Router(r, dbConn)

	r.Run(":3001")
}
