package newsv1

import (
	"encoding/json"
	"net/http"

	"github.com/football-news-backend/pkg/models"
)

func NewsV1Handler(w http.ResponseWriter, _ *http.Request) {
	res := models.NewsV1Response{
		Success: true,
	}
	out, err := json.Marshal(res)
	if err != nil {
		panic("OH NO!")
	}

	w.Write(out)
}
