package post

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createPostQuery = `INSERT INTO posts(user_id, category_name, title, description, create_at) VALUES($1,$2,$3,$4,$5)`
	getAllPost      = `SELECT * FROM posts`

	getIdPost      = `SELECT * FROM posts WHERE id = $1`
	getUserPost    = `SELECT * FROM posts WHERE user_id = $1`
	updateUserPost = `UPDATE posts SET description = $1, title = $2 WHERE id = $3 AND user_id = $4;`
	deleteUserPost = `DELETE FROM posts WHERE id = $1 AND user_id = $2`

	postCommentsQuery = `SELECT c.id, c.post_id, c.user_id, c.description, c.created_at, c.updated_at FROM posts p JOIN comments c ON c.post_id = p.id WHERE p.id = $1`
)

type PostsRepository struct {
	db *sql.DB
	m  sync.RWMutex
}

func NewPostsRepo(db *sql.DB) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}

func (p PostsRepository) CreatePostRepository(ctx context.Context, post model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()

	if _, err := p.db.ExecContext(ctx, createPostQuery, post.UserId, post.CategoryName, post.Title, post.Description, post.CreateDate); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("User Successfully CREATE POST")
	return nil
}

func (p PostsRepository) AllPostRepository(ctx context.Context) ([]*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	rows, err := p.db.QueryContext(ctx, getAllPost)
	if err != nil {
		return nil, err
	}
	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Description, &post.CreateDate)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	fmt.Println("User Successfully  POSTS")
	return posts, nil
}

func (p PostsRepository) IdPostRepository(ctx context.Context, id int) (*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	postId := &model.Post{}
	if err := p.db.QueryRowContext(ctx, getIdPost, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate); err != nil {
		return nil, err
	}
	fmt.Println("User Successfully  POST")
	return postId, nil
}

func (p PostsRepository) UserPostRepository(ctx context.Context, userId int) ([]*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	rows, err := p.db.QueryContext(ctx, getUserPost, userId)
	if err != nil {
		return nil, err
	}
	userPosts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)
		if err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Description, &post.CreateDate); err != nil {
			return nil, err
		}
		userPosts = append(userPosts, post)
	}
	fmt.Println("User Successfully USER")
	return userPosts, nil
}

func (p *PostsRepository) UpdateUserPostRepository(ctx context.Context, post model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()
	if _, err := p.db.ExecContext(ctx, updateUserPost, post.Description, post.Title, post.PostId, post.UserId); err != nil {
		return err
	}
	fmt.Println("User Successfully USERPOSTID")
	return nil
}

func (p *PostsRepository) DeleteUserPostRepository(ctx context.Context, deleteModel *model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()
	if _, err := p.db.ExecContext(ctx, deleteUserPost, deleteModel.PostId, deleteModel.UserId); err != nil {
		return err
	}
	fmt.Println("User Successfully DETELE")
	return nil
}

func (p *PostsRepository) CommentsPostRepository(ctx context.Context, id int) (*model.PostCommentsDTO, error) {
	postComments := &model.PostCommentsDTO{}

	postId := &model.Post{}
	if err := p.db.QueryRowContext(ctx, getIdPost, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate); err != nil {
		return nil, err
	}
	postComments.Post = postId

	rows, err := p.db.QueryContext(ctx, postCommentsQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.Post, &comment.User, &comment.Description, &comment.CreatedDate, &comment.UpdatedDate); err != nil {
			return nil, err
		}
		postComments.Comments = append(postComments.Comments, &comment)
	}

	return postComments, nil
}
