package server

import (
	"log"
	"net/http"
	"strconv"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodGet)
		return
	}
	user := userFromContext(r)
	createComment, err := convert.CreateCommentConvert(r, user.ID)
	if err != nil {
		log.Println(err)
		return
	}
	if err := h.Services.CommentService.CreateCommentService(createComment); err != nil {
		log.Println(err)
		return
	}
	response.WriteJSON(w, http.StatusOK, "Successfully Created")
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		return
	}
	user := userFromContext(r)
	deletedComment, err := convert.DeleteCommentConvert(r, user.ID)
	if err != nil {
		log.Println(err)
		return
	}
	if err := h.Services.CommentService.DeleteCommentService(deletedComment); err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		return
	}
	user := userFromContext(r)
	updComment, err := convert.UpdateCommentConvert(r, user.ID)
	if err != nil {
		log.Println(err)
		return
	}
	if err := h.Services.CommentService.UpdateCommentService(updComment); err != nil {
		log.Println(err)
		return
	}
	response.WriteJSON(w, http.StatusOK, "Successfully Updated")
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
	response.WriteJSON(w, http.StatusOK, comm)
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
	response.WriteJSON(w, http.StatusOK, comments)
}
