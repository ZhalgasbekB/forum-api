package nothification

type INothificationRepository interface {
	Create() error
	Read() error
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
	return nil
}

func (ns *NothificationService) ReadService() error {
	return nil
}

func (ns *NothificationService) DeleteService() error {
	return nil
}

func (ns *NothificationService) UpdateService() error {
	return nil
}
