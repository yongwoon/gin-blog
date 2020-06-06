package initializer

import (
	"os"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	currentEnvironment, ok := os.LookupEnv("ENVIRONMENT")

	var err error
	if ok == true {
		err = godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/yongwoon/gin-blog/.env." + currentEnvironment))
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		panic(err)
	}
}
