package post

import (
	"context"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type PostsRepository interface {
	CreatePostRepository(ctx context.Context, post model.Post) error
	PostsRepository(ctx context.Context) ([]*model.Post, error)
	PostByIdRepository(ctx context.Context, id int) (*model.Post, error)
	PostByUserIdRepository(ctx context.Context, userId int) ([]*model.Post, error)
	UpdatePostByUserIdRepository(ctx context.Context, post model.Post) error
	DeletePostByUserIdRepository(ctx context.Context, deleteModel *model.Post) error
	PostCommentsRepository(context.Context, int) (*model.PostCommentsDTO, error)
}

type IPostsService interface {
	CreatePostService(ctx context.Context, post model.Post) error
	GetAllPostService(ctx context.Context) ([]*model.Post, error)
	GetIdPostService(ctx context.Context, numId int) (*model.Post, error)
	GetUserPostService(ctx context.Context, numId int) ([]*model.Post, error)
	UpdateUserPostService(ctx context.Context, post model.Post) error
	DeleteUserPostService(ctx context.Context, deleteModel *model.Post) error
	CommentsPostService(context.Context, int) (*model.PostCommentsDTO, error)
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
	return p.repo.CreatePostRepository(ctx, post)
}

func (p *PostsService) GetAllPostService(ctx context.Context) ([]*model.Post, error) {
	return p.repo.PostsRepository(ctx)
}

func (p *PostsService) GetIdPostService(ctx context.Context, numId int) (*model.Post, error) {
	return p.repo.PostByIdRepository(ctx, numId)
}

func (p *PostsService) GetUserPostService(ctx context.Context, numId int) ([]*model.Post, error) {
	return p.repo.PostByUserIdRepository(ctx, numId)
}

func (p *PostsService) UpdateUserPostService(ctx context.Context, post model.Post) error {
	return p.repo.UpdatePostByUserIdRepository(ctx, post)
}

func (p *PostsService) DeleteUserPostService(ctx context.Context, deleteModel *model.Post) error {
	return p.repo.DeletePostByUserIdRepository(ctx, deleteModel)
}

func (p *PostsService) CommentsPostService(ctx context.Context, id int) (*model.PostCommentsDTO, error) {
	return p.repo.PostCommentsRepository(ctx, id)
}
