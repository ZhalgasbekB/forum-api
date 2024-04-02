package post

import (
	"fmt"
	"math"

	"gitea.com/lzhuk/forum/internal/model"
)

type ILikePostRepository interface {
	CreateLikePostRepository(*model.LikePost) error
	UpdateLikePostRepository(*model.LikePost) error
	GetLikePostRepository(int, int) (*model.LikePost, error)
	GetLikesAndDislikesPostRepository(int) (int, int, error)
}

type ILikePostService interface {
	LikePostService(like *model.LikePost) error
	GetLikesAndDislikesPostService(like *model.LikePost) error
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
		if err := l.likePostRepo.UpdateLikePostRepository(like); err != nil {
			return err
		}
	}
	if err := l.likePostRepo.CreateLikePostRepository(like); err != nil {
		return err
	}
	return nil
}

func (l *LikePostService) GetLikesAndDislikesPostService(like *model.LikePost) error {
	likes, dislikes, err := l.likePostRepo.GetLikesAndDislikesPostRepository(like.PostId)
	if err != nil {
		return err
	}
	fmt.Println(likes, int(math.Abs(float64(dislikes))))
	return nil
}
