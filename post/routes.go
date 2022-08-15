package post

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes routes for the App
func InitializeRoutes(r *gin.RouterGroup) {
	post := r.Group("/posts")

	post.GET("", getPosts)
	post.GET("/", getPosts)

	post.GET("/:id", getPost)
	post.PUT("/:id", updatePost)
	post.DELETE("/:id", deletePost)

	post.POST("", createNewPost)
	post.POST("/", createNewPost)
}
