package repository

import "github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"

type CommentRepository interface {
	Create(comment *domain.Comment) (err error)
	GetAll() (comments []domain.Comment, err error)
	GetOne(id uint) (comment domain.Comment, err error)
	Update(comment domain.Comment) (photo domain.Photo, err error)
	Delete(id uint) (err error)
}
