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
