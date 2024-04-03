package main

import (
	"github.com/joho/godotenv"
	"log"
)

//TODO change func names

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := InitRouter()
	err = router.Run(":8080")
	if err != nil {
		return
	}

}
