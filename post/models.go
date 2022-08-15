package post

// Post represents a blog post
type BlogPost struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
