package cmd

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	newsv1 "github.com/football-news-backend/pkg/news/v1"
	"github.com/football-news-backend/pkg/storer"
)

// ReportNews executes the scrapper and storer to retrieve and persist latest news
func ReportNews() error {
	latestNews := newsv1.ScrapeNews()
	if len(latestNews) == 0 {
		return errors.New("reporter returned no news")
	}

	storer.StoreNews(latestNews)
	log.Println("End of reporting!")

	return nil
}

// Temporal handler until moving to a cronjob
func ReportHandler(w http.ResponseWriter, _ *http.Request) {
	err := ReportNews()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	successMessage, _ := json.Marshal(map[string]string{"success": "Reporter triggered"})
	w.Write(successMessage)
}
