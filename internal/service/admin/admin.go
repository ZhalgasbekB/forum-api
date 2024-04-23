package admin

type IAdminRepository interface{}

type IAdminService interface{}

type AdminService struct {
	iAdminRepository IAdminRepository
}

func NewAdminService(iAdminRepository IAdminRepository) *AdminService {
	return &AdminService{
		iAdminRepository: iAdminRepository,
	}
}

func (as *AdminService) Users() error      { return nil }
func (as *AdminService) UpdateUser() error { return nil }
func (as *AdminService) DeleteUser() error { return nil }
