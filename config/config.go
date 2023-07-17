package config

import (
	"fmt"
	"os"
	"golang/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	loadEnv()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migration()
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Error load .env file")
  	}
}

func migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Film{})
}
