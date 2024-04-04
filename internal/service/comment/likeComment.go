package comment

import (
	"gitea.com/lzhuk/forum/internal/model"
)

type ILikeCommentRepository interface {
	CreateLikeCommentRepository(*model.LikeComment) error
	DeleteLikeByCommentIdRepository(int, int) error
	LikeCommentRepository(int, int) (*model.LikeComment, error)
	LikesAndDislikesCommentAllRepository() (map[int][]int, map[int]string, error)
}

type ILikeCommentService interface {
	LikeCommentService(*model.LikeComment) error
	LikesAndDislikesCommentService() (map[int][]int, map[int]string, error)
	PostCommentsService(post_id int) ([]model.Comment, error)
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
	oldLike, _ := l.LikeCommentRepository.LikeCommentRepository(like.UserId, like.CommentId)
	if oldLike != nil {
		l.LikeCommentRepository.DeleteLikeByCommentIdRepository(like.UserId, like.CommentId)
		if oldLike.LikeStatus == like.LikeStatus {
			return nil
		}
	}
	return l.LikeCommentRepository.CreateLikeCommentRepository(like)
}

func (l *LikeCommentService) LikesAndDislikesCommentService() (map[int][]int, map[int]string, error) {
	return l.LikeCommentRepository.LikesAndDislikesCommentAllRepository()
}
