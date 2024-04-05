package comment

import (
	"context"
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

const (
	createCommQuery   = `INSERT INTO comments (post_id, user_id, description, created_at, updated_at) VALUES ($1,$2,$3,$4,$5)`
	updateCommQuery   = `UPDATE comments SET description = $1, updated_at = $2 WHERE id = $3`
	deleteCommQuery   = `DELETE FROM comments WHERE id = $1 AND user_id = $2`
	commentsNameQuery = `SELECT c.id, us.name FROM comments c JOIN users us ON c.user_id  = us.id`

	postCommentsQuery = `SELECT c.id, c.post_id, c.user_id, c.description, c.created_at, c.updated_at FROM posts p JOIN comments c ON c.post_id = p.id WHERE p.id = $1`               // CHECK
	postByIdQuery     = `SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name FROM posts ps JOIN users u ON ps.user_id = u.id WHERE ps.id = $1` // CHECK

	likesCommentsByPost = `SELECT comment_id, SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM comments_likes c JOIN comments co ON c.comment_id = co.id JOIN posts p ON p.id = co.post_id  WHERE p.id = $1 GROUP BY comment_id`
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

func (repo *CommentsRepository) CommentsName(ctx context.Context) (map[int]string, error) {
	commentsName := map[int]string{}
	rows, err := repo.db.QueryContext(ctx, commentsNameQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		commentsName[id] = name
	}
	return commentsName, nil
}

func (repo *CommentsRepository) CommentsByPostId(post_id int) ([]model.Comment, error) {
	commentsPost := []model.Comment{}
	rows, err := repo.db.Query(postCommentsQuery, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.Post, &comment.User, &comment.Description, &comment.CreatedDate, &comment.UpdatedDate); err != nil {
			return nil, err
		}
		commentsPost = append(commentsPost, comment)
	}
	return commentsPost, nil
}

func (repo *CommentsRepository) LikesCommentsByPostRepository(id int) (map[int][]int, error) {
	commentsLikes := map[int][]int{}
	rows, err := repo.db.Query(likesCommentsByPost, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment_id, likes, dislikes int
		if err := rows.Scan(&comment_id, &likes, &dislikes); err != nil {
			return nil, nil
		}
		commentsLikes[comment_id] = append(commentsLikes[comment_id], likes, dislikes)
	}
	return commentsLikes, nil
}

func (repo *CommentsRepository) PostCommentsRepository(ctx context.Context, id int) (*model.PostCommentsDTO, error) {
	postId := &model.Post{}
	if err := repo.db.QueryRowContext(ctx, postByIdQuery, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate, &postId.Author); err != nil {
		return nil, err
	}
	return &model.PostCommentsDTO{Post: postId}, nil
}
