package post

import (
	"context"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type IPostsRepository interface {
	CreatePostRepository(context.Context, model.Post) error
	UpdatePostByUserIdRepository(context.Context, model.Post) error
	DeletePostByUserIdRepository(context.Context, *model.Post) error

	PostsCategoryRepository(context.Context, string) ([]*model.Post, error)
	PostByIdRepository(ctx context.Context, id int) (*model.Post, error)
	PostByUserIdRepository(ctx context.Context, userId int) ([]*model.Post, error)
	PostsRepository(ctx context.Context) ([]*model.Post, error)
}

type IPostsService interface {
	CreatePostService(ctx context.Context, post model.Post) error
	UpdateUserPostService(ctx context.Context, post model.Post) error
	DeleteUserPostService(ctx context.Context, deleteModel *model.Post) error

	PostsCategoryService(context.Context, string) ([]*model.Post, error)
	GetIdPostService(ctx context.Context, numId int) (*model.Post, error)
	GetUserPostService(ctx context.Context, numId int) ([]*model.Post, error)
	GetAllPostService(ctx context.Context) ([]*model.Post, error)
}

type PostsService struct {
	postsRepository IPostsRepository
}

func NewPostsService(postsRepository IPostsRepository) *PostsService {
	return &PostsService{
		postsRepository: postsRepository,
	}
}

func (p *PostsService) CreatePostService(ctx context.Context, post model.Post) error {
	post.CreateDate = time.Now()
	return p.postsRepository.CreatePostRepository(ctx, post)
}

func (p *PostsService) GetAllPostService(ctx context.Context) ([]*model.Post, error) {
	return p.postsRepository.PostsRepository(ctx)
}

func (p *PostsService) GetIdPostService(ctx context.Context, numId int) (*model.Post, error) {
	return p.postsRepository.PostByIdRepository(ctx, numId)
}

func (p *PostsService) GetUserPostService(ctx context.Context, numId int) ([]*model.Post, error) {
	return p.postsRepository.PostByUserIdRepository(ctx, numId)
}

func (p *PostsService) UpdateUserPostService(ctx context.Context, post model.Post) error {
	return p.postsRepository.UpdatePostByUserIdRepository(ctx, post)
}

func (p *PostsService) DeleteUserPostService(ctx context.Context, deleteModel *model.Post) error {
	return p.postsRepository.DeletePostByUserIdRepository(ctx, deleteModel)
}

func (p *PostsService) PostsCategoryService(ctx context.Context, category string) ([]*model.Post, error) {
	return p.postsRepository.PostsCategoryRepository(ctx, category)
}
