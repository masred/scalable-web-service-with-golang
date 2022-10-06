package config

import (
	"github.com/jinzhu/gorm"
	"github.com/masred/scalable-web-service-with-golang/session-08/3-practice/models"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.Person{})
	return db
}
