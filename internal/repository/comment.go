package repository

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

const (
	updateCommQuery  = `UPDATE comments SET description = $1, updated_at = $2 WHERE id = $3`
	deleteCommQuery  = `DELETE FROM comments WHERE id = $1 AND user_id = $2`
	commentByIDQuery = `SELECT * FROM comments WHERE id = $1`
	commentsQuery    = `SELECT * FROM comments`
	createCommQuery  = `INSERT INTO comments (post_id, user_id, description, created_at, updated_at) VALUES ($1,$2,$3,$4,$5)`
)

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepo(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}

func (repo *CommentsRepository) CreateComment(comm *model.Comment) error {
	if _, err := repo.db.Exec(createCommQuery, comm.Post, comm.User, comm.Description, comm.CreatedDate, comm.UpdatedDate); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) UpdateComment(comm *model.Comment) error {
	if _, err := repo.db.Exec(updateCommQuery, comm.Description, comm.UpdatedDate, comm.ID); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) DeleteComment(comm *model.Comment) error {
	if _, err := repo.db.Exec(deleteCommQuery, comm.ID, comm.User); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) CommentByID(id int) (*model.Comment, error) {
	var comm model.Comment
	if err := repo.db.QueryRow(commentByIDQuery, id).Scan(&comm.ID, &comm.Post, &comm.User, &comm.Description, &comm.CreatedDate, &comm.UpdatedDate); err != nil {
		return nil, err
	}
	return &comm, nil
}

func (repo *CommentsRepository) Comments() ([]model.Comment, error) {
	var comments []model.Comment
	rows, err := repo.db.Query(commentsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comm model.Comment
		if err := rows.Scan(&comm.ID, &comm.Post, &comm.User, &comm.Description, &comm.CreatedDate, &comm.UpdatedDate); err != nil {
			return nil, err
		}
		comments = append(comments, comm)
	}
	return comments, nil
}
