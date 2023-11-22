package main

import (
	"fmt"
	"html"
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
	server.GET("/", root)
	server.GET("/secrets", secrets)
	server.GET("/page", page)
	server.Use(cors.Default())

	server.Run(fmt.Sprintf("%s:%s", "0.0.0.0", PORT))
}

func root(c *gin.Context) {
	time := time.Now().Format("2000-12-01 12:23:30 pm")
	host := os.Getenv("HOSTNAME")
	c.Header("Content-Type", "text/plain")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization")
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Accepted %s - Time %s", host, time))
}
func secrets(c *gin.Context) {
	time := time.Now().Format("2000-12-01 12:23:30 pm")
	host := os.Getenv("HOSTNAME")
	c.Header("Content-Type", "text/plain")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization")
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("admin_password = %sH4CK-m3 - Time: %s", host, time))
}

func page(c *gin.Context) {
	tag := "<h1>weak</h1>"
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization")
	c.IndentedJSON(http.StatusOK, html.EscapeString(tag))
}
