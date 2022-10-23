package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) Register(user *domain.User) (err error) {
	if err = service.UserRepository.Register(user); err != nil {
		return err
	}

	return
}

func (service *UserServiceImpl) Login(user *domain.User) (err error) {
	if err = service.UserRepository.Login(user); err != nil {
		return err
	}

	return
}

func (service *UserServiceImpl) Update(user *domain.User) (u domain.User, err error) {
	if u, err = service.UserRepository.Update(user); err != nil {
		return u, err
	}

	return u, nil
}

func (service *UserServiceImpl) Delete(id uint) (err error) {
	if err = service.UserRepository.Delete(id); err != nil {
		return err
	}

	return
}
