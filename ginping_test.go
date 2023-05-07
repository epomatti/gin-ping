package ginping_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	ginping "github.com/epomatti/gin-ping"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestPingGET(t *testing.T) {
	// Arrange
	router := gin.Default()
	ginping.Add(router)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPingHEAD(t *testing.T) {
	// Arrange
	router := gin.Default()
	ginping.Add(router)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("HEAD", "/ping", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}
