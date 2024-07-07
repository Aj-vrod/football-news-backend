package main

import (
	"fmt"

	"github.com/football-news-backend/pkg/db"
	"github.com/football-news-backend/pkg/server"
)

func main() {
	fmt.Println("Starting database...")
	db.InitDB()

	fmt.Println("Listening in port :8080")
	server.InitServer()

}
