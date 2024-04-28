package comment

import (
	"context"
	"database/sql"

	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

const (
	createCommQuery = `INSERT INTO comments (post_id, user_id, description, created_at, updated_at) VALUES ($1,$2,$3,$4,$5)`
	updateCommQuery = `UPDATE comments SET description = $1, updated_at = $2 WHERE id = $3`
	deleteCommQuery = `DELETE FROM comments WHERE id = $1 AND user_id = $2`

	commentsByPostIdQuery = `SELECT c.id, c.post_id, c.user_id, us.name AS user_name, c.description, c.created_at, c.updated_at,COALESCE(cl.likes, 0) AS likes,COALESCE(cl.dislikes, 0) AS dislikes FROM comments c JOIN   users us ON c.user_id = us.id LEFT JOIN (SELECT comment_id, SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM comments_likes  GROUP BY comment_id) cl ON c.id = cl.comment_id WHERE  c.post_id = $1;`
)

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepo(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}

func (repo *CommentsRepository) CreateComment(ctx context.Context, comm *model.Comment) error {
	if _, err := repo.db.ExecContext(ctx, createCommQuery, comm.Post, comm.User, comm.Description, comm.CreatedDate, comm.UpdatedDate); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) UpdateComment(ctx context.Context, comm *model.Comment) error {
	if _, err := repo.db.ExecContext(ctx, updateCommQuery, comm.Description, comm.UpdatedDate, comm.ID); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) DeleteComment(ctx context.Context, comm *model.Comment) error {
	if _, err := repo.db.ExecContext(ctx, deleteCommQuery, comm.ID, comm.User); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) CommentsPostByIDRepository(ctx context.Context, id int) ([]model.Comment, error) {
	commentsPost := []model.Comment{}
	rows, err := repo.db.QueryContext(ctx, commentsByPostIdQuery, id)
	if err != nil {
		if err == errors.ErrSQLNoRows {
			return nil, errors.ErrNotFoundData
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := model.Comment{}
		if err := rows.Scan(&comment.ID, &comment.Post, &comment.User, &comment.Name, &comment.Description, &comment.CreatedDate, &comment.UpdatedDate, &comment.Like, &comment.Dislike); err != nil {
			return nil, err
		}
		commentsPost = append(commentsPost, comment)
	}
	return commentsPost, nil
}
