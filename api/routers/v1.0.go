package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yongwoon/gin-blog/controllers"
)


func initV1Router(r *gin.Engine, dbConn *gorm.DB) {
	postHandler := controllers.PostHandler{
		Db: dbConn,
	}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "index page",
			})
		})

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})

		v1.GET("/posts", postHandler.Index)
		v1.POST("/posts", postHandler.Create)
		v1.GET("/posts/:id", postHandler.Show)
		v1.PATCH("/posts/:id", postHandler.Update)
		v1.DELETE("/posts/:id", postHandler.Destroy)
	}
}
