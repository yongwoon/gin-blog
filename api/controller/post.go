package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yongwoon/gin-blog/entity"
	"github.com/yongwoon/gin-blog/serializer"
)

type PostHandler struct {
	Db *gorm.DB
}

func (h *PostHandler) Index(c *gin.Context) {
	var posts []entity.Post
	h.Db.Find(&posts)

	serializer := serializer.PostsSerializer{c, posts}
	c.JSON(http.StatusOK, gin.H{"posts": serializer.Response()})
}

func (h *PostHandler) Create(c *gin.Context) {
	title, _ := c.GetPostForm("title")
	contents, _ := c.GetPostForm("contents")
	h.Db.Create(&entity.Post{Title: title, Contents: contents})

	post := entity.Post{}
	h.Db.Last(&post)

	serializer := serializer.PostSerializer{c, post}
	c.JSON(http.StatusCreated, gin.H{"post": serializer.Response()})
}

func (h *PostHandler) Show(c *gin.Context) {
	post := entity.Post{}
	id := c.Param("id")
	h.Db.First(&post, id)

	serializer := serializer.PostSerializer{c, post}
	c.JSON(http.StatusCreated, gin.H{"post": serializer.Response()})
}

func (h *PostHandler) Update(c *gin.Context) {
	post := entity.Post{}
	id := c.Param("id")
	title, _ := c.GetPostForm("title")
	contents, _ := c.GetPostForm("contents")

	h.Db.First(&post, id)
	post.Title = title
	post.Contents = contents
	h.Db.Save(&post)

	serializer := serializer.PostSerializer{c, post}
	c.JSON(http.StatusOK, gin.H{"post": serializer.Response()})
}

func (h *PostHandler) Destroy(c *gin.Context) {
	post := entity.Post{}
	id := c.Param("id")
	h.Db.First(&post, id)
	h.Db.Delete(&post)

	serializer := serializer.SuccessSerializer{c}
	c.JSON(http.StatusCreated, serializer.Response())
}
