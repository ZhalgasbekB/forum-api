package post

import (
	"math"

	"gitea.com/lzhuk/forum/internal/model"
)

type ILikePostRepository interface {
	CreateLikePostRepository(*model.LikePost) error
	GetLikePostRepository(int, int) (*model.LikePost, error)
	GetLikesAndDislikesPostRepository(int) (int, int, error)
	DeleteLikeByPostIdRepository(int, int) error
}

type ILikePostService interface {
	LikePostService(*model.LikePost) error
	GetLikesAndDislikesPostService(int) (int, int, error)
}

type LikePostService struct {
	likePostRepo ILikePostRepository
}

func NewLikePostService(likePostRepo ILikePostRepository) *LikePostService {
	return &LikePostService{
		likePostRepo: likePostRepo,
	}
}

func (l *LikePostService) LikePostService(like *model.LikePost) error {
	oldLike, _ := l.likePostRepo.GetLikePostRepository(like.UserId, like.PostId)
	if oldLike != nil {
		l.likePostRepo.DeleteLikeByPostIdRepository(like.UserId, like.PostId)
		if oldLike.LikeStatus == like.LikeStatus {
			return nil
		}
	}
	return l.likePostRepo.CreateLikePostRepository(like)
}

func (l *LikePostService) GetLikesAndDislikesPostService(post_id int) (int, int, error) {
	likes, dislikes, err := l.likePostRepo.GetLikesAndDislikesPostRepository(post_id)
	if err != nil {
		return -1, -1, err
	}
	return likes, int(math.Abs(float64(dislikes))), nil
}
