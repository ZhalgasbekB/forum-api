package server

import (
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", h.Register) // Страница регистрации (методы GET и POST)
	mux.HandleFunc("/login", h.Login)       // Страница для входа (методы GET и POST)

	mux.HandleFunc("/userd3", h.Home)                                                            // Главная страница пользователя (метод GET)
	mux.Handle("/userd3/posts", h.RequiredAuthentication(http.HandlerFunc(h.CreatePosts)))       // Cтраница создания новой темы (поста) (методы GET и POST)
	mux.Handle("/userd3/myposts", h.RequiredAuthentication(http.HandlerFunc(h.UserPosts)))       // Страница созданных пользователем тем (постов) (метод GET)
	mux.Handle("/userd3/post/", h.RequiredAuthentication(http.HandlerFunc(h.Post)))              // Получение страницы с конкретной темой по id (метод GET)
	mux.Handle("/userd3/mypostUpdate", h.RequiredAuthentication(http.HandlerFunc(h.UpdatePost))) // Изменение и удаление поста (метод PUT)
	mux.Handle("/userd3/mypostDelete", h.RequiredAuthentication(http.HandlerFunc(h.DeletePost))) // Изменение и удаление поста (метод DELETE)

	mux.HandleFunc("/userd3/post/vote", h.votePost)  // Проставление лайка или дизлайка на пост (метод POST)
	mux.HandleFunc("/userd3/likeposts", h.likePosts) // ??           // Cтраница понравившихся тем пользователем (метод GET)
	//  MY ROUTES
	mux.Handle("/userd3/postComments", h.RequiredAuthentication(http.HandlerFunc(h.PostComments)))
	mux.Handle("/userd3/commentsCreate", h.RequiredAuthentication(http.HandlerFunc(h.CreateComment)))
	mux.Handle("/userd3/commentsDelete", h.RequiredAuthentication(http.HandlerFunc(h.DeleteComment)))
	mux.Handle("/userd3/commentsUpdate", h.RequiredAuthentication(http.HandlerFunc(h.UpdateComment)))
	mux.Handle("/userd3/comment", h.RequiredAuthentication(http.HandlerFunc(h.CommentByID)))
	mux.HandleFunc("/userd3/comments", h.Comments)

	h.IsAuthenticated(mux) // Check Authentication
	return mux
}
