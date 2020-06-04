package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yongwoon/gin-blog/controllers"
)


func Post(rg *gin.RouterGroup, dbConn *gorm.DB) {
	postHandler := controllers.PostHandler{
		Db: dbConn,
	}

	posts := rg.Group("/posts")
	{
		posts.GET("/", postHandler.Index)
		posts.POST("/", postHandler.Create)
		posts.GET("/:id", postHandler.Show)
		posts.PATCH("/:id", postHandler.Update)
		posts.DELETE("/:id", postHandler.Destroy)
	}
}
