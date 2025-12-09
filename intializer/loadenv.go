package intializer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Loadenv(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		c.String(http.StatusInternalServerError, "database not connected")
	}
}
