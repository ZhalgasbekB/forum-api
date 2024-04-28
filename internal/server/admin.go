package server

import (
	"encoding/json"
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	hh "gitea.com/lzhuk/forum/internal/helpers/json"
	"gitea.com/lzhuk/forum/internal/model"

	"gitea.com/lzhuk/forum/internal/errors"
)

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	users, err := h.Services.Admin.UsersService()
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	hh.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) AdminChangeRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uRole, err := convert.UpdateRole(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.UpdateUserService(uRole); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	hh.WriteJSON(w, http.StatusOK, uRole.Role)
}

func (h *Handler) AdminCreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	category := &model.CategoryDTO{}

	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.CreateCategoryService(category.CategoryName); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AdminDeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	category := &model.CategoryDTO{}

	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.DeleteCategoryService(category.CategoryName); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := convert.DeleteUser(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.DeleteUserService(id); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id, err := convert.DeletePost(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.DeletePostService(id); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id, err := convert.DeleteComment(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.DeleteCommentService(id); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

/// Moderator func (???)

func (h *Handler) ModeratorReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	report := &model.ReportCreateDTO{}
	if err := json.NewDecoder(r.Body).Decode(report); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.CreateReportService(report); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AdminReports(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	reports, err := h.Services.Admin.ReportsModeratorService()
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	hh.WriteJSON(w, http.StatusOK, reports)
}

func (h *Handler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	update := &model.ReportResponseDTO{}
	if err := json.NewDecoder(r.Body).Decode(update); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.UpdateReportService(update); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

// WRITE

// USER UP (???)

func (h *Handler) UserUpRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user := &model.WantsDTO{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.UserWantsService(user); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UsersWants(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	users, err := h.Services.Admin.UsersWantsService()
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	hh.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) UserWantRoleAdminResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	temp, err := convert.AdminResponse(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.UserWantRoleAdminResponseService(temp); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	// if temp.Status == 1 {
	// 	t := &model.User{ID: temp.UserID, Role: roles.MODERATOR}
	// 	if err := h.Services.Admin.UpdateUserService(t); err != nil {
	// 		log.Println(err)
	// 		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// }

	// UPDATE STATUS 1 OR -1 //
}
