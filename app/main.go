package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", handler)

	// r.Run(":3000")
	r.Run(":" + port)
}

func handler(ctx *gin.Context) {

	user := User{"User", 20}

	ctx.HTML(200, "index.html", gin.H{
		"user": user,
	})
}
