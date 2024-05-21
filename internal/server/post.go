package server

import (
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/json"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	posts, err := h.Services.PostsService.GetAllPostService(r.Context())
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.Services.LikePosts.GetLikeAndDislikeAllPostService()
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	for i, v := range posts {
		if res[v.PostId] != nil {
			posts[i].Like = res[v.PostId][0]
			posts[i].Dislike = res[v.PostId][1]
		}
	}

	user := contextUser(r)
	newNotifications := false
	if user != nil {
		is_read, err := h.Services.Nothification.NotificationIsReadService(user.ID)
		if err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
			return
		}
		newNotifications = is_read
	}

	json.WriteJSON(w, http.StatusOK, &model.PostsNotification{
		Posts:           posts,
		NewNotification: newNotifications,
	})
}

func (h *Handler) CreatePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	user := contextUser(r)
	post, err := convert.ConvertCreatePost(r, user.ID)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.PostsService.CreatePostService(r.Context(), *post); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
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
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.PostsService.UpdateUserPostService(r.Context(), *post); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
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
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.PostsService.DeleteUserPostService(r.Context(), deleteModel); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
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
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := h.Services.PostsService.GetIdPostService(r.Context(), id)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	comments, err := h.Services.CommentService.CommentsLikesNames(r.Context(), id)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.WriteJSON(w, http.StatusOK, model.PostCommentsDTO{Post: post, Comments: comments})
}

func (h *Handler) PostsUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	user := contextUser(r)
	postU, err := h.Services.PostsService.GetUserPostService(r.Context(), user.ID)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.WriteJSON(w, http.StatusOK, postU)
}

func (h *Handler) LikePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		return
	}

	user := contextUser(r)
	like, user_id, err := convert.LikePostConvertor(r, user.ID)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	isLiked, err := h.Services.LikePosts.LikePostService(like)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	notification := convert.NotificationCreateLikes(like, user_id)
	if err := h.Services.Nothification.CreateService(notification, isLiked); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
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
	likedPosts, err := h.Services.LikePosts.GetUserLikedPostService(user.ID, true)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.WriteJSON(w, http.StatusOK, likedPosts)
}

func (h *Handler) PostCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	postsCategory, err := h.Services.PostsService.PostsCategoryService(r.Context(), r.URL.Query().Get("name"))
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.WriteJSON(w, http.StatusOK, postsCategory)
}

func (h *Handler) UploadPostImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		imageDate, err := convert.ConvertUploadPostImage(r)
		if err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = h.Services.UploadImage.AddImagePostService(imageDate)
		if err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		imageDate, err := convert.ConvertUploadPostImage(r)
		if err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
			return
		}

		pathImage, err := h.Services.UploadImage.GetImagePostService(imageDate.PostId)
		if err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
			return
		}
		json.WriteJSON(w, http.StatusOK, model.PathImagePost{Path: pathImage})
	default:
		return
	}
}
