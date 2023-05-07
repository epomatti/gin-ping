package ginping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	route   = "/health"
	message = "OK"
)

func Add(r *gin.Engine) {
	r.GET(route, stringHandler)
	r.HEAD(route, headHandler)
}

func stringHandler(c *gin.Context) {
	c.String(http.StatusOK, message)
}

func headHandler(c *gin.Context) {
	c.String(http.StatusOK, "")
}
