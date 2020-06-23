package main

import (
	"log"
	"os"

	"github.com/gabrieldavid98/car-control-access/infrastructure/data"
	"github.com/gabrieldavid98/car-control-access/infrastructure/data/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := mysql.GetDB(os.Getenv("DB_URL"))
	data.Migrate(db)
}
