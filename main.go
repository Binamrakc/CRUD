package main

import (
	"crud/controller"
	"crud/intializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	intializer.Loadenv()
	intializer.ConnectDB()
	intializer.DBmigrate()
	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "userlogin",
		})
	})

	r.POST("/submit", controller.Submit)
	r.GET("/view", controller.View)

	r.GET("/update/:id", controller.Take)
	r.POST("/update/:id", controller.Update)

	r.POST("/delete/:id", controller.Delete)
	r.GET("/register", controller.Showregister)
	r.POST("/register", controller.Register)

	r.GET("/auth", controller.Showauth)
	r.Run()
}
