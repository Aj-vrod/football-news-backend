package main

import newsv1 "github.com/football-news-backend/pkg/news/v1"

func main() {
	// fmt.Println("Listening in port :8080")
	// server.InitServer()
	newsv1.Scraper()
}
