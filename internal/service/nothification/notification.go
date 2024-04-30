package nothification

import (
	"gitea.com/lzhuk/forum/internal/model"
)

type INotificationRepository interface {
	Create(*model.Notification) error
	Read(int) ([]model.Notification, error)
	Update(int, int) error
	NotificationIsRead(int) (bool, error)

	DuplicateNotification(*model.Notification) (*model.Notification, error)
	DeleteNotification(*model.Notification) error
}

type INotificationService interface {
	CreateService(*model.Notification, bool) error
	NotificationsService(int) ([]model.Notification, error)
	UpdateService(int, int) error
	NotificationIsReadService(int) (bool, error)
	CreateCommentService(*model.Notification) error
}

type NotificationService struct {
	notificationRepository INotificationRepository
}

func InitNotificationService(notificationRepository INotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
	}
}

func (ns *NotificationService) CreateService(n *model.Notification, isLiked bool) error {
	nOld, err := ns.notificationRepository.DuplicateNotification(n)
	if err != nil {
		return err
	}

	if nOld != nil && (!isLiked || (nOld.Type != n.Type)) {
		if err := ns.notificationRepository.DeleteNotification(nOld); err != nil {
			return err
		}

		if !isLiked {
			return nil
		}
	}

	return ns.notificationRepository.Create(n)
}

func (ns *NotificationService) CreateCommentService(n *model.Notification) error {
	return ns.notificationRepository.Create(n)
}

func (ns *NotificationService) NotificationsService(u_id int) ([]model.Notification, error) {
	return ns.notificationRepository.Read(u_id)
}

func (ns *NotificationService) UpdateService(u_id, id int) error {
	return ns.notificationRepository.Update(u_id, id)
}

func (ns *NotificationService) NotificationIsReadService(u_id int) (bool, error) {
	return ns.notificationRepository.NotificationIsRead(u_id)
}
