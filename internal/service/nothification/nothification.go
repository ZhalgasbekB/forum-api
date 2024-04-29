package nothification

import "gitea.com/lzhuk/forum/internal/model"

type INothificationRepository interface {
	Create() error
	Read() ([]model.Nothification, error)
	Update() error
	NothificationIsRead() (bool, error)
}

type INothificationService interface {
	CreateService() error
	ReadService() ([]model.Nothification, error)
	UpdateService() error
	NothificationIsReadService() (bool, error)
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

func (ns *NothificationService) UpdateService() error {
	return ns.nothificationRepository.Update()
}

func (ns *NothificationService) NothificationIsReadService() (bool, error) {
	return ns.nothificationRepository.NothificationIsRead()
}
