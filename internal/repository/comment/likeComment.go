package comment

import (
	"database/sql"
	"math"

	"gitea.com/lzhuk/forum/internal/model"
)

type LikeCommentRepostory struct {
	db *sql.DB
}

const (
	createLikeCommentQuery  = `INSERT INTO comments_likes(user_id, comment_id, status) VALUES($1, $2, $3)`
	deleteLikeCommentQuery  = `DELETE FROM comments_likes WHERE user_id = $1 AND comment_id = $2`
	existLikeCommentQuery   = `SELECT * FROM comments_likes WHERE user_id = $1 AND comment_id = $2`
	likesAndDislikesQuery   = `SELECT SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM comments_likes WHERE comment_id = $1 GROUP BY comment_id`
	likeANDDislikesAllQuery = `SELECT comment_id, SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM comments_likes GROUP BY comment_id`
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

func (l *LikeCommentRepostory) GetLikeCommentRepository(userId, postId int) (*model.LikeComment, error) {
	likedPost := &model.LikeComment{}
	if err := l.db.QueryRow(existLikeCommentQuery, userId, postId).Scan(&likedPost.UserId, &likedPost.CommentId, &likedPost.LikeStatus); err != nil {
		return nil, err
	}
	return likedPost, nil
}

 

func (l *LikeCommentRepostory) GetUserLikedCommentRepository(like *model.LikeComment) error {
	return nil
}

func (l *LikeCommentRepostory) GetLikesAndDislikesCommentAllRepository() (map[int][]int, error) {
	commentsLikes := map[int][]int{}
	rows, err := l.db.Query(likesAndDislikesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment_id, likes, dislikes int
		if err := rows.Scan(&comment_id, &likes, &dislikes); err != nil {
			return nil, nil
		}
		commentsLikes[comment_id] = append(commentsLikes[comment_id], likes, int(math.Abs(float64(dislikes))))
	}
	return commentsLikes, nil
}
