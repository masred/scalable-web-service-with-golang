package service

import (
	"context"

	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
)

type UserService interface {
	Register(context.Context, *domain.User) error
	Login(context.Context, *domain.User) error
	Update(context.Context, *domain.User) (domain.User, error)
	Delete(context.Context, string) error
}
