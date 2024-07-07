package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitDB() {
	DB_DNS := os.Getenv("POSTGRESQL_DNS")
	MIGRATIONS_DIR := os.Getenv("MIGRATIONS_PATH")

	if MIGRATIONS_DIR == "" {
		fmt.Println("Missing env variable MIGRATIONS_PATH")
		return
	}

	if DB_DNS == "" {
		fmt.Println("Missing env variable POSTGRESQL_DNS")
		return
	}

	db, err := sql.Open("postgres", DB_DNS)
	if err != nil {
		fmt.Println("Failed to connect to db: ", err)
		return
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println("Failed to get driver: ", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		MIGRATIONS_DIR,
		"football", driver)
	if err != nil {
		fmt.Println("Failed to get migrate instance: ", err)
		return
	}

	err = m.Up()
	if err != nil {
		fmt.Println("Failed to run migrations: ", err)
		return
	}
}
