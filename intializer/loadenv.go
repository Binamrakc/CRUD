package intializer

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(http.StatusInternalServerError, "database not connected")
	}
}
