package nothification

import "gitea.com/lzhuk/forum/internal/model"

type INothificationRepository interface {
	Create(*model.Nothification) error
	Read(int) ([]model.Nothification, error)
	Update() error
	NothificationIsRead() (bool, error)
}

type INothificationService interface {
	CreateService(*model.Nothification) error
	ReadService(int) ([]model.Nothification, error)
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

func (ns *NothificationService) CreateService(n *model.Nothification) error {
	return ns.nothificationRepository.Create(n)
}

func (ns *NothificationService) ReadService(id int) ([]model.Nothification, error) {
	return ns.nothificationRepository.Read(id)
}

func (ns *NothificationService) UpdateService() error {
	return ns.nothificationRepository.Update()
}

func (ns *NothificationService) NothificationIsReadService() (bool, error) {
	return ns.nothificationRepository.NothificationIsRead()
}
