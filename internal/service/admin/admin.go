package admin

import (
	"gitea.com/lzhuk/forum/internal/helpers/roles"
	"gitea.com/lzhuk/forum/internal/model"
)

type IAdminRepository interface {
	Users() ([]model.User, error)
	UpdateUser(*model.User) error
	DeleteUser(int) error
	DeletePost(int) error
	DeleteComment(int) error

	CreateCategory(string) error
	DeleteCategory(string) error

	CreateReportRepository(*model.ReportCreateDTO) error
	ReportsByStatus() ([]model.Report, error)
	UpdateReport(*model.ReportResponseDTO) error

	UserWantsRepository(*model.WantsDTO) error
	UserWants() ([]model.WantsDTO, error)

	UpdateWantUser(u *model.AdminResponse) error
}

type IAdminService interface {
	UsersService() ([]model.User, error)
	UpdateUserService(*model.User) error
	DeleteUserService(int) error
	DeletePostService(int) error
	DeleteCommentService(int) error

	CreateCategoryService(string) error
	DeleteCategoryService(string) error

	CreateReportService(*model.ReportCreateDTO) error
	ReportsModeratorService() ([]model.Report, error)
	UpdateReportService(*model.ReportResponseDTO) error

	UserWantsService(*model.WantsDTO) error
	UsersWantsService() ([]model.WantsDTO, error)
	UserWantRoleAdminResponseService(user *model.AdminResponse) error
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

func (as *AdminService) DeletePostService(id int) error {
	return as.iAdminRepository.DeletePost(id)
}

func (as *AdminService) DeleteCommentService(id int) error {
	return as.iAdminRepository.DeleteComment(id)
}

func (as *AdminService) CreateReportService(r *model.ReportCreateDTO) error {
	return as.iAdminRepository.CreateReportRepository(r)
}

func (as *AdminService) ReportsModeratorService() ([]model.Report, error) {
	return as.iAdminRepository.ReportsByStatus()
}

func (as *AdminService) CreateCategoryService(category string) error {
	return as.iAdminRepository.CreateCategory(category)
}

func (as *AdminService) DeleteCategoryService(category string) error {
	return as.iAdminRepository.DeleteCategory(category)
}

func (as *AdminService) UpdateReportService(update *model.ReportResponseDTO) error {
	return as.iAdminRepository.UpdateReport(update)
}

func (as *AdminService) UserWantsService(m *model.WantsDTO) error {
	return as.iAdminRepository.UserWantsRepository(m)
}

func (as *AdminService) UsersWantsService() ([]model.WantsDTO, error) {
	return as.iAdminRepository.UserWants()
}

func (as *AdminService) UserWantRoleAdminResponseService(adminR *model.AdminResponse) error {
	if adminR.Status == 1 {
		if err := as.iAdminRepository.UpdateUser(&model.User{ID: adminR.UserID, Role: roles.MODERATOR}); err != nil {
			return err
		}
	}
	return as.iAdminRepository.UpdateWantUser(adminR) /// SOME CODE SAYS REJECTED OR APROUVE
}
