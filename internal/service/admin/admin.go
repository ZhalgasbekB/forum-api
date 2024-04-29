package admin

import (
	"gitea.com/lzhuk/forum/internal/helpers/roles"
	"gitea.com/lzhuk/forum/internal/model"
)

type IAdminRepository interface {
	Users() ([]model.User, error)
	ChangeRole(*model.User) error
	DeleteUser(int) error
	DeletePost(int) error
	DeleteComment(int) error

	CreateCategory(string) error
	DeleteCategory(string) error

	CreateReportModerator(*model.ReportCreateDTO) error
	ReportsByStatus() ([]model.Report, error)
	ResponseReportAdmin(*model.ReportResponseDTO) error

	UserWant(*model.User) error
	UsersWantRole() ([]model.Wants2DTO, error)
	UpdateUserWantStatus(*model.AdminResponse) error

	MonderatorReports(int) ([]model.Report, error)
	UserWants(int) ([]model.Wants1DTO, error)
}

type IAdminService interface {
	UsersService() ([]model.User, error)
	ChangeRoleService(*model.User) error
	DeleteUserService(int) error
	DeletePostService(int) error
	DeleteCommentService(int) error

	CreateCategoryService(string) error
	DeleteCategoryService(string) error

	CreateReportModeratorService(*model.ReportCreateDTO) error
	ReportsByStatusService() ([]model.Report, error)
	ResponseReportAdminService(*model.ReportResponseDTO) error

	UserWantService(*model.User) error
	UsersWantRoleService() ([]model.Wants2DTO, error)
	UpdateUserWantStatusService(user *model.AdminResponse) error

	MonderatorReportsService(int) ([]model.Report, error)
	UserWantsService(int) ([]model.Wants1DTO, error)
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

func (as *AdminService) ChangeRoleService(us *model.User) error {
	return as.iAdminRepository.ChangeRole(us)
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

func (as *AdminService) CreateReportModeratorService(r *model.ReportCreateDTO) error {
	return as.iAdminRepository.CreateReportModerator(r)
}

func (as *AdminService) ReportsByStatusService() ([]model.Report, error) {
	return as.iAdminRepository.ReportsByStatus()
}

func (as *AdminService) ResponseReportAdminService(update *model.ReportResponseDTO) error {
	return as.iAdminRepository.ResponseReportAdmin(update)
}

func (as *AdminService) CreateCategoryService(category string) error {
	return as.iAdminRepository.CreateCategory(category)
}

func (as *AdminService) DeleteCategoryService(category string) error {
	return as.iAdminRepository.DeleteCategory(category)
}

func (as *AdminService) UserWantService(m *model.User) error {
	return as.iAdminRepository.UserWant(m)
}

func (as *AdminService) UsersWantRoleService() ([]model.Wants2DTO, error) {
	return as.iAdminRepository.UsersWantRole()
}

func (as *AdminService) UpdateUserWantStatusService(adminR *model.AdminResponse) error {
	if adminR.Status == 1 {
		if err := as.iAdminRepository.ChangeRole(&model.User{ID: adminR.UserID, Role: roles.MODERATOR}); err != nil {
			return err
		}
	}
	return as.iAdminRepository.UpdateUserWantStatus(adminR)
}

func (as *AdminService) MonderatorReportsService(m int) ([]model.Report, error) {
	return as.iAdminRepository.MonderatorReports(m)
}

func (as *AdminService) UserWantsService(u int) ([]model.Wants1DTO, error) {
	return as.iAdminRepository.UserWants(u)
}
