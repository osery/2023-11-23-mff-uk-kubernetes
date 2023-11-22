package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	time.Sleep(10 * time.Second)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", handleIndex)
	r.GET("/killz", handleKillz)

	r.Run(":8080")
}

func handleIndex(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%s: Hello MFF-UK!\n", hostname))
}

func handleKillz(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%s: Shutting down.\n", hostname))
	stop()
}
