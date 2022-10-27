package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository) SocialMediaService {
	return &SocialMediaServiceImpl{SocialMediaRepository: socialMediaRepository}
}

func (socialMediaService *SocialMediaServiceImpl) Create(socialMedia *domain.SocialMedia) (err error) {

	if err = socialMediaService.SocialMediaRepository.Create(socialMedia); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceImpl) GetAll() (socialMedias []domain.SocialMedia, err error) {

	if socialMedias, err = socialMediaService.SocialMediaRepository.GetAll(); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceImpl) GetOne(id uint) (socialMedia domain.SocialMedia, err error) {

	if socialMedia, err = socialMediaService.SocialMediaRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceImpl) Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error) {

	if updatedSocialMedia, err = socialMediaService.SocialMediaRepository.Update(socialMedia); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceImpl) Delete(id uint) (err error) {

	if err = socialMediaService.SocialMediaRepository.Delete(id); err != nil {
		return
	}

	return
}
