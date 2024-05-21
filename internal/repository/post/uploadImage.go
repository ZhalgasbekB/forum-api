package post

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	addImage    = "INSERT INTO images(path, post_id) VALUES($1, $2)"
	getImage    = "SELECT path FROM images WHERE post_id = $1 AND is_active = 1"
	checkImage  = "UPDATE images SET path = $1 WHERE post_id = $2 AND is_active = 1"
	deleteImage = "UPDATE images SET is_active = 0 WHERE post_id = $1 AND is_active = 1 "
)

type UploadImagePostRepository struct {
	db *sql.DB
}

func NewUploadImagePostRepository(db *sql.DB) *UploadImagePostRepository {
	return &UploadImagePostRepository{
		db: db,
	}
}

func (u *UploadImagePostRepository) AddImagePostRepository(image *model.UploadPostImage) error {
	if _, err := u.db.Exec(addImage, image.Path, image.PostId); err != nil {
		return err
	}
	return nil
}

func (u *UploadImagePostRepository) GetImagePostRepository(postId int) (string, error) {
	var path string
	if err := u.db.QueryRow(getImage, postId).Scan(&path); err != nil {
		return "", err
	}
	return path, nil
}

func (u *UploadImagePostRepository) UpdateImagePostRepository(image *model.UploadPostImage) error {
	if _, err := u.db.Exec(checkImage, image.Path, image.PostId); err != nil {
		return err
	}
	return nil
}

func (u *UploadImagePostRepository) DeleteImagePostRepository(postId int) error {
	if _, err := u.db.Exec(deleteImage, postId); err != nil {
		return err
	}
	return nil
}
