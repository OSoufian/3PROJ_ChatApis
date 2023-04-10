package main

import (
	"fmt"
	"log"
	"os"

	"chatsapi/internal/domain"
	"chatsapi/internal/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {

	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load(".env")
	}

	postgresHost := os.Getenv("PostgresHost")
	postgresUser := os.Getenv("PostgresUser")
	postgresPassword := os.Getenv("PostgresPassword")
	postgresDatabase := os.Getenv("PostgresDatabase")
	postgresPort := os.Getenv("PostgresPort")

	appListen := os.Getenv("AppListen")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort)
	domain.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	domain.Db.AutoMigrate(&domain.LiveMessage{})
	domain.Db.AutoMigrate(&domain.Message{})

	// Start Fiber app
	http.Http().Listen(appListen)
}
