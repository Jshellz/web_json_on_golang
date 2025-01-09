package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"web/db/models"
)

var dsn = "host=localhost user=local password=pass dbname=app port=5432 TimeZone=Europe/Moscow"
var DB *gorm.DB

func ToDB() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connected to db")
	}

	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		return
	}

	DB = db
}
