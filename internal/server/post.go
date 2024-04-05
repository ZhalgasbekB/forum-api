package server

import (
	"database/sql"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	posts, err := h.Services.PostsService.GetAllPostService(r.Context())
	if err != nil {
		return
	}

	res, err := h.Services.LikePosts.GetLikeAndDislikeAllPostService()   
	if err != nil {
		return
	}

	for i, v := range posts {
		if res[v.PostId] != nil {
			posts[i].Like = res[v.PostId][0]
			posts[i].Dislike = res[v.PostId][1]
		}
	}
	response.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) CreatePosts(w http.ResponseWriter, r *http.Request) {  
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	user := contextUser(r)
	post, err := convert.ConvertCreatePost(r, user.ID)
	if err != nil {
		return
	}
	if err := h.Services.PostsService.CreatePostService(r.Context(), *post); err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) { 
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	user := contextUser(r)
	post, err := convert.ConvertUpdatePost(r, user.ID)
	if err != nil {
		return
	}
	err = h.Services.PostsService.UpdateUserPostService(r.Context(), *post)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {  
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	user := contextUser(r)
	deleteModel, err := convert.ConvertDeletePost(r, user.ID)
	if err != nil {
		return
	}

	if err := h.Services.PostsService.DeleteUserPostService(r.Context(), deleteModel); err != nil {
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	id, err := convert.ConvertParamID(r)
	if err != nil {
		return
	}
	likes, dislike, err := h.Services.LikePosts.GetLikesAndDislikesPostService(id)
	if err == sql.ErrNoRows {
		likes = 0
		dislike = 0
	}
	arr, err := h.Services.CommentService.CommentsLikesNames(r.Context(), id)
	arr.Post.Like = likes
	arr.Post.Dislike = dislike

	response.WriteJSON(w, http.StatusOK, arr)
}

func (h *Handler) PostsUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	user := contextUser(r)
	postU, err := h.Services.PostsService.GetUserPostService(r.Context(), user.ID)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, postU)
}

func (h *Handler) LikePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		return
	}
	user := contextUser(r)

	like, err := convert.LikePostConvertor(r, user.ID)
	if err != nil {
		return
	}
	if err := h.Services.LikePosts.LikePostService(like); err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) LikedPostsUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	user := contextUser(r)
	likedPosts, err := h.Services.LikePosts.GetUserLikedPostService(user.ID)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, likedPosts)
}
