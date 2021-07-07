package controller

import (
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"testing"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var test_db *gorm.DB

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
