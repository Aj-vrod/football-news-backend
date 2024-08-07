package askv1

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/football-news-backend/pkg/llm"
)

type AskV1Response struct {
	Answer string `json:"answer"`
}

func AskV1Handler(w http.ResponseWriter, req *http.Request) {
	gemini := initGemini()
	values := req.URL.Query()
	queries := values["query"]
	if len(queries) <= 0 {
		fmt.Fprint(w, "Missing query in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ans := gemini.GenerateAnswer(queries[0])
	out := AskV1Response{
		Answer: ans,
	}
	outJSON, _ := json.Marshal(out)
	w.Write(outJSON)
}

func initGemini() llm.GeminiClient {
	log.Println("Starting gemini client...")
	ctx := context.Background()
	cfg := llm.GenerateConfig(ctx)
	return llm.InitClient(cfg)
}
