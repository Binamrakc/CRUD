package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"primaryKey"`
	UserName string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}

var DB *gorm.DB

func Database() {
	dsn := "gouser:go12345!@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection failed ", err)
	}
	DB = db

	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Println("auto migration failed", err)
		return
	}

	log.Println("database created and runed sucessfully")
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	Database()
	r.LoadHTMLGlob("templates/*")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "userlogin",
		})
	})
	r.POST("/submit", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")

		user := User{
			UserName: username,
			Email:    email,
		}
		result := DB.Create(&user)

		if result.Error != nil {
			c.String(500, "error saving user")
			return
		}

		c.Redirect(http.StatusSeeOther, "/home")
	})
	r.GET("/view", func(c *gin.Context) {
		var users []User
		DB.Find(&users)
		c.HTML(http.StatusOK, "view.html", gin.H{
			"users": users,
		})
		c.Redirect(http.StatusSeeOther, "/view")
	})
	r.GET("/update/:id", func(c *gin.Context) {
		id := c.Param("id")

		var user User
		if err := DB.First(&user, id).Error; err != nil {
			c.String(http.StatusNotFound, "user not found")
			return
		}
		c.HTML(http.StatusOK, "update.html", gin.H{
			"user": user,
		})
	})
	r.POST("/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		email := c.PostForm("email")
		username := c.PostForm("username")

		newresult := DB.Model(&User{}).
			Where("id=?", id).
			Updates(User{
				UserName: username,
				Email:    email,
			})
		if newresult.Error != nil {
			c.String(http.StatusInternalServerError, "server error")
			return
		}
		c.Redirect(http.StatusSeeOther, "/view")
	})
	r.Run()
}
