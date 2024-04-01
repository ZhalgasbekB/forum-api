package post

import (
	"context"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type PostsRepository interface {
	CreatePostRepository(ctx context.Context, post model.CreatePost) error
	AllPostRepository() ([]*model.Post, error)
	IdPostRepository(id int) (*model.Post, error)
	UserPostRepository(userId int) ([]*model.Post, error)
	UpdateUserPostRepository(post model.UpdatePost) error
	DeleteUserPostRepository(deleteModel *model.DeletePost) error	 
}

type IPostsService interface {
	CreatePostService(ctx context.Context, post model.CreatePost) error
	GetAllPostService() ([]*model.Post, error)
	GetIdPostService(numId int) (*model.Post, error)
	GetUserPostService(numId int) ([]*model.Post, error)
	UpdateUserPostService(post model.UpdatePost) error
	DeleteUserPostService(deleteModel *model.DeletePost) error	 
}

type PostsService struct {
	repo PostsRepository
}

func NewPostsService(repo PostsRepository) *PostsService {
	return &PostsService{
		repo: repo,
	}
}

func (p *PostsService) CreatePostService(ctx context.Context, post model.CreatePost) error {
	post.CreateDate = time.Now()
	if err := p.repo.CreatePostRepository(ctx, post); err != nil {
		return err
	}

	return nil
}

func (p *PostsService) GetAllPostService() ([]*model.Post, error) {
	allPosts, err := p.repo.AllPostRepository()
	if err != nil {
		return nil, err
	}
	return allPosts, nil
}

func (p *PostsService) GetIdPostService(numId int) (*model.Post, error) {
	postId, err := p.repo.IdPostRepository(numId)
	if err != nil {
		return nil, err
	}
	return postId, nil
}

func (p *PostsService) GetUserPostService(numId int) ([]*model.Post, error) {
	userPosts, err := p.repo.UserPostRepository(numId)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}

func (p *PostsService) UpdateUserPostService(post model.UpdatePost) error {
	post.CreateDate = time.Now()
	err := p.repo.UpdateUserPostRepository(post)
	if err != nil {
		return err
	}
	
	return nil
}

func (p *PostsService) DeleteUserPostService(deleteModel *model.DeletePost) error {
	err := p.repo.DeleteUserPostRepository(deleteModel)
	if err != nil {
		return err
	}
	return nil
}

 