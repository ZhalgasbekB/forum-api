package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

///// home page GET for getting all posts

// Cтраница создания новой темы (поста) (методы GET и POST)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	allPost, err := h.Services.PostsService.GetAllPostService()
	if err != nil {
		return
	}

	// for _, v := range allPost {
	// 	fmt.Println(v)
	// }
	response.WriteJSON(w, 200, allPost)
}

func (h *Handler) createPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/userd3/posts" {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		allPost, err := h.Services.PostsService.GetAllPostService()
		if err != nil {
			return
		}
		for _, v := range allPost {
			fmt.Println(v)
		}

		break
	case http.MethodPost:
		uuid, _ := r.Cookie("CookieUUID")
		session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

		post, err := convert.NewConvertCreatePost(r, session)
		if err != nil {
			return
		}
		h.Services.PostsService.CreatePostService(ctx, *post)
	default:
		break
	}
}

// Получение страницы с конкретной темой по id (метод GET)
func (h *Handler) getPost(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/userd3/post/") {
		return
	}
	if r.Method != http.MethodGet {
		return
	}
	numId, err := convert.ConvertDatePost(r.URL.Path)
	if err != nil {
		return
	}
	postId, err := h.Services.PostsService.GetIdPostService(numId)
	if err != nil {
		return
	}
	fmt.Println(postId)
}

// Получение страницы с темой (постами) созданных конкретным пользователем
func (h *Handler) myPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/userd3/myposts" {
		return
	}
	if r.Method != http.MethodGet {
		return
	}

	uuid, _ := r.Cookie("CookieUUID")
	session, _ := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)

	userPosts, err := h.Services.PostsService.GetUserPostService(session.UserID)
	if err != nil {
		return
	}

	for _, v := range userPosts {
		fmt.Println(v)
	}
}

// Получение страницы с постами понравившихся пользователю
func (h *Handler) likePosts(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	return
	// }
	// uuid, _ := r.Cookie("CookieUUID")
	// session, _ := h.sessionService.GetSessionByUUIDService(uuid.Value)

	// postsVote, err := h.sessionService.LikePostsService(session.UserID)
	// if err != nil {
	// 	return
	// }
	// for _, v := range postsVote {
	// 	fmt.Println(v)
	// }
}

// Изменение или удаления описания и текста раннее созданной темы
func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/userd3/mypost/") {
		return
	}
	switch r.Method {
	case http.MethodPut:

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
	case http.MethodDelete:

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
	}
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
