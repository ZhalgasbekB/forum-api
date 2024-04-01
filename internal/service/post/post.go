package post

import (
	"context"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type PostsRepository interface {
	CreatePostRepository(ctx context.Context, post model.Post) error
	AllPostRepository(ctx context.Context) ([]*model.Post, error)
	IdPostRepository(ctx context.Context, id int) (*model.Post, error)
	UserPostRepository(ctx context.Context, userId int) ([]*model.Post, error)
	UpdateUserPostRepository(ctx context.Context, post model.Post) error
	DeleteUserPostRepository(ctx context.Context, deleteModel *model.Post) error
}

type IPostsService interface {
	CreatePostService(ctx context.Context, post model.Post) error
	GetAllPostService(ctx context.Context) ([]*model.Post, error)
	GetIdPostService(ctx context.Context, numId int) (*model.Post, error)
	GetUserPostService(ctx context.Context, numId int) ([]*model.Post, error)
	UpdateUserPostService(ctx context.Context, post model.Post) error
	DeleteUserPostService(ctx context.Context, deleteModel *model.Post) error
}

type PostsService struct {
	repo PostsRepository
}

func NewPostsService(repo PostsRepository) *PostsService {
	return &PostsService{
		repo: repo,
	}
}

func (p *PostsService) CreatePostService(ctx context.Context, post model.Post) error {
	post.CreateDate = time.Now()
	if err := p.repo.CreatePostRepository(ctx, post); err != nil {
		return err
	}
	return nil
}

func (p *PostsService) GetAllPostService(ctx context.Context) ([]*model.Post, error) {
	allPosts, err := p.repo.AllPostRepository(ctx)
	if err != nil {
		return nil, err
	}
	return allPosts, nil
}

func (p *PostsService) GetIdPostService(ctx context.Context, numId int) (*model.Post, error) {
	postId, err := p.repo.IdPostRepository(ctx, numId)
	if err != nil {
		return nil, err
	}
	return postId, nil
}

func (p *PostsService) GetUserPostService(ctx context.Context, numId int) ([]*model.Post, error) {
	userPosts, err := p.repo.UserPostRepository(ctx, numId)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}

func (p *PostsService) UpdateUserPostService(ctx context.Context, post model.Post) error {
	if err := p.repo.UpdateUserPostRepository(ctx, post); err != nil {
		return err
	}
	return nil
}

func (p *PostsService) DeleteUserPostService(ctx context.Context, deleteModel *model.Post) error {
	if err := p.repo.DeleteUserPostRepository(ctx, deleteModel); err != nil {
		return err
	}
	return nil
}
