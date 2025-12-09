package intializer

import "log"

func DBmigrate() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Println("auto migration failed", err)
		return
	}

}
