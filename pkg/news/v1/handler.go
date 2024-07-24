package newsv1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/football-news-backend/pkg/models"
)

func NewsV1Handler(w http.ResponseWriter, _ *http.Request) {
	news := ScraperNews()
	res := models.NewsV1Response{
		Results: news,
	}
	out, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to marshall response for /news/v1")
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
