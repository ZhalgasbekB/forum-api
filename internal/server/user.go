package server

import (
	"fmt"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/cookies"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userReq, err := convert.UserLoginRequestBody(r)
	user, err := h.Services.UserService.UserByEmailService(userReq.Email, userReq.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := h.Services.SessionService.CreateSessionService(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	response.WriteJSON(w, http.StatusOK, cookies.CreateCookie(w, session))
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := convert.UserRegisterRequestBody(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := h.Services.UserService.CreateUserService(user); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User Successfully Created")
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	uuid, err := convert.UUID(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := h.Services.SessionService.DeleteSessionService(uuid.UUID); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User Successfully Logout")
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
}
