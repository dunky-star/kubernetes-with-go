package main

import (
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	return ready == "1"
}

func (c Client) RunMigration() error {
	if !c.Ready() {
		log.Fatal("Database is not ready")
	}
	var tableExists bool
	err := c.db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = ?)", "users").Scan(&tableExists).Error
	if err != nil {
		log.Fatalf("Error checking table existence: %v", err)
		return err
	}

	if !tableExists {
		log.Println("Table `users` does not exist, running migrations...")
		err = c.db.AutoMigrate(&User{}) // Migrate only if table doesn't exist
		if err != nil {
			log.Fatalf("Error running migrations: %v", err)
			return err
		}
		log.Println("Migration successful!")
	} else {
		log.Println("Table `users` already exists, skipping migration.")
	}

	return nil
}

// NewDBClient function constructs a database connection from environment variables,
// creating a Client that can be used to interact with the database.
func NewDBClient() (Client, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read `DATABASE_URL` from environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("Error: DATABASE_URL is not set")
	}

	// Open the connection using `pgx`
	dbConfig, err := pgx.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Failed to parse DB config: %v", err)
		return Client{}, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*dbConfig),
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
	})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return Client{}, err
	}
	client := Client{db}
	err = client.RunMigration()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	return client, nil
}
