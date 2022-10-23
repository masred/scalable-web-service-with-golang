package repository

import (
	"errors"

	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/helper"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Register(user *domain.User) (err error) {
	if err = repository.DB.Create(&user).Error; err != nil {
		return err
	}

	return
}

func (repository *UserRepositoryImpl) Login(user *domain.User) (err error) {
	password := user.Password

	err = repository.DB.Where("email = ?", user.Email).Take(&user).Error
	isValid := helper.Compare([]byte(user.Password), []byte(password))

	if err != nil || !isValid {
		return errors.New("invalid email or password")
	}

	return
}

func (repository *UserRepositoryImpl) Update(user *domain.User) (updatedUser domain.User, err error) {
	updatedUser = domain.User{}

	if err = repository.DB.First(&updatedUser, &user.ID).Error; err != nil {
		return updatedUser, err
	}

	if err = repository.DB.Model(&updatedUser).Updates(user).Error; err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (repository *UserRepositoryImpl) Delete(id uint) (err error) {
	if err = repository.DB.First(&domain.User{}, &id).Error; err != nil {
		return err
	}

	if err = repository.DB.Delete(&domain.User{}, &id).Error; err != nil {
		return err
	}

	return
}
