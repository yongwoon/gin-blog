package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yongwoon/gin-blog/controller"
)

func Post(rg *gin.RouterGroup, dbConn *gorm.DB) {
	postHandler := controller.PostHandler{
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
