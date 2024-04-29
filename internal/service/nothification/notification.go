package nothification

import "gitea.com/lzhuk/forum/internal/model"

type INotificationRepository interface {
	Create(*model.Notification) error
	Read(int) ([]model.Notification, error)
	Update() error
	NotificationIsRead() (bool, error)
}

type INotificationService interface {
	CreateService(*model.Notification) error
	ReadService(int) ([]model.Notification, error)
	UpdateService() error
	NotificationIsReadService() (bool, error)
}

type NotificationService struct {
	notificationRepository INotificationRepository
}

func InitNothificationService(nothificationRepository INotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository: nothificationRepository,
	}
}

func (ns *NotificationService) CreateService(n *model.Notification) error {
	return ns.notificationRepository.Create(n)
}

func (ns *NotificationService) ReadService(id int) ([]model.Notification, error) {
	return ns.notificationRepository.Read(id)
}

func (ns *NotificationService) UpdateService() error {
	return ns.notificationRepository.Update()
}

func (ns *NotificationService) NotificationIsReadService() (bool, error) {
	return ns.notificationRepository.NotificationIsRead()
}
