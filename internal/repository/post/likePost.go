package post

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	updateLikePostQuery = "UPDATE posts_likes SET status = $1, like_code = $2 like_code  WHERE user_id = $4 AND post_id = $5"
	createLikePostQuery = "INSERT INTO posts_likes(user_id, post_id, status, like_code) VALUES($1, $2, $3, $4)"
	existLikePostQuery  = "SELECT * FROM posts_likes WHERE user_id = $4 AND post_id = $5"
	likesAndDislikes    = "SELECT SUM(CASE WHEN like_code = 1 THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN like_code = -1 THEN 1 ELSE 0 END) AS dislikes FROM posts_likes WHERE post_id = $1 GROUP BY post_id"
)

type LikePostRepository struct {
	db *sql.DB
}

func NewLikePostRepository(db *sql.DB) *LikePostRepository {
	return &LikePostRepository{
		db: db,
	}
}

func (l *LikePostRepository) CreateLikePostRepository(like *model.LikePost) error {
	if _, err := l.db.Exec(createLikePostQuery, like.UserId, like.PostId, like.LikeStatus, like.LikeCode); err != nil {
		return err
	}
	return nil
}

func (l *LikePostRepository) UpdateLikePostRepository(like *model.LikePost) error {
	if _, err := l.db.Exec(updateLikePostQuery, like.LikeStatus, like.LikeCode, like.UserId, like.PostId); err != nil {
		return err
	}
	return nil
}

func (l *LikePostRepository) GetLikePostRepository(userId, postId int) (*model.LikePost, error) {
	likedPost := &model.LikePost{}
	if err := l.db.QueryRow(existLikePostQuery, userId, postId).Scan(&likedPost.UserId, &likedPost.PostId, &likedPost.LikeStatus, &likedPost.LikeCode); err != nil {
		return nil, err
	}
	return likedPost, nil
}

func (l *LikePostRepository) GetLikesAndDislikesPostRepository(postId int) (int, int, error) {
	var likes, dislikes int
	if err := l.db.QueryRow(likesAndDislikes, postId).Scan(&likes, &dislikes); err != nil {
		return -1, -1, err
	}
	return likes, dislikes, nil
}
func (l *LikePostRepository) GetUserLikedPostRepository(like *model.LikePost) error { return nil }
