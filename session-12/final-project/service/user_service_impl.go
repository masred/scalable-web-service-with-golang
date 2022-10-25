package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (userService *UserServiceImpl) Register(user *domain.User) (err error) {
	if err = userService.UserRepository.Register(user); err != nil {
		return err
	}

	return
}

func (userService *UserServiceImpl) Login(user *domain.User) (err error) {
	if err = userService.UserRepository.Login(user); err != nil {
		return err
	}

	return
}

func (userService *UserServiceImpl) Update(user *domain.User) (u domain.User, err error) {
	if u, err = userService.UserRepository.Update(user); err != nil {
		return u, err
	}

	return u, nil
}

func (userService *UserServiceImpl) Delete(id uint) (err error) {
	if err = userService.UserRepository.Delete(id); err != nil {
		return err
	}

	return
}
