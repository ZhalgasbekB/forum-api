package comment

import (
	"context"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type ICommentRepository interface {
	CreateComment(context.Context, *model.Comment) error
	UpdateComment(context.Context, *model.Comment) error
	DeleteComment(context.Context, *model.Comment) error
	CommentsName(ctx context.Context) (map[int]string, error)
}

type ICommentService interface {
	CreateCommentService(context.Context, *model.Comment) error
	UpdateCommentService(context.Context, *model.Comment) error
	DeleteCommentService(context.Context, *model.Comment) error
	CommentsNameService(ctx context.Context) (map[int]string, error)
}

type CommentService struct {
	iCommentRepository ICommentRepository
}

func NewCommentsService(iCommentRepository ICommentRepository) *CommentService {
	return &CommentService{iCommentRepository: iCommentRepository}
}

func (r *CommentService) CreateCommentService(ctx context.Context, comm *model.Comment) error {
	comm.CreatedDate = time.Now()
	comm.UpdatedDate = time.Now()
	return r.iCommentRepository.CreateComment(ctx, comm)
}

func (r *CommentService) UpdateCommentService(ctx context.Context, comm *model.Comment) error {
	comm.UpdatedDate = time.Now()
	return r.iCommentRepository.UpdateComment(ctx, comm)
}

func (r *CommentService) DeleteCommentService(ctx context.Context, comm *model.Comment) error {
	return r.iCommentRepository.DeleteComment(ctx, comm)
}

func (r *CommentService) CommentsNameService(ctx context.Context) (map[int]string, error) {
	return r.iCommentRepository.CommentsName(ctx)
}
