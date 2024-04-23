package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitea.com/lzhuk/forum/internal/model"
	comment2 "gitea.com/lzhuk/forum/internal/repository/comment"
	post2 "gitea.com/lzhuk/forum/internal/repository/post"
	user2 "gitea.com/lzhuk/forum/internal/repository/user"
	"gitea.com/lzhuk/forum/internal/server"
	"gitea.com/lzhuk/forum/internal/service"
	"gitea.com/lzhuk/forum/internal/service/comment"
	"gitea.com/lzhuk/forum/internal/service/post"
	"gitea.com/lzhuk/forum/internal/service/user"
	"gitea.com/lzhuk/forum/pkg/config"
	"gitea.com/lzhuk/forum/pkg/db/driver"
)

type KeyUser string

func TestStartPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}
	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.Home)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/userd3", nil)
	if err != nil {
		t.Fatal(err)
	}
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}
	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.Home)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"post_id":1,"user_id":1,"category_name":"SSSSS","title":"CHANGENAME1","description":"DDUPDATE333","create_at":"2024-04-03T10:05:22.067903164+06:00","name":"Saken","likes":1,"dislikes":0},{"post_id":2,"user_id":1,"category_name":"reurowur","title":"one1","description":"fdgjklkvsklhfl","create_at":"2024-04-03T12:21:46.563880426+06:00","name":"Saken","likes":0,"dislikes":0},{"post_id":3,"user_id":1,"category_name":"reurowur","title":"one2","description":"fdgjklkvsklhfl","create_at":"2024-04-03T12:23:06.494829715+06:00","name":"Saken","likes":0,"dislikes":1},{"post_id":4,"user_id":1,"category_name":"reurowur","title":"one3","description":"fdgjklkvsklhfl","create_at":"2024-04-03T12:24:14.415795836+06:00","name":"Saken","likes":0,"dislikes":0},{"post_id":5,"user_id":1,"category_name":"reurowur","title":"one3","description":"fdgjklkvsklhfl","create_at":"2024-04-04T09:50:06.026867507+06:00","name":"Saken","likes":0,"dislikes":0},{"post_id":6,"user_id":1,"category_name":"reurowurPWDPPWPDPW","title":"one3","description":"fdgjklkvsklhfl","create_at":"2024-04-04T09:50:25.359260573+06:00","name":"Saken","likes":0,"dislikes":0},{"post_id":7,"user_id":1,"category_name":"reurowurPWDPPWPDPW","title":"one35453543","description":"fdgjklkvsklhfl","create_at":"2024-04-04T09:50:39.338054155+06:00","name":"Saken","likes":0,"dislikes":0},{"post_id":8,"user_id":1,"category_name":"","title":"fgvnmfn","description":"vsddsd","create_at":"2024-04-04T10:27:51.250234384+06:00","name":"Saken","likes":0,"dislikes":0},{"post_id":9,"user_id":1,"category_name":"","title":"vckgf","description":"bccbc","create_at":"2024-04-04T10:50:17.50870632+06:00","name":"Saken","likes":0,"dislikes":0}]
`
	if rr.Body.String() != expected {
		println(len(expected))
		println(len(rr.Body.String()))
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRegisterExistingAccount(t *testing.T) {
	payload := map[string]string{
		"name":     "Danial",
		"email":    "danial@gmail.com", // Should be an existing gmail
		"password": "1234512345",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/register", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}
	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.Register)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "{\"status\":500,\"message\":\"Email already exist\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRegisterNewAccount(t *testing.T) {
	payload := map[string]string{
		"name":     "Danial",
		"email":    "danial@gmail.com.somerandom.comx", // Should be random always.
		"password": "1234512345",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/register", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}

	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.Register)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestLoginIncorrectCredentials(t *testing.T) {
	payload := map[string]string{
		"email":    "danial@gmail.com",
		"password": "someincorrectpassword",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}

	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	expectedMessage := "Invalid Credentials"
	json.Unmarshal(rr.Body.Bytes(), &response)
	if value, ok := response["message"]; !ok || value != expectedMessage {
		fmt.Println(response)
		t.Errorf("handler returned unexpected body, missing 'name'")
	}
}

func TestLoginCorrectCredentials(t *testing.T) {
	payload := map[string]string{
		"email":    "danial@gmail.com",
		"password": "1234512345",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}

	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	if _, ok := response["name"]; !ok {
		fmt.Println(response)
		t.Errorf("handler returned unexpected body, missing 'name'")
	}
	if _, ok := response["email"]; !ok {
		t.Errorf("handler returned unexpected body, missing 'email'")
	}
}

func TestLikePosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/userd3/likeposts", nil)
	if err != nil {
		t.Fatal(err)
	}

	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}

	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.LikePosts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMyPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/userd3/myposts", nil)
	if err != nil {
		t.Fatal(err)
	}

	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}

	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.PostsUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreatePosts(t *testing.T) {
	payload := map[string]string{
		"cookie_uuid":   "ad07db8a-a88b-4a31-b37e-49c398886756",
		"category_name": "Станции",
		"title":         "Станция Люберцы",
		"discription":   "Новые поезда",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/userd3/posts", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	user := &model.User{
		ID:    1337,
		Name:  "Danial",
		Email: "danial@example.com",
	}

	req = setUserContext(req, user)
	cookie := &http.Cookie{
		Name:  "UserData",
		Value: "ad07db8a-a88b-4a31-b37e-49c398886756",
	}
	req.AddCookie(cookie)
	services, rr, err := initServices()
	if err != nil {
		t.Errorf("Error initializing services: %s", err)
	}

	h := server.NewHandler(services)
	handler := http.HandlerFunc(h.CreatePosts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func initServices() (service.Service, *httptest.ResponseRecorder, error) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		return service.Service{}, nil, errors.New("Ошибка загрузки конфигурации")
	}

	db, err := driver.NewDB(cfg)
	if err != nil {
		log.Println("Ошибка инциализации базы данных %w", err)
		return service.Service{}, nil, errors.New("Ошибка загрузки конфигурации")
	}

	usersRepo := user2.NewUserRepo(db)
	sessionRepo := user2.NewSessionRepository(db)
	postsRepo := post2.NewPostsRepo(db)
	likePostRepo := post2.NewLikePostRepository(db)
	commentsRepo := comment2.NewCommentsRepo(db)
	likeCommentsRepo := comment2.NewLikeCommentRepository(db)

	usersService := user.NewUserService(usersRepo)
	sessionsService := user.NewSessionService(sessionRepo)
	postsService := post.NewPostsService(postsRepo)
	likePostsService := post.NewLikePostService(likePostRepo)
	commentsService := comment.NewCommentsService(commentsRepo)
	likeCommentsService := comment.NewLikeCommentService(likeCommentsRepo)

	services := service.NewService(usersService, sessionsService, postsService, commentsService, likePostsService, likeCommentsService)

	rr := httptest.NewRecorder()
	return services, rr, nil
}

// func setUserContext(r *http.Request, user *model.User) *http.Request {
// 	ctx := context.WithValue(r.Context(), KeyUser("UserData"), user)
// 	return r.WithContext(ctx)
// }
// func addContext(req *http.Request, key, value interface{}) *http.Request {
// 	ctx := context.WithValue(req.Context(), key, value)
// 	return req.WithContext(ctx)
// }
