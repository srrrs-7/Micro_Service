package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Connect to PostgreSQL database using gorm
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Query users from database
	var users []User
	err = db.Find(&users).Error
	if err != nil {
		log.Fatalf("Error querying users from database: %s", err)
	}

	// Create CSV file and write users to it
	file, err := os.Create("users.csv")
	if err != nil {
		log.Fatalf("Error creating CSV file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, user := range users {
		record := []string{
			strconv.Itoa(user.ID),
			user.Username,
			user.Email,
		}
		err := writer.Write(record)
		if err != nil {
			log.Fatalf("Error writing record to CSV: %s", err)
		}
	}

	log.Println("Successfully wrote users to CSV file")
}
