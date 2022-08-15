package post

import (
	"database/sql"
	"example/go_blog/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ---- CURL COMMANDS FOR TESTING ----
// GET:       curl http://localhost:8080/api/posts
// GET_BY_ID: curl http://localhost:8080/api/posts/:id
// CREATE:    curl http://localhost:8080/api/posts/ --include --header "Content-Type: application/json" --request "POST" --data '{"title": "Blog Post 2","body": "The story of Golang"}'
// UPDATE:    curl http://localhost:8080/api/posts/:id --include --header "Content-Type: application/json" --request "PUT" --data '{"title": "Blog Post 2","body": "The most AWESOME one"}'
// DELETE:    curl http://localhost:8080/api/posts/:id --include --header "Content-Type: application/json" --request "DELETE"

// GET /posts
func getPosts(c *gin.Context) {
	// Retrieve our Postgres DB
	db := utils.GetDB()

	// Create a post slice to hold data from returned rows
	var allPosts []BlogPost

	// Query our database
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Println("getPosts FAILURE: ", err)
		return
	}

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var newPost BlogPost
		if err := rows.Scan(&newPost.ID, &newPost.Title, &newPost.Body); err != nil {
			log.Println("getPosts FAILURE:", err)
			return
		}

		allPosts = append(allPosts, newPost)
	}

	// Return the http status and response
	c.IndentedJSON(http.StatusOK, allPosts)
	// c.Redirect(http.StatusOK, "/home/")
}

// GET /posts/:id
func getPost(c *gin.Context) {
	// Retrieve our Postgres DB
	db := utils.GetDB()

	// Get the id of our item
	id := c.Param("id")

	// Query our database and save the new post
	var newPost BlogPost
	err := db.QueryRow("SELECT * FROM posts WHERE id= $1", id).Scan(&newPost.ID, &newPost.Title, &newPost.Body)

	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		} else {
			log.Println("getPost FAILURE: ", err)
		}
		return
	}

	// Return the http status and response
	c.IndentedJSON(http.StatusOK, newPost)
}

// POST /posts/:id
func createNewPost(c *gin.Context) {
	// Retrieve our Postgres DB
	db := utils.GetDB()

	// Call BindJSON to bind the received JSON to our newPost object
	var newPost BlogPost
	if err := c.BindJSON(&newPost); err != nil {
		return
	}

	// Insert the new Post into our DB
	sqlStatement := `INSERT INTO posts (title, body) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, newPost.Title, newPost.Body)
	if err != nil {
		log.Println("createPost FAILURE: ", err)
		return
	}

	// Return the http status and response
	c.IndentedJSON(http.StatusCreated, newPost)
	// c.Redirect(http.StatusCreated, "/home/")
}

// PUT /posts/:id
func updatePost(c *gin.Context) {
	// Retrieve our Postgres DB
	db := utils.GetDB()

	// Call BindJSON to bind the received JSON to our newPost object
	var targetPost BlogPost
	targetPost.ID = c.Param("id")
	if err := c.BindJSON(&targetPost); err != nil {
		return
	}

	// Update the target post in our DB
	sqlStatement := `UPDATE posts SET title = $2, body = $3 WHERE id = $1;`
	res, err := db.Exec(sqlStatement, targetPost.ID, targetPost.Title, targetPost.Body)
	if err != nil {
		log.Println("updatePost FAILURE: ", err)
		return
	}

	// Confirm DB changes
	count, err := res.RowsAffected()
	if err != nil {
		log.Println("deletePost Rows Affected FAILURE: ", err)
	}

	log.Println("Post Updated. Rows Affected " + strconv.Itoa(int(count)))

	// Return the http status and response
	c.IndentedJSON(http.StatusOK, targetPost)
}

// DELETE /posts/:id
func deletePost(c *gin.Context) {
	// Retrieve our Postgres DB
	db := utils.GetDB()

	// Get the id of our item
	id := c.Param("id")

	// Update the target post in our DB
	sqlStatement := `DELETE FROM posts WHERE id = $1;`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Println("deletePost FAILURE: ", err)
		return
	}

	// Confirm DB changes
	count, err := res.RowsAffected()
	if err != nil {
		log.Println("deletePost Rows Affected FAILURE: ", err)
	}

	c.IndentedJSON(http.StatusOK, "Post Deleted. Rows Affected "+strconv.Itoa(int(count)))
}
