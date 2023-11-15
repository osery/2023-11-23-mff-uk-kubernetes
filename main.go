package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", handle)
	r.Run(":8080")
}

func handle(c *gin.Context) {
	c.String(http.StatusOK, "Hello MFF-UK!")
}
