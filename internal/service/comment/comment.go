package comment

import (
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type ICommentRepository interface {
	CreateComment(*model.Comment) error
	UpdateComment(*model.Comment) error
	DeleteComment(*model.Comment) error
	CommentByID(int) (*model.Comment, error)
	Comments() ([]model.Comment, error)
}
type ICommentService interface {
	CreateCommentService(*model.Comment) error
	UpdateCommentService(*model.Comment) error
	DeleteCommentService(*model.Comment) error
	CommentByIDService(id int) (*model.Comment, error)
	CommentsService() ([]model.Comment, error)
}

type CommentService struct {
	iCommentRepository ICommentRepository
}

func NewCommentsService(iCommentRepository ICommentRepository) *CommentService {
	return &CommentService{iCommentRepository: iCommentRepository}
}

func (r *CommentService) CreateCommentService(comm *model.Comment) error {
	comm.CreatedDate = time.Now()
	comm.UpdatedDate = time.Now()
	if err := r.iCommentRepository.CreateComment(comm); err != nil {
		return err
	}
	return nil
}

func (r *CommentService) UpdateCommentService(comm *model.Comment) error {
	comm.UpdatedDate = time.Now()
	if err := r.iCommentRepository.UpdateComment(comm); err != nil {
		return err
	}
	return nil
}

func (r *CommentService) DeleteCommentService(comm *model.Comment) error {
	if err := r.iCommentRepository.DeleteComment(comm); err != nil {
		return err
	}
	return nil
}

func (r *CommentService) CommentByIDService(id int) (*model.Comment, error) {
	comm, err := r.iCommentRepository.CommentByID(id)
	if err != nil {
		return nil, err
	}
	return comm, nil
}

func (r *CommentService) CommentsService() ([]model.Comment, error) {
	comments, err := r.iCommentRepository.Comments()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
