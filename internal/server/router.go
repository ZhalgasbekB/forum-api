package server

import (
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", h.Login)                                           // (POST METHOD)
	mux.HandleFunc("/register", h.Register)                                     // (POST METHOD)
	mux.Handle("/logout", h.RequiredAuthentication(http.HandlerFunc(h.Logout))) // (POST METHOD)
	mux.HandleFunc("/d3", h.Home)                                               // (GET METHOD) get all posts

	mux.Handle("/d3/user-likes", h.RequiredAuthentication(http.HandlerFunc(h.LikedPostsUser))) // (GET METHOD)			  // ???
	mux.Handle("/d3/user-posts", h.RequiredAuthentication(http.HandlerFunc(h.PostsUser)))      // (GET METHOD) user posts // ???

	mux.Handle("/d3/post", h.RequiredAuthentication(http.HandlerFunc(h.Post)))                  // (GET METHOD) post and his comments
	mux.Handle("/d3/post-create", h.RequiredAuthentication(http.HandlerFunc(h.CreatePosts)))    // (POST METHOD) create post
	mux.Handle("/d3/post-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdatePost)))     // (PUT METHOD) update
	mux.Handle("/userd3/post-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeletePost))) // (DELETE METHOD) delete

	mux.Handle("/d3/post-like", h.RequiredAuthentication(http.HandlerFunc(h.LikePosts)))           // (POST METHOD)
	mux.Handle("/userd3/comment-like", h.RequiredAuthentication(http.HandlerFunc(h.LikeComments))) // (POST METHOD)

	mux.Handle("/d3/comment-create", h.RequiredAuthentication(http.HandlerFunc(h.CreateComment))) // (POST METHOD) create
	mux.Handle("/d3/comment-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdateComment))) // (PUT METHOD) update
	mux.Handle("/d3/comment-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeleteComment)))  // (DELETE METHOD) delete

	return h.IsAuthenticated(mux)
}
