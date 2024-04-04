package post

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createPostQuery       = `INSERT INTO posts(user_id, category_name, title, description, create_at) VALUES($1,$2,$3,$4,$5)`
	postsQuery            = `SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name FROM posts ps JOIN users u ON ps.user_id = u.id`
	postByIdQuery         = `SELECT ps.id, ps.user_id, ps.category_name, ps.title, ps.description, ps.create_at, u.name FROM posts ps JOIN users u ON ps.user_id = u.id WHERE ps.id = $1` // CHECK
	postsByUserIdQuery    = `SELECT * FROM posts WHERE user_id = $1`
	updatePostUserIdQuery = `UPDATE posts SET description = $1, title = $2 WHERE id = $3 AND user_id = $4;`
	deletePostUserIdQuery = `DELETE FROM posts WHERE id = $1 AND user_id = $2`
	postCommentsQuery     = `SELECT c.id, c.post_id, c.user_id, c.description, c.created_at, c.updated_at FROM posts p JOIN comments c ON c.post_id = p.id WHERE p.id = $1`
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

func (p PostsRepository) PostsRepository(ctx context.Context) ([]*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	rows, err := p.db.QueryContext(ctx, postsQuery)
	if err != nil {
		return nil, err
	}
	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Description, &post.CreateDate, &post.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	fmt.Println("User Successfully Returned Posts")
	return posts, nil
}

func (p PostsRepository) PostByIdRepository(ctx context.Context, id int) (*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock() 
	postId := &model.Post{}
	if err := p.db.QueryRowContext(ctx, postByIdQuery, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate, &postId.Author); err != nil {
		return nil, err
	}
	fmt.Println("User Successfully Return Post")
	return postId, nil
}

func (p PostsRepository) PostByUserIdRepository(ctx context.Context, userId int) ([]*model.Post, error) {
	p.m.Lock()
	defer p.m.Unlock()
	rows, err := p.db.QueryContext(ctx, postsByUserIdQuery, userId)
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
	fmt.Println("User Successfully User Posts Return")
	return userPosts, nil
}

func (p *PostsRepository) UpdatePostByUserIdRepository(ctx context.Context, post model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()
	if _, err := p.db.ExecContext(ctx, updatePostUserIdQuery, post.Description, post.Title, post.PostId, post.UserId); err != nil {
		return err
	}
	fmt.Println("User Successfully User Post Return")
	return nil
}

func (p *PostsRepository) DeletePostByUserIdRepository(ctx context.Context, deleteModel *model.Post) error {
	p.m.Lock()
	defer p.m.Unlock()
	if _, err := p.db.ExecContext(ctx, deletePostUserIdQuery, deleteModel.PostId, deleteModel.UserId); err != nil {
		return err
	}
	fmt.Println("User Successfully Delete Post")
	return nil
}

func (p *PostsRepository) PostCommentsRepository(ctx context.Context, id int) (*model.PostCommentsDTO, error) {
	postComments := &model.PostCommentsDTO{}

	postId := &model.Post{}
	if err := p.db.QueryRowContext(ctx, postByIdQuery, id).Scan(&postId.PostId, &postId.UserId, &postId.CategoryName, &postId.Title, &postId.Description, &postId.CreateDate, &postId.Author); err != nil {
		fmt.Println(err)
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
	fmt.Println("User Successfully Post Comments")
	return postComments, nil
}
