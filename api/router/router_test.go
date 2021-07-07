package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yongwoon/gin-blog/config/initializer"
	"github.com/yongwoon/gin-blog/db"
)

func TestApiV1PingRoute(t *testing.T) {
	// Prepare go test
	initializer.InitDotenv()
	engine := gin.Default()
	dbConn := db.Init()

	router := Router(engine, dbConn)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
