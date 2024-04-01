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
		Discription:  createPost.Discription,
	}, nil
}

func ConvertUpdatePost(r *http.Request, user_id int) (*model.Post, error) {
	idS := r.URL.Query().Get("id")
	numId, err := strconv.Atoi(idS)
	if err != nil {
		return nil, err
	}

	updatePost := &model.UpdatePostDTO{}
	if err := json.NewDecoder(r.Body).Decode(updatePost); err != nil {
		return nil, err
	}
	return &model.Post{
		PostId:      numId,
		UserId:      user_id,
		Title:       updatePost.Title,
		Discription: updatePost.Discription,
	}, nil
}

func ConvertDeletePost(r *http.Request, user_id int) (*model.DeletePost, error) {
	idS := r.URL.Query().Get("id")
	numId, err := strconv.Atoi(idS)
	if err != nil {
		return nil, err
	}
	deletePost := &model.DeletePost{}
	if err := json.NewDecoder(r.Body).Decode(deletePost); err != nil {
		return nil, err
	}
	return &model.DeletePost{
		UserId: user_id,
		PostId: numId,
	}, nil
}
