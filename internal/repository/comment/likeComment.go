package comment

// const (
// 	addVote      = `INSERT INTO posts_votes( post_id, user_id, vote ) VALUES($1,$2,$3)`
// 	getLikePosts = `SELECT t1.post_id, t1.user_id, t1.category_name, t1.title, t1.discription, t1.create_at FROM posts t1 JOIN posts_votes t2 ON t1.id = t2.post_id WHERE t2.user_id = $1 AND t2.vote = 1`
// 	checkVote    = `SELECT vote FROM posts_vote WHERE post_id = $1 AND user_id = $2`
// 	deleteVote   = `DELETE  FROM posts_vote WHERE post_id = $1 AND user_id = $2`
// )

// type LikePostRepository struct {
// 	db *sql.DB
// }

// func NewLikePostRepository(db *sql.DB) *LikePostRepository {
// 	return &LikePostRepository{
// 		db: db,
// 	}
// }

// func (l *LikePostRepository) VotePostsRepository(post model.LikePost) error {
// 	_, err := l.db.Exec(addVote, post.PostId, post.UserId, post.IsLike)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (l *LikePostRepository) LikePostsRepository(userId int) ([]*model.Post, error) {
// 	rows, err := l.db.Query(getLikePosts, userId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	likePosts := make([]*model.Post, 0)
// 	for rows.Next() {
// 		post := new(model.Post)
// 		err := rows.Scan(&post.PostId, &post.UserId, &post.CategoryName, &post.Title, &post.Discription, &post.CreateDate)
// 		if err != nil {
// 			return nil, err
// 		}
// 		likePosts = append(likePosts, post)
// 	}
// 	return likePosts, nil
// }

// func (l *LikePostRepository) CheckVotePost(post model.LikePost) (string, error) {
// 	rows, err := l.db.Query(checkVote, post.PostId, post.UserId)
// 	if err != nil {
// 		return "", err
// 	}
// 	resVote := new(int)
// 	for rows.Next() {
// 		err := rows.Scan(
// 			&resVote)
// 		if err != nil {
// 			return "", err
// 		}
// 		if 1 == *resVote || -1 == *resVote {
// 			return "yes", nil
// 		}
// 	}
// 	return "no", nil
// }

// func (l *LikePostRepository) DeleteVotePost(post model.LikePost) error {
// 	if _, err := l.db.Exec(deleteVote, post.PostId, post.UserId); err != nil {
// 		return err
// 	}
// 	return nil
// }
