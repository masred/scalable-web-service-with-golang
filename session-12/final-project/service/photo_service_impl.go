package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
)

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &PhotoServiceImpl{PhotoRepository: photoRepository}
}

func (photoService *PhotoServiceImpl) Create(photo *domain.Photo) (err error) {

	if err = photoService.PhotoRepository.Create(photo); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) GetAll() (photos []domain.Photo, err error) {

	if photos, err = photoService.PhotoRepository.GetAll(); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) GetOne(id uint) (photos domain.Photo, err error) {

	if photos, err = photoService.PhotoRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) Update(photo domain.Photo) (updatedPhoto domain.Photo, err error) {

	if updatedPhoto, err = photoService.PhotoRepository.Update(photo); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceImpl) Delete(id uint) (err error) {

	if err = photoService.PhotoRepository.Delete(id); err != nil {
		return
	}

	return
}
