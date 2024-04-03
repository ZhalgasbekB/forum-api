package comment

import (
	"gitea.com/lzhuk/forum/internal/model"
)

type ILikeCommentRepository interface {
	CreateLikeCommentRepository(*model.LikeComment) error
	DeleteLikeByCommentIdRepository(int, int) error
	GetLikeCommentRepository(int, int) (*model.LikeComment, error)
	GetLikesAndDislikesCommentAllRepository() (map[int][]int, error)
	GetUserLikedCommentRepository(*model.LikeComment) error
}

type ILikeCommentService interface {
	LikeCommentService(*model.LikeComment) error
	GetLikesAndDislikesCommentService() (map[int][]int, error)
}

type LikeCommentService struct {
	LikeCommentRepository ILikeCommentRepository
}

func NewLikeCommentService(LikeCommentRepository ILikeCommentRepository) *LikeCommentService {
	return &LikeCommentService{
		LikeCommentRepository: LikeCommentRepository,
	}
}

func (l *LikeCommentService) LikeCommentService(like *model.LikeComment) error {
	oldLike, _ := l.LikeCommentRepository.GetLikeCommentRepository(like.UserId, like.CommentId)
	if oldLike != nil {
		l.LikeCommentRepository.DeleteLikeByCommentIdRepository(like.UserId, like.CommentId)
		if oldLike.LikeStatus == like.LikeStatus {
			return nil
		}
	}
	return l.LikeCommentRepository.CreateLikeCommentRepository(like)
}

func (l *LikeCommentService) GetLikesAndDislikesCommentService() (map[int][]int, error) {
	return l.LikeCommentRepository.GetLikesAndDislikesCommentAllRepository()
}
