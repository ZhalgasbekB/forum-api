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
	CommentsByPostId(int) ([]model.Comment, error)
	LikesCommentsByPostRepository(int) (map[int][]int, error)
	PostCommentsRepository(context.Context, int) (*model.PostCommentsDTO, error)
}

type ICommentService interface {
	CreateCommentService(context.Context, *model.Comment) error
	UpdateCommentService(context.Context, *model.Comment) error
	DeleteCommentService(context.Context, *model.Comment) error

	CommentsLikesNames(context.Context, int) (*model.PostCommentsDTO, error)
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

func (r *CommentService) CommentsLikesNames(ctx context.Context, post_id int) (*model.PostCommentsDTO, error) {
	postUname, err := r.iCommentRepository.PostCommentsRepository(ctx, post_id)

	commentsPost, err := r.iCommentRepository.CommentsByPostId(post_id)
	if err != nil {
		return nil, err
	}

	commentName, err := r.iCommentRepository.CommentsName(ctx)
	if err != nil {
		return nil, err
	}

	comm, err := r.iCommentRepository.LikesCommentsByPostRepository(post_id)
	if err != nil {
		return nil, err
	}

	for i, v := range commentsPost {
		for k1, v1 := range commentName {
			if v.ID == k1 {
				commentsPost[i].Name = v1
			}
		}
		for k2, v2 := range comm {
			if v.ID == k2 {
				commentsPost[i].Like = v2[0]
				commentsPost[i].Dislike = v2[1]
			}
		}

	}
	postUname.Comments = commentsPost
	return postUname, nil
}
