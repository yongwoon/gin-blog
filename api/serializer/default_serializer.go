package serializer

import (
	"github.com/gin-gonic/gin"
)

type SuccessSerializer struct {
	C *gin.Context
}

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (s *SuccessSerializer) Response() SuccessResponse {
	return SuccessResponse{
		Status:  "success",
		Message: "",
	}
}

//  TODO: error message format
// {
// 	"error": {
// 		"code": 190,
// 		"message": "Message describing the error",
// 		"error_user_title": "A title",
// 		"error_user_msg": "A message"
// 	}
// }
