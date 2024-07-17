package askv1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/football-news-backend/pkg/llm"
	"github.com/gorilla/mux"
)

type askV1Response struct {
	answer string
}

type askV1FailureResponse struct {
	warning string
}

func AskV1Handler(w http.ResponseWriter, req *http.Request) {
	gemini := initGemini()
	query := mux.Vars(req)["query"]

	if query == "" {
		json.NewEncoder(w).Encode(askV1FailureResponse{warning: "Missing query for your request."})
	}

	ans := gemini.GenerateAnswer(query)
	json.NewEncoder(w).Encode(askV1Response{answer: ans})
}

func initGemini() llm.GeminiClient {
	log.Println("Starting gemini client...")
	ctx := context.Background()
	cfg := llm.GenerateConfig(ctx)
	return llm.InitClient(cfg)
}
