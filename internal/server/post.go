package server

import (
	"net/http"
	"strings"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Services.PostsService.GetAllPostService()
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) CreatePosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // SEND TO DB LIKE REQUEST
	if r.Method != http.MethodPost {
		// ERROR STRUCT
		return
	}

	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

	post, err := convert.NewConvertCreatePost(r, session)
	if err != nil {
		return
	}
	h.Services.PostsService.CreatePostService(ctx, *post)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	id, err := convert.ConvertDatePost(r.URL.Path)
	if err != nil {
		return
	}
	post, err := h.Services.PostsService.GetIdPostService(id)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, post)
}

func (h *Handler) UserPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

	postsU, err := h.Services.PostsService.GetUserPostService(session.UserID)
	if err != nil {
		return
	}

	response.WriteJSON(w, http.StatusOK, postsU)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		return
	}
	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

	post, err := convert.NewConvertUpdatePost(r, session)
	if err != nil {
		return
	}
	err = h.Services.PostsService.UpdateUserPostService(*post)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, "VSE NORM BRAT")
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}
	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

	deleteModel, err := convert.NewConvertDeletePost(r, session)
	if err != nil {
		return
	}
	err = h.Services.PostsService.DeleteUserPostService(deleteModel)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, "VSE UDALIL BRAT CHETCO")
}

// Проставление лайка или дизлайка на тему (пост)
func (h *Handler) votePost(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/userd3/post/vote") {
		return
	}
	if r.Method != http.MethodPost {
		return
	}

	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

	postVote, err := convert.NewConvertVote(r, session)
	if err != nil {
		return
	}

	err = h.Services.PostsService.VotePostsService(*postVote)
	if err != nil {
		return
	}
}

// Получение страницы с постами понравившихся пользователю
func (h *Handler) likePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)
	postsVote, err := h.Services.SessionService.CreateSessionService(session.UserID) // NEED ???
	// postsVote, err := h.Services.SessionService.LikePostsService(session.UserID) // NEED ???
	if err != nil {
		return
	}
	for _, v := range []string{"DD"} {
		println(postsVote, v)
	}
}
