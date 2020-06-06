package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(r *gin.Engine, dbConn *gorm.DB) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})

		Post(v1, dbConn)
	}
}
