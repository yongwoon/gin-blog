package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yongwoon/gin-blog/controllers"
)

// Router router list
func Router(dbConn *gorm.DB) {
	postHandler := controllers.PostHandler{
		Db: dbConn,
	}

	r := gin.Default()

	// Set cors
	// @see https://github.com/gin-contrib/cors
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index page",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.GET("/posts", postHandler.GetAll)
	r.POST("/posts", postHandler.CreatePost)
	r.GET("/posts/:id", postHandler.ShowPost)
	r.PATCH("/posts/:id", postHandler.UpdatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)
	r.Run(":3001")
}
