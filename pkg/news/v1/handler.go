package newsv1

import (
	"encoding/json"
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
		panic("OH NO!")
	}

	w.Write(out)
}
