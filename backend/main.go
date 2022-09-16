package main

import (
	//"fmt"
	"net/http"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/gin-gonic/gin"
)

//https://github.com/gin-gonic/gin

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
