package controllers

import (


	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yongwoon/gin-blog/entities"
)

type PostHandler struct {
	Db *gorm.DB
}

func (h *PostHandler) GetAll(c *gin.Context) {
	var posts []entities.Post
	h.Db.Find(&posts)
	c.JSON(200, posts)
}
func (h *PostHandler) CreatePost(c *gin.Context) {
	title, _ := c.GetPostForm("title")
	contents, _ := c.GetPostForm("contents")
	h.Db.Create(&entities.Post{Title: title, Contents: contents})

	post := entities.Post{}
	h.Db.Last(&post)
	c.JSON(200, post)
}
func (h *PostHandler) ShowPost(c *gin.Context) {
	post := entities.Post{}
	id := c.Param("id")
	h.Db.First(&post, id)
	c.JSON(200, post)
}
func (h *PostHandler) UpdatePost(c *gin.Context) {
	post := entities.Post{}
	id := c.Param("id")
	title, _ := c.GetPostForm("title")
	contents, _ := c.GetPostForm("contents")

	h.Db.First(&post, id)
	post.Title = title
	post.Contents = contents
	h.Db.Save(&post)
	c.JSON(200, post)
}
func (h *PostHandler) DeletePost(c *gin.Context) {
	post := entities.Post{}
	id := c.Param("id")
	h.Db.First(&post, id)
	h.Db.Delete(&post)
	c.JSON(200, gin.H{
		"result": "success",
	})
}
