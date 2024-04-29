package nothification

import "gitea.com/lzhuk/forum/internal/model"

type INothificationRepository interface {
	Create() error
	Read() ([]model.Nothification, error)
	Delete() error
	Update() error
}

type INothificationService interface {
	CreateService() error
	ReadService() error
	DeleteService() error
	UpdateService() error
}

type NothificationService struct {
	nothificationRepository INothificationRepository
}

func InitNothificationService(nothificationRepository INothificationRepository) *NothificationService {
	return &NothificationService{
		nothificationRepository: nothificationRepository,
	}
}

func (ns *NothificationService) CreateService() error {
	return ns.nothificationRepository.Create()
}

func (ns *NothificationService) ReadService() ([]model.Nothification, error) {
	return ns.nothificationRepository.Read()
}

func (ns *NothificationService) DeleteService() error {
	return ns.nothificationRepository.Delete()
}

func (ns *NothificationService) UpdateService() error {
	return ns.nothificationRepository.Update()
}
