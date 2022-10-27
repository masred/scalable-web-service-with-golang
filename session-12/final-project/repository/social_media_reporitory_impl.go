package repository

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"gorm.io/gorm"
)

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{DB: db}
}

func (socialMediaRepository *SocialMediaRepositoryImpl) Create(socialMedia *domain.SocialMedia) (err error) {

	if err = socialMediaRepository.DB.Create(&socialMedia).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryImpl) GetAll() (socialMedias []domain.SocialMedia, err error) {

	if err = socialMediaRepository.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Find(&socialMedias).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryImpl) GetOne(id uint) (socialMedia domain.SocialMedia, err error) {

	if err = socialMediaRepository.DB.Find(&socialMedia, id).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryImpl) Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error) {

	if err = socialMediaRepository.DB.First(&updatedSocialMedia, socialMedia.ID).Error; err != nil {
		return
	}

	if err = socialMediaRepository.DB.Model(&updatedSocialMedia).Updates(socialMedia).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryImpl) Delete(id uint) (err error) {

	if err = socialMediaRepository.DB.Delete(&domain.SocialMedia{}, id).Error; err != nil {
		return
	}

	return
}
