package controller

import (
	"crud/intializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")

	user := intializer.User{
		UserName: username,
		Email:    email,
	}
	result := intializer.DB.Create(&user)

	if result.Error != nil {
		c.String(500, "error saving user")
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}
func View(c *gin.Context) {
	var users []intializer.User
	intializer.DB.Find(&users)
	c.HTML(http.StatusOK, "view.html", gin.H{
		"users": users,
	})
	c.Redirect(http.StatusSeeOther, "/view")
}
func Take(c *gin.Context) {
	id := c.Param("id")

	var user intializer.User
	if err := intializer.DB.First(&user, id).Error; err != nil {
		c.String(http.StatusNotFound, "user not found")
		return
	}
	c.HTML(http.StatusOK, "update.html", gin.H{
		"user": user,
	})
}
func Update(c *gin.Context) {
	id := c.Param("id")
	email := c.PostForm("email")
	username := c.PostForm("username")

	newresult := intializer.DB.Model(&intializer.User{}).
		Where("id=?", id).
		Updates(intializer.User{
			UserName: username,
			Email:    email,
		})
	if newresult.Error != nil {
		c.String(http.StatusInternalServerError, "server error")
		return
	}
	c.Redirect(http.StatusSeeOther, "/view")
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	delete := intializer.DB.Delete(&intializer.User{}, id)

	if delete.Error != nil {
		c.String(http.StatusInternalServerError, "unable to delete data")
		return
	}
	c.Redirect(http.StatusSeeOther, "/view")

}
