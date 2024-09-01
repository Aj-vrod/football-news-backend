package storer

import (
	"log"

	"github.com/football-news-backend/pkg/db"
	"github.com/football-news-backend/pkg/models"
)

// StoreNews inserts every piece of news into the database
func StoreNews(latestNews []models.NewsItem) {
	for _, n := range latestNews {
		log.Printf("Inserting news with title: %s \n", n.Title)
		if err := db.DBInstance.CreateNews(n); err != nil {
			log.Printf("Failed to insert news with error: %v", err)
		}
	}
}
