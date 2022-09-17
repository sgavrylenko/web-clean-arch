package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sgavrylenko/web-clean-arch/model"
	"log"
	"os"
)

func DB() *gorm.DB {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", username+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db
}
