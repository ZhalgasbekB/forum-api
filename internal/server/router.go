package server

import (
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", h.Register) // (POST METHOD)
	mux.HandleFunc("/login", h.Login)       // (POST METHOD)

	mux.HandleFunc("/userd3", h.Home)                                                           // (GET METHOD) get all posts
	mux.Handle("/userd3/posts", h.RequiredAuthentication(http.HandlerFunc(h.CreatePosts)))      // (POST METHOD) create post
	mux.Handle("/userd3/myposts", h.RequiredAuthentication(http.HandlerFunc(h.PostsUser)))      // (GET METHOD) user posts
	mux.Handle("/userd3/post", h.RequiredAuthentication(http.HandlerFunc(h.Post)))              // (GET METHOD) post and his comments
	mux.Handle("/userd3/post-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdatePost))) // (PUT METHOD) update
	mux.Handle("/userd3/post-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeletePost))) // (DELETE METHOD) delete

	// mux.HandleFunc("/userd3/likeposts", h.LikePosts)											   // (GET METHOD)
	mux.Handle("/userd3/post-like", h.RequiredAuthentication(http.HandlerFunc(h.LikePosts)))       // (POST METHOD)
	mux.Handle("/userd3/comment-like", h.RequiredAuthentication(http.HandlerFunc(h.LikeComments))) // (POST METHOD)

	mux.HandleFunc("/userd3/comments", h.Comments)                                                    // (GET METHOD) comments
	mux.Handle("/userd3/comment", h.RequiredAuthentication(http.HandlerFunc(h.CommentByID)))          // (GET METHOD) comment by id
	mux.Handle("/userd3/comment-create", h.RequiredAuthentication(http.HandlerFunc(h.CreateComment))) // (POST METHOD) create
	mux.Handle("/userd3/comment-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdateComment))) // (PUT METHOD) update
	mux.Handle("/userd3/comment-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeleteComment))) // (DELETE METHOD) delete

	return h.IsAuthenticated(mux) // Check Authentication
}
