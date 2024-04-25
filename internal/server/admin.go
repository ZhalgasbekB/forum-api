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

// FOR CHECKING  ADMIN
// user := contextUser(r)
// if user.Role != roles.ADMIN {
// 	log.Println("Not Admin")
// 	errors.ErrorSend(w, http.StatusInternalServerError, "Because You Are Not Admin")
// 	return
// }
//

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
	hh.WriteJSON(w, http.StatusOK, id)
}

func (h *Handler) AdminUpdateAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := convert.UpdateUserAdmin(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.UpdateUserNewDateService(user); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

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

	
}
