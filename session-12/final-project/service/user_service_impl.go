package service

import (
	"context"

	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) Register(ctx context.Context, user *domain.User) (err error) {
	if err = service.UserRepository.Register(ctx, user); err != nil {
		return err
	}

	return
}

func (service *UserServiceImpl) Login(ctx context.Context, user *domain.User) (err error) {
	if err = service.UserRepository.Login(ctx, user); err != nil {
		return err
	}

	return
}

func (service *UserServiceImpl) Update(ctx context.Context, user *domain.User) (u domain.User, err error) {
	if u, err = service.UserRepository.Update(ctx, user); err != nil {
		return u, err
	}

	return u, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, id uint) (err error) {
	if err = service.UserRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
