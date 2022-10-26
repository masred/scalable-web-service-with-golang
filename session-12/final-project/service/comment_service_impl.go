package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &CommentServiceImpl{CommentRepository: commentRepository}
}

func (commentService *CommentServiceImpl) Create(comment *domain.Comment) (err error) {

	if err = commentService.CommentRepository.Create(comment); err != nil {
		return
	}
	return
}

func (commentService *CommentServiceImpl) GetAll() (comments []domain.Comment, err error) {

	if comments, err = commentService.CommentRepository.GetAll(); err != nil {
		return
	}

	return
}

func (commentService *CommentServiceImpl) GetOne(id uint) (comment domain.Comment, err error) {

	if comment, err = commentService.CommentRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (commentService *CommentServiceImpl) Update(comment domain.Comment) (photo domain.Photo, err error) {

	if photo, err = commentService.CommentRepository.Update(comment); err != nil {
		return
	}

	return
}

func (commentService *CommentServiceImpl) Delete(id uint) (err error) {

	if err = commentService.CommentRepository.Delete(id); err != nil {
		return
	}

	return
}
