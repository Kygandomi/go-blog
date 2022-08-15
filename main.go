package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example/go_blog/post"
	"example/go_blog/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(quote.Go())

	// ============================================
	// Connect to Postgres DB
	db, err := utils.ConnectDB()
	if err != nil {
		log.Println(err)
	}

	// Close db at function end
	defer db.Close()

	// Test our connection
	ping := db.Ping()
	if ping != nil {
		log.Fatal(ping)
	}

	fmt.Println("Postgres DB Connected.")

	// ============================================

	// Create Gin Router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowOrigins = []string{"*"}

	router.Use(cors.New(config))

	api := router.Group("/api")
	post.InitializeRoutes(api)

	// Serve index
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.GET("/home/", func(c *gin.Context) {
		c.File("./public/index.html")
	})
	router.GET("/home/:id", func(c *gin.Context) {
		c.File("./public/edit.html")
	})

	router.Run("localhost:8080")
}
