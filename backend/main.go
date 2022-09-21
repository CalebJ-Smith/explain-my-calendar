package main

import (
	//"fmt"
	"context"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"
	"os"

	"google.golang.org/api/calendar/v3"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

//https://github.com/gin-gonic/gin

func get_cal_id() string {
	s := os.Getenv("CAL_ID")
	fmt.Printf("s: %v\n", s)
	return s
}

func calendar_things() {
	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	cal, c_err := calendarService.Calendars.Get(get_cal_id()).Do()
	if c_err != nil {
		log.Fatalln(c_err)
	}
	fmt.Printf("cal.Summary: %v\n", cal.Summary)
}

//To look into next: https://developers.google.com/identity/gsi/web/guides/display-button#javascript
// get users authenticated cause that's my problem for local development now.

func main() {
	calendar_things()
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
