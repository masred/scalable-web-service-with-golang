package app

import (
	"fmt"

	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/helper"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	pass   = "root"
	port   = "5432"
	dbName = "assignment_2"
)

func NewDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(config))
	helper.PanicIfError(err)

	err = db.AutoMigrate(&domain.Item{}, &domain.Order{})
	helper.PanicIfError(err)

	return db
}
