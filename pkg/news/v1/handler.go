package newsv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/football-news-backend/pkg/db"
	"github.com/football-news-backend/pkg/models"
)

func NewsV1Handler(w http.ResponseWriter, req *http.Request) {
	var news models.NewsV1Response
	values := req.URL.Query()
	dateValues := values["date"]
	if len(dateValues) <= 0 {
		allNews, err := db.DBInstance.GetNews()
		if err != nil {
			log.Println(err)
		}
		news.Results = allNews
	} else {
		if !validateDate(dateValues[0]) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "The valid date format is DD.MM.YYYY")
			return
		}
		datedNews, err := db.DBInstance.GetDatedNews(dateValues[0])
		if err != nil {
			log.Println(err)
		}
		news.Results = datedNews
	}

	out, err := json.Marshal(news)
	if err != nil {
		log.Println("Failed to marshall response for /news/v1")
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}

func validateDate(date string) bool {
	// https://regex101.com/r/QgSa5i/1
	regex := regexp.MustCompile(`[0-9]{2}\.[0-9]{2}\.[0-9]{4}`)
	return regex.MatchString(date)
}
