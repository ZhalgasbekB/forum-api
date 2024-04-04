package post

import (
	"database/sql"
	"fmt"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createLikePostQuery = "INSERT INTO posts_likes(user_id, post_id, status) VALUES($1, $2, $3)"
	deleteLikePostQuery = "DELETE FROM posts_likes WHERE user_id = $1 AND post_id = $2"

	likePostQuery            = "SELECT * FROM posts_likes WHERE user_id = $1 AND post_id = $2"
	likesAndDislikesQuery    = "SELECT SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM posts_likes WHERE post_id = $1 GROUP BY post_id"
	likesAndDislikesAllQuery = "SELECT post_id, SUM(CASE WHEN status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN status = false THEN 1 ELSE 0 END) AS dislikes FROM posts_likes GROUP BY post_id"
	likedPostAndHisLikes     = "SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name, SUM(CASE WHEN p.status = true THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN p.status = false THEN 1 ELSE 0 END) AS dislikes FROM posts_likes p JOIN posts ps ON ps.id = p.post_id JOIN users u ON ps.user_id = u.id WHERE ps.user_id = $1 GROUP BY ps.id"
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
	if err := l.db.QueryRow(likePostQuery, userId, postId).Scan(&likedPost.UserId, &likedPost.PostId, &likedPost.LikeStatus); err != nil {
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

func (l *LikePostRepository) GetUserLikedPostRepository(user_id int) ([]model.Post, error) {
	likedPosts := []model.Post{}

	rows, err := l.db.Query(likedPostAndHisLikes, user_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := model.Post{}
		if err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Description, &post.CreateDate, &post.Author, &post.Like, &post.Dislike); err != nil {
			return nil, err
		}
		likedPosts = append(likedPosts, post)
	}
	return likedPosts, nil
}

func (l *LikePostRepository) GetLikeAndDislikeAllPostRepository() (map[int][]int, error) {
	postLikes := map[int][]int{}
	rows, err := l.db.Query(likesAndDislikesAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post_id, likes, dislikes int
		if err := rows.Scan(&post_id, &likes, &dislikes); err != nil {
			return nil, nil
		}
		postLikes[post_id] = append(postLikes[post_id], likes, dislikes)
	}
	return postLikes, nil
}
