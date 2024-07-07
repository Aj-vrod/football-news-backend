package main

import (
	"log"

	"github.com/football-news-backend/pkg/db"
	"github.com/football-news-backend/pkg/server"
)

func main() {
	log.Println("Starting database...")
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All ready!")

	log.Println("Listening in port :8080")
	server.InitServer()

}
