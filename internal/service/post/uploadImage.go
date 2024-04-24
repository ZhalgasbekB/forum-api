package post

import "gitea.com/lzhuk/forum/internal/model"

type IUploadImagePostRepository interface {
	AddImagePostRepository(image *model.UploadPostImage) error
	GetImagePostRepository(postId int) (string, error)
	UpdateImagePostRepository(image *model.UploadPostImage) error
	DeleteImagePostRepository(postId int) error
}

type UploadImagePostService struct {
	uploadImagePostRepository IUploadImagePostRepository
}

func NewUploadImagePostService(uploadImagePostRepository IUploadImagePostRepository) *UploadImagePostService {
	return &UploadImagePostService{
		uploadImagePostRepository: uploadImagePostRepository,
	}
}

func (u *UploadImagePostService) AddImagePostService(image *model.UploadPostImage) error {
	return u.uploadImagePostRepository.AddImagePostRepository(image)
}

func (u *UploadImagePostService) GetImagePostService(postId int) (string, error) {
	return u.uploadImagePostRepository.GetImagePostRepository(postId)
}

func (u *UploadImagePostService) UpdateImageService(image *model.UploadPostImage) error {
	return u.uploadImagePostRepository.UpdateImagePostRepository(image)
}

func (u *UploadImagePostService) DeleteImagePostService(postId int) error {
	return u.uploadImagePostRepository.DeleteImagePostRepository(postId)
}
