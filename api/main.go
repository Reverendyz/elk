package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	PORT = os.Getenv("API_PORT")
)

func main() {
	server := gin.Default()
	server.GET("/", handler)
	server.Use(cors.Default())

	server.Run(fmt.Sprintf("%s:%s", "0.0.0.0", PORT))
}

func handler(c *gin.Context) {
	time := time.Now().Format("2006-01-02 15:04:05")
	host := os.Getenv("HOSTNAME")
	c.Header("Content-Type", "text/plain")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization")
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Accepted %s - Time %s", host, time))
}
