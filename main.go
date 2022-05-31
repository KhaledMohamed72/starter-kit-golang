package main

import "github.com/gin-gonic/gin"

func main() {
	app := app()
	application := app()

	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)
		request.Context.JSON(200, gin.H{
			"message": "khaled",
		})
	})
	application.Gin.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// We have two closure
// Application closure --> hold global values
// Request closure --> hold temp values when a user call an api
