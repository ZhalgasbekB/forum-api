package post

import (
	"context"
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createPostQuery = `INSERT INTO posts(user_id, category_name, title, discription, create_at) VALUES($1,$2,$3,$4,$5)`
	getAllPost      = `SELECT * FROM posts`

	getIdPost      = `SELECT * FROM posts WHERE id = $1`
	getUserPost    = `SELECT * FROM posts WHERE user_id = $1`
	updateUserPost = `UPDATE posts SET discription = $1, create_at = $2 WHERE id = $3 AND user_id = $4;`
	deleteUserPost = `DELETE FROM posts WHERE id = $1 AND user_id = $2`
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepo(db *sql.DB) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}

func (p PostsRepository) CreatePostRepository(ctx context.Context, post model.CreatePost) error {
	if _, err := p.db.Exec(createPostQuery, post.UserId, post.CategoryName, post.Title, post.Discription, post.CreateDate); err != nil {
		return err
	}
	return nil
}

func (p PostsRepository) AllPostRepository() ([]*model.Post, error) {
	rows, err := p.db.Query(getAllPost)
	if err != nil {
		return nil, err
	}
	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Discription, &post.CreateDate)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p PostsRepository) IdPostRepository(id int) (*model.Post, error) {
	postId := &model.Post{}
	if err := p.db.QueryRow(getIdPost, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Discription, &postId.CreateDate); err != nil {
		return nil, err
	}
	return postId, nil
}

func (p PostsRepository) UserPostRepository(userId int) ([]*model.Post, error) {
	rows, err := p.db.Query(getUserPost, userId)
	if err != nil {
		return nil, err
	}
	userPosts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)
		if err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Discription, &post.CreateDate); err != nil {
			return nil, err
		}
		userPosts = append(userPosts, post)
	}
	return userPosts, nil
}

func (p *PostsRepository) UpdateUserPostRepository(post model.UpdatePost) error {
	if _, err := p.db.Exec(updateUserPost, post.Discription, post.CreateDate, post.PostId, post.UserId); err != nil {
		return err
	}
	return nil
}

func (p *PostsRepository) DeleteUserPostRepository(deleteModel *model.DeletePost) error {
	if _, err := p.db.Exec(deleteUserPost, deleteModel.PostId, deleteModel.UserId); err != nil {
		return err
	}
	return nil
}
