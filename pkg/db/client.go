package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

var DBInstance Database

func InitDB() error {
	DB_DNS, MIGRATIONS_DIR, err := getEnvs()
	if err != nil {
		return err
	}

	DBInstance.DB, err = sql.Open("postgres", DB_DNS)
	if err != nil {
		log.Println("Failed to open connection with DB.")
		return err
	}
	driver, err := postgres.WithInstance(DBInstance.DB, &postgres.Config{})
	if err != nil {
		log.Println("Failed to get driver.")
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		MIGRATIONS_DIR,
		"football", driver)
	if err != nil {
		log.Println("Failed to get migrate instance.")
		return err
	}

	err = m.Up()
	if err != nil {
		log.Println("Failed to run migrations.")
		return err
	}

	return nil
}

func getEnvs() (string, string, error) {
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	dbDns := os.Getenv("POSTGRESQL_DNS")
	if migrationsPath == "" {
		return "", "", fmt.Errorf("missing env variable MIGRATIONS_PATH")
	}
	if dbDns == "" {
		return "", "", fmt.Errorf("missing env variable POSTGRESQL_DNS")
	}

	return dbDns, migrationsPath, nil
}
