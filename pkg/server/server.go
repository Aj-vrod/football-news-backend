package server

import (
	"net/http"

	newsv1 "github.com/football-news-backend/pkg/news/v1"
	muxtrace "github.com/gorilla/mux"
)

func InitServer() {
	mux := muxtrace.NewRouter()
	mux.HandleFunc("v1/news", newsv1.NewsV1Handler).Methods("GET")
	mux.HandleFunc("v1/ask", newsv1.NewsV1Handler).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
