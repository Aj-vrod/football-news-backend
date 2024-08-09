package storer

import (
	"github.com/football-news-backend/pkg/db"
	newsv1 "github.com/football-news-backend/pkg/news/v1"
)

// SyncNews calls scraper to get the latest news and stores them in the db
func SyncNews() {
	latestNews := newsv1.ScraperNews()

	for _, n := range latestNews {
		db.DBInstance.CreateNews(n)
	}
}
