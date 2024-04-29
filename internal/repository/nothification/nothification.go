package nothification

import "database/sql"

const query = ""

type NothificationRepository struct {
	DB *sql.DB
}

func InitNothificationRepository(db *sql.DB) *NothificationRepository {
	return &NothificationRepository{
		DB: db,
	}
}

const (
	nothCreateQuery = `INSERT INTO nothifications (user_id, post_id, type, created_user_id, message) VALUES($1, $2, $3, $4, $5)`
	readedQuery     = `UPDATE nothifications SET is_read = TRUE` /// ??? ADDS
)

func (n *NothificationRepository) Create() error {
	if _, err := n.DB.Exec(query); err != nil {
		return err
	}
	return nil
}

func (n *NothificationRepository) Read() error {
	return nil
}

func (n *NothificationRepository) Update() error {
	if _, err := n.DB.Exec(query); err != nil {
		return err
	}
	return nil
}

func (n *NothificationRepository) Delete() error {
	if _, err := n.DB.Exec(query); err != nil {
		return err
	}
	return nil
}
