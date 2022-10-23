package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
)

type UserService interface {
	Register(user *domain.User) (err error)
	Login(user *domain.User) (err error)
	Update(user *domain.User) (updatedUser domain.User, err error)
	Delete(id uint) (err error)
}
