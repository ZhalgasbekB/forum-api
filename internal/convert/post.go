package convert

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitea.com/lzhuk/forum/internal/model"
)

func ConvertCreatePost(r *http.Request, user_id int) (*model.Post, error) {
	createPost := &model.CreatePostDTO{}
	if err := json.NewDecoder(r.Body).Decode(createPost); err != nil {
		return nil, err
	}
	return &model.Post{
		UserId:       user_id,
		CategoryName: createPost.CategoryName,
		Title:        createPost.Title,
		Description:  createPost.Description,
	}, nil
}

func ConvertUpdatePost(r *http.Request, user_id int) (*model.Post, error) {
	postID, err := ConvertParamID(r)
	if err != nil {
		return nil, err
	}

	updatePost := &model.UpdatePostDTO{}
	if err := json.NewDecoder(r.Body).Decode(updatePost); err != nil {
		return nil, err
	}
	return &model.Post{
		PostId:      postID,
		UserId:      user_id,
		Title:       updatePost.Title,
		Description: updatePost.Description,
	}, nil
}

func ConvertDeletePost(r *http.Request, user_id int) (*model.Post, error) {
	postID, err := ConvertParamID(r)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		UserId: user_id,
		PostId: postID,
	}, nil
}

func ConvertParamID(r *http.Request) (int, error) {
	idS := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		return -1, err
	}
	return id, err
}
