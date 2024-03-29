package user

import (
	"fmt"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
	"github.com/gofrs/uuid"
)

type ISessionRepository interface {
	CreateSession(*model.Sessinon) error
	DeleteSession(string) error
	UserIDBySession(*model.Sessinon) (int, error)
	SessionByID(int) (*model.Sessinon, error)
	SessionByUUID(string) (*model.Sessinon, error)
}

type ISessionService interface {
	CreateSessionService(id int) (*model.Sessinon, error)
	DeleteSessionService(uuid string) error
	UserIDService(session *model.Sessinon) (int, error)
	GetSessionByUUIDService(uuid string) (*model.Sessinon, error)
}

type SessinonService struct {
	iSessionRepository ISessionRepository
}

func NewSessionService(iSessionRepository ISessionRepository) *SessinonService {
	return &SessinonService{iSessionRepository: iSessionRepository}
}

func (ss *SessinonService) CreateSessionService(id int) (*model.Sessinon, error) {
	oldSession, _ := ss.iSessionRepository.SessionByID(id)
	if oldSession != nil {
		if err := ss.iSessionRepository.DeleteSession(oldSession.UUID); err != nil {
			return nil, err
		}
	}

	// work
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	session := &model.Sessinon{
		UUID:     string(uuid.String()),
		UserID:   id,
		ExpireAt: time.Now().Add(time.Minute * 30),
	}
	if err := ss.iSessionRepository.CreateSession(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (ss *SessinonService) DeleteSessionService(uuid string) error {
	if err := ss.iSessionRepository.DeleteSession(uuid); err != nil {
		return err
	}
	return nil
}

func (ss *SessinonService) UserIDService(session *model.Sessinon) (int, error) {
	user_id, err := ss.iSessionRepository.UserIDBySession(session)
	if err != nil {
		return -1, err
	}
	return user_id, nil
}

func (ss *SessinonService) GetSessionByUUIDService(uuid string) (*model.Sessinon, error) {
	session, err := ss.iSessionRepository.SessionByUUID(uuid)
	switch err {
	case nil:
		if session.ExpireAt.Before(time.Now()) {
			return nil, fmt.Errorf("Time Session Expired")
		}
		return session, nil
	default:
		return nil, err
	}
}
