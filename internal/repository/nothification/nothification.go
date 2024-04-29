package nothification

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

type NothificationRepository struct {
	DB *sql.DB
}

func InitNothificationRepository(db *sql.DB) *NothificationRepository {
	return &NothificationRepository{
		DB: db,
	}
}

const (
	nothCreateQuery     = `INSERT INTO nothifications (user_id, post_id, type, created_user_id, message) VALUES($1, $2, $3, $4, $5)`
	readedQuery         = `UPDATE nothifications SET is_read = TRUE`
	nothificationsQuery = `SELECT * FROM nothifications WHERE user_id = $1 AND is_read = FALSE`
	noth                = `SELECT EXISTS (SELECT 1 FROM nothifications WHERE user_id = $1 AND is_read = FALSE) AS check`
)

func (n *NothificationRepository) Create() error {
	if _, err := n.DB.Exec(nothCreateQuery); err != nil {
		return err
	}
	return nil
}

func (n *NothificationRepository) NothificationIsRead() (bool, error) { // CHECK
	var isRead bool
	if err := n.DB.QueryRow(noth).Scan(&isRead); err != nil {
		return false, err
	}
	return isRead, nil
}

func (n *NothificationRepository) Read(user_id int) ([]model.Nothification, error) {
	userNothifications := []model.Nothification{}
	rows, err := n.DB.Query(nothificationsQuery, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		nothification := &model.Nothification{}
		if err := rows.Scan(&nothification.ID, &nothification.UserID, &nothification.PostID, &nothification.Type, &nothification.CreatedUserID, &nothification.Message, &nothification.IsRead, &nothification.CreatedAt); err != nil {
			return nil, err
		}
		userNothifications = append(userNothifications, *nothification)
	}
	return userNothifications, nil
}

func (n *NothificationRepository) Update() error {
	if _, err := n.DB.Exec(readedQuery); err != nil {
		return err
	}
	return nil
}
