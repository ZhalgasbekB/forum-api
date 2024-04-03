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
	CommentByID(context.Context, int) (*model.Comment, error)
	 
}

type ICommentService interface {
	CreateCommentService(context.Context, *model.Comment) error
	UpdateCommentService(context.Context, *model.Comment) error
	DeleteCommentService(context.Context, *model.Comment) error
	CommentByIDService(context.Context, int) (*model.Comment, error)
	 
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
	if err := r.iCommentRepository.CreateComment(ctx, comm); err != nil {
		return err
	}
	return nil
}

func (r *CommentService) UpdateCommentService(ctx context.Context, comm *model.Comment) error {
	comm.UpdatedDate = time.Now()
	if err := r.iCommentRepository.UpdateComment(ctx, comm); err != nil {
		return err
	}
	return nil
}

func (r *CommentService) DeleteCommentService(ctx context.Context, comm *model.Comment) error {
	if err := r.iCommentRepository.DeleteComment(ctx, comm); err != nil {
		return err
	}
	return nil
}

func (r *CommentService) CommentByIDService(ctx context.Context, id int) (*model.Comment, error) {
	comm, err := r.iCommentRepository.CommentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return comm, nil
}

 