package nothification

import "gitea.com/lzhuk/forum/internal/model"

type INotificationRepository interface {
	Create(*model.Notification) error
	Read(int) ([]model.Notification, error)
	Update(int) error
	NotificationIsRead(int) (bool, error)
}

type INotificationService interface {
	CreateService(*model.Notification) error
	ReadService(int) ([]model.Notification, error)
	UpdateService(int) error
	NotificationIsReadService(int) (bool, error)
}

type NotificationService struct {
	notificationRepository INotificationRepository
}

func InitNotificationService(notificationRepository INotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
	}
}

func (ns *NotificationService) CreateService(n *model.Notification) error {
	return ns.notificationRepository.Create(n)
}

func (ns *NotificationService) ReadService(u_id int) ([]model.Notification, error) {
	return ns.notificationRepository.Read(u_id)
}

func (ns *NotificationService) UpdateService(u_id int) error {
	return ns.notificationRepository.Update(u_id)
}

func (ns *NotificationService) NotificationIsReadService(u_id int) (bool, error) {
	return ns.notificationRepository.NotificationIsRead(u_id)
}
