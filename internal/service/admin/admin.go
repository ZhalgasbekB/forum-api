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

	UserWant(*model.WantsDTO) error
	UsersWantRole() ([]model.WantsDTO, error)
	UpdateUserWantStatus(u *model.AdminResponse) error
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

	UserWantService(*model.WantsDTO) error
	UsersWantRoleService() ([]model.WantsDTO, error)
	UpdateUserWantStatusService(user *model.AdminResponse) error
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

func (as *AdminService) UserWantService(m *model.WantsDTO) error {
	return as.iAdminRepository.UserWant(m)
}

func (as *AdminService) UsersWantRoleService() ([]model.WantsDTO, error) {
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
