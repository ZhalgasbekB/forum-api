package server

import (
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/google/login", h.Google) // POST
	mux.HandleFunc("/google/callback", h.GoogleCallback)
	mux.HandleFunc("/github/login", h.GitHub) // POST
	mux.HandleFunc("/github/callback", h.GitHubCallback)

	mux.HandleFunc("/d3", h.Home)           // 200 (GET METHOD) get all posts
	mux.HandleFunc("/login", h.Login)       // 200 (POST METHOD)
	mux.HandleFunc("/register", h.Register) // 201 (POST METHOD)

	mux.Handle("/logout", h.RequiredAuthentication(http.HandlerFunc(h.Logout)))                // 200 (POST METHOD)
	mux.Handle("/d3/category", h.RequiredAuthentication(http.HandlerFunc(h.PostCategory)))     // 200 (GET METHOD) user posts
	mux.Handle("/d3/user-likes", h.RequiredAuthentication(http.HandlerFunc(h.LikedPostsUser))) // 200 (GET METHOD)
	mux.Handle("/d3/user-posts", h.RequiredAuthentication(http.HandlerFunc(h.PostsUser)))      // 200 (GET METHOD) user posts

	mux.Handle("/d3/post", h.RequiredAuthentication(http.HandlerFunc(h.Post)))               // 200    	(GET METHOD) post and his comments id ?? category
	mux.Handle("/d3/post-create", h.RequiredAuthentication(http.HandlerFunc(h.CreatePosts))) // 201   // (POST METHOD) create post
	mux.Handle("/d3/post-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdatePost)))  // 202   // (PUT METHOD) update
	mux.Handle("/d3/post-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeletePost)))  // 202   // (DELETE METHOD) delete

	mux.Handle("/d3/post-like", h.RequiredAuthentication(http.HandlerFunc(h.LikePosts)))       // 200  // (POST METHOD)
	mux.Handle("/d3/comment-like", h.RequiredAuthentication(http.HandlerFunc(h.LikeComments))) // 200  // (POST METHOD)

	mux.Handle("/d3/comment-create", h.RequiredAuthentication(http.HandlerFunc(h.CreateComment))) // 201 	(POST METHOD) create
	mux.Handle("/d3/comment-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdateComment))) // 202	(PUT METHOD) update
	mux.Handle("/d3/comment-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeleteComment))) // 202 	(DELETE METHOD) delete

	return h.IsAuthenticated(mux)
}
