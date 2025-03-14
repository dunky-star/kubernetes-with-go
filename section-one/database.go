package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
}

type Client struct {
	db *gorm.DB
}

// Ready method we perform a RAW SQL query to check database readiness. It will return a boolean response (true or false).
func (c Client) Ready() bool {
	var ready string
	result := c.db.Raw("SELECT 1 as ready").Scan(&ready)
	if result.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func (c Client) RunMigration() error {
	if !c.Ready() {
		log.Fatal("Database is not ready")
	}
	err := c.db.AutoMigrate(&User{}) // Model to be added
	if err != nil {
		return err
	}
	return nil
}

// NewDBClient function constructs a database connection from environment variables,
// creating a Client that can be used to interact with the database.
func NewDBClient() (Client, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Error, Invalid DB Port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbHost, dbUsername, dbPassword, dbName, databasePort, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Client{}, err
	}
	client := Client{db}
	return client, nil
}
