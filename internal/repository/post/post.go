package post

import (
	"context"
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	createPostQuery = `
	INSERT INTO posts(
		user_id,
		category_name,
		title,
		discription,
		create_at
	)
	VALUES($1,$2,$3,$4,$5)
	`
	getAllPost = `
	SELECT *
	FROM posts`

	getIdPost = `
	SELECT *
	FROM posts
	WHERE id = $1
	`
	getUserPost = `
	SELECT *
	FROM posts
	WHERE user_id = $1
	`
	updateUserPost = `
	UPDATE posts
	SET
    	discription = $1,
    	create_at = $2
	WHERE
    id = $3 AND user_id = $4;
	`
	deleteUserPost = `
	DELETE FROM posts
	WHERE id = $1 AND user_id = $2
	`
	addVote = `
	INSERT INTO posts_votes(
		post_id,
		user_id,
		vote
	)
	VALUES($1,$2,$3)
	`
	checkTitle = `
	SELECT title
	FROM posts
	WHERE title = $1
	`
	checkDiscription = `
	SELECT discription
	FROM posts
	WHERE discription = $1
	`
	getLikePosts = `
	SELECT t1.post_id, t1.user_id, t1.category_name, t1.title, t1.discription, t1.create_at
	FROM posts t1
	JOIN posts_votes t2
	ON t1.id = t2.post_id
	WHERE
	t2.user_id = $1 AND t2.vote = 1
	`
	checkVote = `
	SELECT vote
	FROM posts_vote
	WHERE post_id = $1 AND user_id = $2
	`
	deleteVote = `
	DELETE 
	FROM posts_vote
	WHERE post_id = $1 AND user_id = $2
	`
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepo(db *sql.DB) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}

// Метод создания новой темы(поста) в БД
func (p PostsRepository) CreatePostRepo(ctx context.Context, post model.CreatePost) error {
	_, err := p.db.Exec(
		createPostQuery,
		post.UserId,
		post.CategoryName,
		post.Title,
		post.Discription,
		post.CreateDate)
	if err != nil {
		return err
	}
	return nil
}

// Метод получения всех тем(постов) из БД
func (p PostsRepository) GetAllPostRepo() ([]*model.Post, error) {
	rows, err := p.db.Query(getAllPost)
	if err != nil {
		return nil, err
	}
	// Срез для записи всех тем(постов) из БД
	allPosts := make([]*model.Post, 0)

	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(
			&post.PostId,
			&post.UserId,
			&post.CategoryName,
			&post.Title,
			&post.Discription,
			&post.CreateDate,
		)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}

// Метод получения конкретной темы (поста) по id
func (p PostsRepository) GetIdPostRepo(id int) (*model.Post, error) {
	postId := &model.Post{}
	err := p.db.QueryRow(getIdPost, id).Scan(
		&postId.PostId,
		&postId.UserId,
		&postId.CategoryName,
		&postId.Title,
		&postId.Discription,
		&postId.CreateDate,
	)
	if err != nil {
		return nil, err
	}
	return postId, nil
}

// Метод получения всех созданных тем (постов) конкретным пользователем
func (p PostsRepository) GetUserPostRepo(userId int) ([]*model.Post, error) {
	rows, err := p.db.Query(getUserPost, userId)
	if err != nil {
		return nil, err
	}
	// Срез для записи всех тем(постов) из БД
	userPosts := make([]*model.Post, 0)

	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(
			&post.PostId,
			&post.UserId,
			&post.CategoryName,
			&post.Title,
			&post.Discription,
			&post.CreateDate,
		)
		if err != nil {
			return nil, err
		}
		userPosts = append(userPosts, post)
	}
	return userPosts, nil
}

// Метод изменения поста созданного пользователем
func (p *PostsRepository) UpdateUserPostRepo(post model.UpdatePost) error {
	_, err := p.db.Exec(
		updateUserPost,
		post.Discription,
		post.CreateDate,
		post.PostId,
		post.UserId)
	if err != nil {
		return err
	}
	return nil
}

// Метод для удаления поста пользователем
func (p *PostsRepository) DeleteUserPostRepo(deleteModel *model.DeletePost) error {
	_, err := p.db.Exec(
		deleteUserPost,
		deleteModel.PostId,
		deleteModel.UserId)
	if err != nil {
		return err
	}
	return nil
}

// Метод для постановки голоса
func (p *PostsRepository) VotePostsRepo(post model.Vote) error {
	_, err := p.db.Exec(
		addVote,
		post.PostId,
		post.UserId,
		post.Vote)
	if err != nil {
		return err
	}
	return nil
}

// Метод для проверки наличия уже темы в БД
func (p *PostsRepository) CheckTitlePost(title string) (string, error) {
	rows, err := p.db.Query(
		checkTitle,
		title)
	if err != nil {
		return "", err
	}
	resTitle := new(string)
	for rows.Next() {
		err := rows.Scan(
			&resTitle)
		if err != nil {
			return "", err
		}
		if title == *resTitle {
			return "yes", nil
		}
	}
	return "no", nil
}

// Метод получения всех понравившихся тем (постов) пользователем
func (p *PostsRepository) LikePostsRepo(userId int) ([]*model.Post, error) {
	rows, err := p.db.Query(getLikePosts, userId)
	if err != nil {
		return nil, err
	}
	// Срез для записи всех тем(постов) из БД
	likePosts := make([]*model.Post, 0)

	for rows.Next() {
		post := new(model.Post)
		err := rows.Scan(
			&post.PostId,
			&post.UserId,
			&post.CategoryName,
			&post.Title,
			&post.Discription,
			&post.CreateDate,
		)
		if err != nil {
			return nil, err
		}
		likePosts = append(likePosts, post)
	}
	return likePosts, nil
}

// Метод для проверки наличия голоса в теме
func (p *PostsRepository) CheckVotePost(post model.Vote) (string, error) {
	rows, err := p.db.Query(
		checkVote,
		post.PostId,
		post.UserId)
	if err != nil {
		return "", err
	}
	resVote := new(int)
	for rows.Next() {
		err := rows.Scan(
			&resVote)
		if err != nil {
			return "", err
		}
		if 1 == *resVote || -1 == *resVote {
			return "yes", nil
		}
	}
	return "no", nil
}

// Метод для удаления голоса с темы (поста)
func (p *PostsRepository) DeleteVotePost(post model.Vote) error {
	_, err := p.db.Exec(
		deleteVote,
		post.PostId,
		post.UserId)
	if err != nil {
		return err
	}
	return nil
}
