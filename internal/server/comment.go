package server

import (
	"log"
	"net/http"
	"strconv"

	"gitea.com/lzhuk/forum/internal/convert"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	uuid, err := r.Cookie("CookieUUID")
	if err != nil {
		log.Println(err)
		return
	}
	session, err := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)
	if err != nil {
		log.Println(err)
		return
	}
	createComment, err := convert.CreateCommentConvert(r, session)
	if err != nil {
		log.Println(err)
		return
	}
	if err := h.Services.CommentService.CreateCommentService(createComment); err != nil {
		log.Println(err)
		return
	}
	log.Printf("Successfully Created")
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		return
	}
	uuid, err := r.Cookie("CookieUUID")
	if err != nil {
		log.Println(err)
		return
	}
	session, err := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)
	if err != nil {
		log.Println(err)
		return
	}
	deletedComment, err := convert.DeleteCommentConvert(r, session)
	if err := h.Services.CommentService.DeleteCommentService(deletedComment); err != nil {
		log.Println(err)
		return
	}
	log.Printf("Successfully Deleted")
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		return
	}
	uuid, err := r.Cookie("CookieUUID")
	if err != nil {
		log.Println(err)
		return
	}
	session, err := h.Services.SessionService.GetSessionByUUIDService(uuid.Value)
	if err != nil {
		log.Println(err)
		return
	}
	updComment, err := convert.UpdateCommentConvert(r, session)
	if err != nil {
		log.Println(err)
		return
	}
	if err := h.Services.CommentService.UpdateCommentService(updComment); err != nil {
		log.Println(err)
		return
	}
	log.Printf("Successfully Updated")
}

func (h *Handler) CommentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	idS := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println(err)
		return
	}
	comm, err := h.Services.CommentService.CommentByIDService(id)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(comm)
}

func (h *Handler) Comments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	comments, err := h.Services.CommentService.CommentsService()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(comments)
}
