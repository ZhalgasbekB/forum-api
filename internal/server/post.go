package server

import (
	"fmt"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Services.PostsService.GetAllPostService(r.Context())
	if err != nil {
		return
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
	h.Services.PostsService.CreatePostService(r.Context(), *post)
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
	response.WriteJSON(w, http.StatusOK, "VSE NORM BRAT")
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
	err = h.Services.PostsService.DeleteUserPostService(r.Context(), deleteModel)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, "VSE UDALIL BRAT CHETCO")
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
	postComments, err := h.Services.PostsService.CommentsPostService(r.Context(), id)
	if err != nil {
		return
	}
	h.Services.LikePosts.GetLikesAndDislikesPostService(postComments.Post.PostId)
	response.WriteJSON(w, http.StatusOK, postComments)
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

	like, err := convert.LikeConvertor(r, user.ID)
	if err != nil {
		return
	}
	fmt.Println(like.LikeStatus)
	if err := h.Services.LikePosts.LikePostService(like); err != nil {
		return
	}
}
