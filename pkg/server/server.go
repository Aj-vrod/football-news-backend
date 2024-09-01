package server

import (
	"log"
	"net/http"

	"github.com/football-news-backend/cmd"
	askv1 "github.com/football-news-backend/pkg/ask/v1"
	newsv1 "github.com/football-news-backend/pkg/news/v1"
	muxtrace "github.com/gorilla/mux"
)

func InitServer() {
	mux := muxtrace.NewRouter()
	v1 := mux.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/news", newsv1.NewsV1Handler).Methods("GET")
	v1.HandleFunc("/ask", askv1.AskV1Handler).Methods("GET")
	v1.HandleFunc("/trigger-reporter", cmd.ReportHandler).Methods("GET")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
