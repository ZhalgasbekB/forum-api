package comment

import (
	"database/sql"
	"fmt"

	"gitea.com/lzhuk/forum/internal/model"
)

type LikeCommentRepostory struct {
	db *sql.DB
}

const (
	createLikeCommentQuery = `INSERT INTO comments_likes(user_id, comment_id, status) VALUES($1, $2, $3)`
	deleteLikeCommentQuery = `DELETE FROM comments_likes WHERE user_id = $1 AND comment_id = $2`
	checkCommentQuery      = `SELECT * FROM comments_likes WHERE user_id = $1 AND comment_id = $2`
	likeAllQuery           = `SELECT comment_id, us.name, COUNT(CASE WHEN status = true THEN 1 END) AS likes, COUNT(CASE WHEN status = false THEN 1 END) AS dislikes FROM comments_likes c JOIN users us ON us.id = c.user_id GROUP BY c.comment_id;`
)

func NewLikeCommentRepository(db *sql.DB) *LikeCommentRepostory {
	return &LikeCommentRepostory{
		db: db,
	}
}

func (l *LikeCommentRepostory) CreateLikeCommentRepository(like *model.LikeComment) error {
	if _, err := l.db.Exec(createLikeCommentQuery, like.UserId, like.CommentId, like.LikeStatus); err != nil {
		return err
	}
	return nil
}

func (l *LikeCommentRepostory) DeleteLikeByCommentIdRepository(user_id, post_id int) error {
	if _, err := l.db.Exec(deleteLikeCommentQuery, user_id, post_id); err != nil {
		return err
	}
	return nil
}

func (l *LikeCommentRepostory) LikeCommentRepository(userId, postId int) (*model.LikeComment, error) {
	likedPost := &model.LikeComment{}
	if err := l.db.QueryRow(checkCommentQuery, userId, postId).Scan(&likedPost.UserId, &likedPost.CommentId, &likedPost.LikeStatus); err != nil {
		return nil, err
	}
	return likedPost, nil
}

func (l *LikeCommentRepostory) LikesAndDislikesCommentAllRepository() (map[int][]int, map[int]string, error) {
	commentsLikes := map[int][]int{}
	commentsNames := map[int]string{}

	rows, err := l.db.Query(likeAllQuery)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment_id, likes, dislikes int
		var name string

		if err := rows.Scan(&comment_id, &name, &likes, &dislikes); err != nil {
			return nil, nil, nil
		}
		commentsNames[comment_id] = name
		commentsLikes[comment_id] = append(commentsLikes[comment_id], likes, dislikes)
	}
	return commentsLikes, commentsNames, nil
}

func (l *LikeCommentRepostory) PostCommentsRepository(post_id int) ([]model.Comment, error) {
	likeComments := []model.Comment{}
	rows, err := l.db.Query("", post_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("s")
	defer rows.Close()
	for rows.Next() {
		comment := model.Comment{}
		if err := rows.Scan(&comment.ID, &comment.User, &comment.Post, &comment.Description, &comment.CreatedDate, &comment.UpdatedDate, &comment.Name, &comment.Like, &comment.Dislike); err != nil {
			return nil, err
		}

		likeComments = append(likeComments, comment)
	}
	return likeComments, nil
}
