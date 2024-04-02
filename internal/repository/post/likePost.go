package post

import (
	"database/sql"
	"fmt"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createLikePostQuery   = "INSERT INTO posts_likes(user_id, post_id, status) VALUES($1, $2, $3)"
	deleteLikePostQuery   = "DELETE FROM posts_likes WHERE user_id = $1 AND post_id = $2"
	existLikePostQuery    = "SELECT * FROM posts_likes WHERE user_id = $1 AND post_id = $2"
	likesAndDislikesQuery = "SELECT SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM posts_likes WHERE post_id = $1 GROUP BY post_id"
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
	fmt.Println(like)
	if _, err := l.db.Exec(createLikePostQuery, like.UserId, like.PostId, like.LikeStatus); err != nil {
		return err
	}
	return nil
}

func (l *LikePostRepository) DeleteLikeByPostIdRepository(user_id, post_id int) error {
	if _, err := l.db.Exec(deleteLikePostQuery, user_id, post_id); err != nil {
		return err
	}
	return nil
}

func (l *LikePostRepository) GetLikePostRepository(userId, postId int) (*model.LikePost, error) {
	likedPost := &model.LikePost{}
	if err := l.db.QueryRow(existLikePostQuery, userId, postId).Scan(&likedPost.UserId, &likedPost.PostId, &likedPost.LikeStatus); err != nil {
		return nil, err
	}
	return likedPost, nil
}

func (l *LikePostRepository) GetLikesAndDislikesPostRepository(postId int) (int, int, error) {
	var likes, dislikes int
	if err := l.db.QueryRow(likesAndDislikesQuery, postId).Scan(&likes, &dislikes); err != nil {
		return -1, -1, err
	}
	return likes, dislikes, nil
}

func (l *LikePostRepository) GetUserLikedPostRepository(like *model.LikePost) error { return nil }
