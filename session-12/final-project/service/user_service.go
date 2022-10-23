package service

import (
	"context"

	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
)

type UserService interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) (domain.User, error)
	Delete(ctx context.Context, id uint) error
}
