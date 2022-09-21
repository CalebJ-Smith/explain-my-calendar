package main

import (
	//"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

//https://github.com/gin-gonic/gin

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong" + sb,
		})
	})
	router.Run(":8080")
}
