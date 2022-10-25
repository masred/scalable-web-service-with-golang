package repository

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{DB: db}
}

func (photoRepository *PhotoRepositoryImpl) Create(photo *domain.Photo) (err error) {

	if err = photoRepository.DB.Create(&photo).Error; err != nil {
		return
	}

	return
}

func (photoRepository *PhotoRepositoryImpl) GetAll() (photos []domain.Photo, err error) {

	if err = photoRepository.DB.Preload("User").Find(&photos).Error; err != nil {
		return
	}

	return
}

func (photoRepository *PhotoRepositoryImpl) GetOne(id uint) (photo domain.Photo, err error) {

	if err = photoRepository.DB.First(&photo, id).Error; err != nil {
		return
	}

	return
}

func (photoRepository *PhotoRepositoryImpl) Update(photo domain.Photo) (updatedPhoto domain.Photo, err error) {

	if err = photoRepository.DB.First(&updatedPhoto, photo.ID).Error; err != nil {
		return
	}

	if err = photoRepository.DB.Model(&updatedPhoto).Updates(photo).Error; err != nil {
		return
	}

	return
}

func (photoRepository *PhotoRepositoryImpl) Delete(id uint) (err error) {

	if err = photoRepository.DB.First(&domain.Photo{}, id).Error; err != nil {
		return
	}

	if err = photoRepository.DB.Delete(&domain.Photo{}, id).Error; err != nil {
		return
	}

	return
}
