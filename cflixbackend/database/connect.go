package database

import (
	"github.com/TechMaster/golang/08Fiber/Repository/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:@/cflix"), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to the database")
	}
	DB = database
	DB.AutoMigrate(models.Film{})
	DB.AutoMigrate(models.Category{})
	DB.AutoMigrate(models.Filmcategory{})
}
