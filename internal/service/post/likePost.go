package post

import (
	"fmt"
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
	GetLikesAndDislikesPostService(int) error
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

func (l *LikePostService) GetLikesAndDislikesPostService(post_id int) error {
	likes, dislikes, err := l.likePostRepo.GetLikesAndDislikesPostRepository(post_id)
	if err != nil {
		return err
	}
	fmt.Println("DDDDD")
	fmt.Println(likes, int(math.Abs(float64(dislikes))))
	return nil
}
