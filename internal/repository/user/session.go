package user

import (
	"database/sql"
	"fmt"

	"gitea.com/lzhuk/forum/internal/model"
)

type SessinonRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessinonRepository {
	return &SessinonRepository{db: db}
}

const (
	createSessionQuery   = `INSERT INTO sessions (uuid, user_id, expire_at) VALUES ($1,$2,$3)`
	deleteSessionQuery   = `DELETE FROM sessions WHERE uuid = $1`
	userIDQueryBySession = `SELECT user_id FROM sessions WHERE uuid = $1`
	sessionQueryByID     = `SELECT * FROM sessions WHERE user_id = $1`
	sessionQueryByUUID   = `SELECT * FROM sessions WHERE uuid = $1`
)

func (s *SessinonRepository) CreateSession(session *model.Session) error {
	if _, err := s.db.Exec(createSessionQuery, session.UUID, session.UserID, session.ExpireAt); err != nil {
		return err
	}
	return nil
}

func (s *SessinonRepository) DeleteSession(uuid string) error {
	if _, err := s.db.Exec(deleteSessionQuery, uuid); err != nil {
		return err
	}
	return nil
}

func (s *SessinonRepository) SessionByID(userID int) (*model.Session, error) {
	var session model.Session
	if err := s.db.QueryRow(sessionQueryByID, userID).Scan(&session.UUID, &session.UserID, &session.ExpireAt); err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessinonRepository) SessionByUUID(uuid string) (*model.Session, error) {
	session := &model.Session{}
	fmt.Println("check", uuid)
	if err := s.db.QueryRow(sessionQueryByUUID, uuid).Scan(&session.UUID, &session.UserID, &session.ExpireAt); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessinonRepository) UserIDBySession(session *model.Session) (int, error) {
	var userId int
	if err := s.db.QueryRow(userIDQueryBySession, session.UUID).Scan(&userId); err != nil {
		return -1, err
	}
	return userId, nil
}
