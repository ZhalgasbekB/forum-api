package post

import "database/sql"

const (
	upLikeQuery         = ""
	downLikeQuery       = ""
	createLikePostQuery = ""
)

type LikePostRepository struct {
	db *sql.DB
}

func NewLikePostRepository(db *sql.DB) *LikePostRepository {
	return &LikePostRepository{
		db: db,
	}
}

func CreateLikePostRepository() error { return nil }
func UpdateLikePostRepository() error { return nil }
func DeleteLikePostRepository() error { return nil }
func GetLikePostRepository() error    { return nil }
