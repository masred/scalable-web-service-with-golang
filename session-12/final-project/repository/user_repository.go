package repository

import (
	"context"

	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, user *domain.User) (err error)
	Login(ctx context.Context, user *domain.User) (err error)
	Update(ctx context.Context, user *domain.User) (u domain.User, err error)
	Delete(ctx context.Context, id string) (err error)
}
