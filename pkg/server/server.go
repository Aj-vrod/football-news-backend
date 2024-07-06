package server

import (
	"net/http"

	newsv1 "github.com/football-news-backend/pkg/news/v1"
	muxtrace "github.com/gorilla/mux"
)

func InitServer() {
	mux := muxtrace.NewRouter()
	mux.HandleFunc("/news/v1", newsv1.NewsV1Handler)

	http.ListenAndServe(":8080", mux)
}
