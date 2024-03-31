package post

import (
	"context"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type PostsRepository interface {
	CreatePostRepo(ctx context.Context, post model.CreatePost) error
	GetAllPostRepo() ([]*model.Post, error)
	GetIdPostRepo(id int) (*model.Post, error)
	GetUserPostRepo(userId int) ([]*model.Post, error)
	LikePostsRepo(userId int) ([]*model.Post, error)
	UpdateUserPostRepo(post model.UpdatePost) error
	DeleteUserPostRepo(deleteModel *model.DeletePost) error
	VotePostsRepo(post model.Vote) error

	CheckVotePost(post model.Vote) (string, error)
	DeleteVotePost(post model.Vote) error
}

type IPostsService interface {
	CreatePostService(ctx context.Context, post model.CreatePost) error
	GetAllPostService() ([]*model.Post, error)
	GetIdPostService(numId int) (*model.Post, error)
	GetUserPostService(numId int) ([]*model.Post, error)
	LikePostsService(userId int) ([]*model.Post, error)
	UpdateUserPostService(post model.UpdatePost) error
	DeleteUserPostService(deleteModel *model.DeletePost) error
	VotePostsService(post model.Vote) error
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
	if err := p.repo.CreatePostRepo(ctx, post); err != nil {
		return err
	}

	return nil
}

func (p *PostsService) GetAllPostService() ([]*model.Post, error) {
	allPosts, err := p.repo.GetAllPostRepo()
	if err != nil {
		return nil, err
	}
	return allPosts, nil
}

func (p *PostsService) GetIdPostService(numId int) (*model.Post, error) {
	postId, err := p.repo.GetIdPostRepo(numId)
	if err != nil {
		return nil, err
	}
	return postId, nil
}

func (p *PostsService) GetUserPostService(numId int) ([]*model.Post, error) {
	userPosts, err := p.repo.GetUserPostRepo(numId)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}

func (p *PostsService) LikePostsService(userId int) ([]*model.Post, error) {
	votePosts, err := p.repo.LikePostsRepo(userId)
	if err != nil {
		return nil, err
	}
	return votePosts, nil
}

func (p *PostsService) UpdateUserPostService(post model.UpdatePost) error {
	post.CreateDate = time.Now()
	err := p.repo.UpdateUserPostRepo(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostsService) DeleteUserPostService(deleteModel *model.DeletePost) error {
	err := p.repo.DeleteUserPostRepo(deleteModel)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostsService) VotePostsService(post model.Vote) error {
	// Бизнес-логика на проверку голоса для темы в БД
	// при наличии голоса пользователя голос убирается
	check, err := p.repo.CheckVotePost(post)
	if check == "yes" {
		err = p.repo.DeleteVotePost(post)
		if err != nil {
			return err
		}
		return nil
	}

	err = p.repo.VotePostsRepo(post)
	if err != nil {
		return err
	}
	return nil
}
