package post

import "gitea.com/lzhuk/forum/internal/model"

type ILikePostRepository interface {
	LikePostsRepository(userId int) ([]*model.Post, error)
	VotePostsRepository(post model.Vote) error
	CheckVotePost(post model.Vote) (string, error)
	DeleteVotePost(post model.Vote) error
}

type ILikePostService interface {
	LikePostsService(userId int) ([]*model.Post, error)
	VotePostsService(post model.Vote) error
}

type LikePostService struct {
	repo ILikePostRepository
}

func NewLikePostsService(repo ILikePostRepository) *LikePostService {
	return &LikePostService{
		repo: repo,
	}
}

func (p *LikePostService) VotePostsService(post model.Vote) error {
	check, err := p.repo.CheckVotePost(post)
	if check == "yes" {
		err = p.repo.DeleteVotePost(post)
		if err != nil {
			return err
		}
		return nil
	}

	err = p.repo.VotePostsRepository(post)
	if err != nil {
		return err
	}
	return nil
}

func (p *LikePostService) LikePostsService(userId int) ([]*model.Post, error) {
	votePosts, err := p.repo.LikePostsRepository(userId)
	if err != nil {
		return nil, err
	}
	return votePosts, nil
}
