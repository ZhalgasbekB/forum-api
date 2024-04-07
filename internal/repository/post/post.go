package post

import (
	"context"
	"database/sql"
	"sync"

	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createPostQuery       = `INSERT INTO posts(user_id, category_name, title, description, create_at) VALUES($1,$2,$3,$4,$5)`
	updatePostUserIdQuery = `UPDATE posts SET description = $1, title = $2 WHERE id = $3 AND user_id = $4;`
	deletePostUserIdQuery = `DELETE FROM posts WHERE id = $1 AND user_id = $2`

	postsQuery          = `SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name FROM posts ps JOIN users u ON ps.user_id = u.id`
	postByIdQuery       = `SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name FROM posts ps JOIN users u ON ps.user_id = u.id WHERE ps.id = $1`
	postsByUserIdQuery = `SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name, SUM(CASE WHEN pl.status = TRUE THEN 1 ELSE 0 END) AS likes, SUM(CASE WHEN pl.status = FALSE THEN 1 ELSE 0 END) AS dislikes FROM posts ps JOIN users u ON ps.user_id = u.id LEFT JOIN posts_likes pl ON ps.id = pl.post_id WHERE ps.user_id = $1 GROUP BY ps.id, u.id ORDER BY ps.create_at DESC;`
	postCommentsQuery   = `SELECT c.id, c.post_id, c.user_id, c.description, c.created_at, c.updated_at FROM posts p JOIN comments c ON c.post_id = p.id WHERE p.id = $1`
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
		if err == sql.ErrNoRows {
			return errors.ErrNotFoundData
		}
		return err
	}
	return nil
}

func (p PostsRepository) PostsRepository(ctx context.Context) ([]*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	rows, err := p.db.QueryContext(ctx, postsQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFoundData
		}
		return nil, err
	}
	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := &model.Post{}
		err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Description, &post.CreateDate, &post.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p PostsRepository) PostByIdRepository(ctx context.Context, id int) (*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	postId := &model.Post{}
	if err := p.db.QueryRowContext(ctx, postByIdQuery, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate, &postId.Author); err != nil {
		return nil, err
	}
	return postId, nil
}

func (p PostsRepository) PostByUserIdRepository(ctx context.Context, userId int) ([]*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()

	rows, err := p.db.QueryContext(ctx, postsByUserIdQuery, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFoundData
		}
		return nil, err
	}
	userPosts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)
		if err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Description, &post.CreateDate, &post.Author, &post.Like, &post.Dislike); err != nil {
			return nil, err
		}
		userPosts = append(userPosts, post)
	}
	return userPosts, nil
}

func (p *PostsRepository) UpdatePostByUserIdRepository(ctx context.Context, post model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()
	if _, err := p.db.ExecContext(ctx, updatePostUserIdQuery, post.Description, post.Title, post.PostId, post.UserId); err != nil {
		return err
	}
	return nil
}

func (p *PostsRepository) DeletePostByUserIdRepository(ctx context.Context, deleteModel *model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()
	if _, err := p.db.ExecContext(ctx, deletePostUserIdQuery, deleteModel.PostId, deleteModel.UserId); err != nil {
		return err
	}
	return nil
}

func (p *PostsRepository) PostCommentsRepository(ctx context.Context, id int) (*model.PostCommentsDTO, error) {
	p.m.Lock()
	defer p.m.Unlock()
	postComments := &model.PostCommentsDTO{}
	postId := &model.Post{}
	if err := p.db.QueryRowContext(ctx, postByIdQuery, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate, &postId.Author); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFoundData
		}
		return nil, err
	}
	postComments.Post = postId
	return postComments, nil
}
