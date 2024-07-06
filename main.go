package main

import (
	"fmt"

	"github.com/football-news-backend/pkg/server"
)

func main() {
	fmt.Println("Listening in port :8080")
	server.InitServer()

}
