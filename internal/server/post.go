package server

import (
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	posts, err := h.Services.PostsService.GetAllPostService(r.Context())
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}

	res, err := h.Services.LikePosts.GetLikeAndDislikeAllPostService()
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}
	if err := h.Services.PostsService.CreatePostService(r.Context(), *post); err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}

	if err := h.Services.PostsService.UpdateUserPostService(r.Context(), *post); err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}

	if err := h.Services.PostsService.DeleteUserPostService(r.Context(), deleteModel); err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}

	postComments, err := h.Services.CommentService.CommentsLikesNames(r.Context(), id)
	if err != nil {
		if err == errors.ErrSQLNoRows {
			response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, errors.ErrNotFoundData.Error()))
		}
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
	}

	likes, dislike, err := h.Services.LikePosts.GetLikesAndDislikesPostService(id)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}
	postComments.Post.Like = likes
	postComments.Post.Dislike = dislike
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}
	if err := h.Services.LikePosts.LikePostService(like); err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
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
		response.WriteJSON(w, http.StatusInternalServerError, errors.NewError(500, err.Error()))
		return
	}
	response.WriteJSON(w, http.StatusOK, likedPosts)
}
