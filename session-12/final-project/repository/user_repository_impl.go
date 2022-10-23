package repository

import (
	"context"
	"errors"
	"time"

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

func (repository *UserRepositoryImpl) Register(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = repository.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	password := user.Password

	err = repository.DB.WithContext(ctx).Where("email = ?", user.Email).Take(&user).Error
	isValid := helper.Compare([]byte(user.Password), []byte(password))

	if err != nil || !isValid {
		return errors.New("invalid email or password")
	}

	return
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user *domain.User) (updatedUser domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updatedUser = domain.User{}

	if err = repository.DB.WithContext(ctx).First(&updatedUser, &user.ID).Error; err != nil {
		return updatedUser, err
	}

	if err = repository.DB.WithContext(ctx).Model(&updatedUser).Updates(user).Error; err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = repository.DB.WithContext(ctx).First(&domain.User{}, &id).Error; err != nil {
		return err
	}

	if err = repository.DB.WithContext(ctx).Delete(&domain.User{}, &id).Error; err != nil {
		return err
	}

	return
}
