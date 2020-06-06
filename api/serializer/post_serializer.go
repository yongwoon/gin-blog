package serializer

import (
	"github.com/gin-gonic/gin"
	"github.com/yongwoon/gin-blog/entity"
)

type PostSerializer struct {
	C *gin.Context
	entity.Post
}

type PostsSerializer struct {
	C     *gin.Context
	Posts []entity.Post
}

type PostResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Contents  string `json:"contents"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (s *PostSerializer) Response() PostResponse {
	return PostResponse{
		ID:        s.ID,
		Title:     s.Title,
		Contents:  s.Contents,
		CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		//UpdatedAt:  s.UpdatedAt.UTC().Format(time.RFC3339Nano),
		UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
	}
}

func (s *PostsSerializer) Response() []PostResponse {
	response := []PostResponse{}
	for _, post := range s.Posts {
		serializer := PostSerializer{s.C, post}
		response = append(response, serializer.Response())
	}
	return response
}
