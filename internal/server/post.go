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

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	// id, err := convert.ConvertDatePost(r.URL.Path)
	// if err != nil {
	// 	return
	// }
	post, err := h.Services.PostsService.GetIdPostService(r.Context(), 1)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, post)
}

func (h *Handler) UserPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	user := contextUser(r)
	postsU, err := h.Services.PostsService.GetUserPostService(r.Context(), user.ID)
	if err != nil {
		return
	}
	response.WriteJSON(w, http.StatusOK, postsU)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	fmt.Println("DDDD")
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

// Получение страницы с постами понравившихся пользователю
// func (h *Handler) likePosts(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		return
// 	}
// 	uuid, _ := r.Cookie("CookieUUID")
// 	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)
// 	postsVote, err := h.Services.SessionService.CreateSessionService(session.UserID) // NEED ???
// 	// postsVote, err := h.Services.SessionService.LikePostsService(session.UserID) // NEED ???
// 	if err != nil {
// 		return
// 	}
// 	for _, v := range []string{"DD"} {
// 		println(postsVote, v)
// 	}
// }
// Проставление лайка или дизлайка на тему (пост)
// func (h *Handler) votePost(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodGet)
// 		return
// 	}

// 	uuid, _ := r.Cookie("CookieUUID")
// 	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

// 	postVote, err := convert.NewConvertVote(r, session)
// 	if err != nil {
// 		return
// 	}

// 	err = h.Services.PostsService.VotePostsService(*postVote)
// 	if err != nil {
// 		return
// 	}
// }
