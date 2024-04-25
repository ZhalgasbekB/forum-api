package admin

import "gitea.com/lzhuk/forum/internal/model"

type IAdminRepository interface {
	Users() ([]model.User, error)
	UpdateUser(*model.User) error
	DeleteUser(int) error
	UpdateUserNewDate(*model.User) error
	CreateReportRepository(*model.ReportCreateDTO) error
	ReportsByStatus() ([]model.Report, error)
	UpdateReport(*model.ReportResponseDTO) error
}

type IAdminService interface {
	UsersService() ([]model.User, error)
	UpdateUserService(*model.User) error
	DeleteUserService(int) error
	UpdateUserNewDateService(*model.User) error
	CreateReportService(*model.ReportCreateDTO) error
	ReportsModeratorService() ([]model.Report, error)
	UpdateReportService(*model.ReportResponseDTO) error
}

type AdminService struct {
	iAdminRepository IAdminRepository
}

func NewAdminService(iAdminRepository IAdminRepository) *AdminService {
	return &AdminService{
		iAdminRepository: iAdminRepository,
	}
}

func (as *AdminService) UsersService() ([]model.User, error) {
	return as.iAdminRepository.Users()
}

func (as *AdminService) UpdateUserService(us *model.User) error {
	return as.iAdminRepository.UpdateUser(us)
}

func (as *AdminService) DeleteUserService(id int) error {
	return as.iAdminRepository.DeleteUser(id)
}

func (as *AdminService) UpdateUserNewDateService(user *model.User) error {
	return as.iAdminRepository.UpdateUserNewDate(user)
}

func (as *AdminService) CreateReportService(r *model.ReportCreateDTO) error {
	return as.iAdminRepository.CreateReportRepository(r)
}

func (as *AdminService) ReportsModeratorService() ([]model.Report, error) {
	return as.iAdminRepository.ReportsByStatus()
}

func (as *AdminService) UpdateReportService(update *model.ReportResponseDTO) error {
	return as.iAdminRepository.UpdateReport(update)
}
