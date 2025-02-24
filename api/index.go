// api/index.go
package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpErr, ok := err.(HttpError); ok {
			c.JSON(httpErr.Status, H{
				"message": httpErr.Error(),
			})
		} else {
			c.JSON(500, H{
				"message": fmt.Sprintf("%v", err),
			})
		}
	}))

	
	server.GET("/", func(c *Context) {
		c.JSON(200, H{
			"message": "OK",
		})
	})

	// Hugs Endpoint
	server.GET("/hugs", func(c *Context) {
		totalGifs := 29
		rand.NewSource(time.Now().UnixNano())
		randomNumber := rand.Intn(totalGifs) + 1
		gifURL := fmt.Sprintf("https://cdn.lucia-dev.com/hug%02d.gif", randomNumber)

		c.JSON(200, H{
			"url": gifURL,
		})
	})

	// Pats Endpoint
	server.GET("/pats", func(c *Context) {
		totalGifs := 30
		rand.NewSource(time.Now().UnixNano())
		randomNumber := rand.Intn(totalGifs) + 1
		gifURL := fmt.Sprintf("https://cdn.lucia-dev.com/pat%02d.gif", randomNumber)

		c.JSON(200, H{
			"url": gifURL,
		})
	})

	server.Handle(w, r)
}
