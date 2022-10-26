package repository

import (
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (commentRepository *CommentRepositoryImpl) Create(comment *domain.Comment) (err error) {

	if err = commentRepository.DB.Create(&comment).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryImpl) GetAll() (comments []domain.Comment, err error) {

	if err = commentRepository.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "user_id", "title", "photo_url", "caption")
	}).Find(&comments).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryImpl) GetOne(id uint) (comment domain.Comment, err error) {

	if err = commentRepository.DB.First(&comment, id).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryImpl) Update(comment domain.Comment) (photo domain.Photo, err error) {

	var c domain.Comment

	if err = commentRepository.DB.First(&c, comment.ID).Error; err != nil {
		return
	}

	if err = commentRepository.DB.Model(&c).Updates(comment).Error; err != nil {
		return
	}

	if err = commentRepository.DB.First(&photo, c.PhotoID).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryImpl) Delete(id uint) (err error) {

	if err = commentRepository.DB.First(&domain.Comment{}, id).Error; err != nil {
		return
	}

	if err = commentRepository.DB.Delete(&domain.Comment{}, id).Error; err != nil {
		return
	}

	return
}
