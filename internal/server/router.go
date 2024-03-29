package server

import (
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	//  USERS
	mux.HandleFunc("/register", h.Register) // Страница регистрации (методы GET и POST)
	mux.HandleFunc("/login", h.Login)       // Страница для входа (методы GET и POST)

	
	mux.HandleFunc("/userd3", h.Home)               // Главная страница пользователя (метод GET)
	mux.HandleFunc("/userd3/posts", h.CreatePosts)  // Cтраница создания новой темы (поста) (методы GET и POST)
	mux.HandleFunc("/userd3/myposts", h.UserPosts)  // Страница созданных пользователем тем (постов) (метод GET)
	mux.HandleFunc("/userd3/post/", h.Post)         // Получение страницы с конкретной темой по id (метод GET)
	mux.HandleFunc("/userd3/mypost/", h.UpdatePost) // Изменение и удаление поста (метод PUT)
	mux.HandleFunc("/userd3/mypost/", h.DeletePost) // Изменение и удаление поста (метод DELETE)
	
	
	mux.HandleFunc("/userd3/post/vote", h.votePost) // Проставление лайка или дизлайка на пост (метод POST)
	mux.HandleFunc("/userd3/likeposts", h.likePosts) // ??           // Cтраница понравившихся тем пользователем (метод GET)
	//  MY ROUTES
	mux.HandleFunc("/userd3/commentsCreate", h.CreateComment)
	mux.HandleFunc("/userd3/commentsDelete", h.DeleteComment)
	mux.HandleFunc("/userd3/commentsUpdate", h.UpdateComment)
	mux.HandleFunc("/userd3/comment", h.CommentByID)
	mux.HandleFunc("/userd3/comments", h.Comments)
	mux.HandleFunc("/userd3/postComments", h.Comments)
	return mux
}
