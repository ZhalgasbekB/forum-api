package convert

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"gitea.com/lzhuk/forum/internal/model"
)

func NewConvertCreatePost(r *http.Request, session *model.Sessinon) (*model.CreatePost, error) {
	createPost := &model.CreatePost{}
	if err := json.NewDecoder(r.Body).Decode(createPost); err != nil {
		return nil, err
	}
	return &model.CreatePost{
		UserId:       session.UserID,
		CategoryName: createPost.CategoryName,
		Title:        createPost.Title,
		Discription:  createPost.Discription,
	}, nil
}

func NewConvertUpdatePost(r *http.Request, session *model.Sessinon) (*model.UpdatePost, error) {
	numIdPost, err := ConvertDatePost(r.URL.Path)
	if err != nil {
		return nil, err
	}
	updatePost := &model.UpdatePost{}
	if err := json.NewDecoder(r.Body).Decode(updatePost); err != nil {
		return nil, err
	}
	return &model.UpdatePost{
		PostId:      numIdPost,
		Discription: updatePost.Discription,
		UserId:      session.UserID,
	}, nil
}

func NewConvertDeletePost(r *http.Request, session *model.Sessinon) (*model.DeletePost, error) {
	numIdPost, err := ConvertDatePost(r.URL.Path)
	if err != nil {
		return nil, err
	}
	deletePost := &model.DeletePost{}
	if err := json.NewDecoder(r.Body).Decode(deletePost); err != nil {
		return nil, err
	}
	return &model.DeletePost{
		UserId: session.UserID,
		PostId: numIdPost,
	}, nil
}


func ConvertDatePost(path string) (int, error) {
	switch {
	case strings.HasPrefix(path, "/userd3/post/"):
		idStr := strings.ReplaceAll(path, "/userd3/post/", "")

		numId, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, err
		}
		return numId, nil

	case strings.HasPrefix(path, "/userd3/mypost/"):
		idStr := strings.ReplaceAll(path, "/userd3/mypost/", "")

		numId, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, err
		}
		return numId, nil

	}
	return 0, nil
}
