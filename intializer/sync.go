package intializer

import "log"

func DBmigrate() {
	err := DB.AutoMigrate(&User{}, &Auth{})
	if err != nil {
		log.Println("auto migration failed", err)
		return
	}

}
