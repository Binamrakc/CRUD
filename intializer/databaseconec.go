package intializer

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"primaryKey"`
	UserName string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed ", err)
	}
	DB = db

	log.Println("database created and runed sucessfully")
}
