package server

import (
	"net/http"
	"time"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()
	rateLimiter := RateLimitMiddleware(5, 1*time.Second)

	mux.HandleFunc("/d3", h.Home)                                               // 200 (GET METHOD) get all posts
	mux.HandleFunc("/login", h.Login)                                           // 200 (POST METHOD)
	mux.HandleFunc("/register", h.Register)                                     // 201 (POST METHOD)
	mux.Handle("/logout", h.RequiredAuthentication(http.HandlerFunc(h.Logout))) // 200 (POST METHOD)

	// ADMIN, MODERATOR, USER
	mux.Handle("/admin", h.AdminVerification(http.HandlerFunc(h.Admin)))                       // GET
	mux.Handle("/admin/wants", h.AdminVerification(http.HandlerFunc(h.UsersWants)))            // GET
	mux.Handle("/admin/reports", h.AdminVerification(http.HandlerFunc(h.AdminReports)))        // GET
	mux.Handle("/admin/role-update", h.AdminVerification(http.HandlerFunc(h.AdminChangeRole))) // PUT

	mux.Handle("/admin/response-moderator", h.AdminVerification(http.HandlerFunc(h.AdminResponseModerator))) // PUT
	mux.Handle("/admin/response-user", h.AdminVerification(http.HandlerFunc(h.AdminResponseUser)))

	mux.Handle("/admin/user-delete", h.AdminVerification(http.HandlerFunc(h.AdminDeleteUser)))       // DELETE
	mux.Handle("/admin/post-delete", h.AdminVerification(http.HandlerFunc(h.AdminDeletePost)))       // DELETE
	mux.Handle("/admin/comment-delete", h.AdminVerification(http.HandlerFunc(h.AdminDeleteComment))) // DELETE

	mux.Handle("/admin/create-category", h.AdminVerification(http.HandlerFunc(h.AdminCreateCategory)))
	mux.Handle("/admin/delete-category", h.AdminVerification(http.HandlerFunc(h.AdminDeleteCategory))) // DELETE

	mux.Handle("/moderator/reports", h.ModeratorVerification(http.HandlerFunc(h.ModeratorReports))) // GET
	mux.Handle("/moderator/report", h.ModeratorVerification(http.HandlerFunc(h.ModeratorReport)))   // SAME DONT NEED JSON

	mux.Handle("/user/wants", h.RequiredAuthentication(http.HandlerFunc(h.UserWants))) // GET
	mux.Handle("/user/want", h.RequiredAuthentication(http.HandlerFunc(h.UserWant)))   // CHECK ONLY USER ??? DONT NEED JSON
	/// ADMIN

	///// ADVANCED FEATURE
	mux.Handle("/notifications", h.RequiredAuthentication(http.HandlerFunc(h.Notifications)))      // GET
	mux.Handle("/notification/update", h.RequiredAuthentication(http.HandlerFunc(h.Notification))) // PUT

	///// ACTIVITY
	// mux.Handle("/activity", h.RequiredAuthentication(http.HandlerFunc(h.Notification))) //  ???

	mux.Handle("/d3/category", h.RequiredAuthentication(http.HandlerFunc(h.PostCategory)))     // 200 (GET METHOD) user posts
	mux.Handle("/d3/user-likes", h.RequiredAuthentication(http.HandlerFunc(h.LikedPostsUser))) // 200 (GET METHOD)
	mux.Handle("/d3/user-posts", h.RequiredAuthentication(http.HandlerFunc(h.PostsUser)))      // 200 (GET METHOD) user posts

	mux.Handle("/d3/post", h.RequiredAuthentication(http.HandlerFunc(h.Post)))               // 200   (GET METHOD) post and his comments id ?? category
	mux.Handle("/d3/post-create", h.RequiredAuthentication(http.HandlerFunc(h.CreatePosts))) // 201   (POST METHOD) create post
	mux.Handle("/d3/post-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdatePost)))  // 202   (PUT METHOD) update
	mux.Handle("/d3/post-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeletePost)))  // 202   (DELETE METHOD) delete

	mux.Handle("/d3/post-like", h.RequiredAuthentication(http.HandlerFunc(h.LikePosts)))       // 200  // (POST METHOD)
	mux.Handle("/d3/comment-like", h.RequiredAuthentication(http.HandlerFunc(h.LikeComments))) // 200  // (POST METHOD)

	mux.Handle("/d3/comment-create", h.RequiredAuthentication(http.HandlerFunc(h.CreateComment))) // 201 	(POST METHOD) create
	mux.Handle("/d3/comment-update", h.RequiredAuthentication(http.HandlerFunc(h.UpdateComment))) // 202	(PUT METHOD) update
	mux.Handle("/d3/comment-delete", h.RequiredAuthentication(http.HandlerFunc(h.DeleteComment))) // 202 	(DELETE METHOD) delete

	return rateLimiter(h.IsAuthenticated(mux))
}
