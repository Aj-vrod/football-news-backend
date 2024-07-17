package llm

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateConfig(context context.Context) GeminiConfig {
	return GeminiConfig{
		ctx: context,
	}
}

func InitClient(cfg GeminiConfig) GeminiClient {
	geminiClient, err := genai.NewClient(cfg.ctx, option.WithAPIKey(os.Getenv("GEMINI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer geminiClient.Close()

	model := geminiClient.GenerativeModel("gemini-1.5-flash")
	return GeminiClient{
		model: model,
		ctx:   cfg.ctx,
	}
}

func (gc GeminiClient) GenerateAnswer(prompt string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	return gc.parseAnswer(resp)

}

func (gc GeminiClient) parseAnswer(resp *genai.GenerateContentResponse) string {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				return fmt.Sprint(part)
			}
		}
	}
	return "---"
}
