package llm

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

type GeminiConfig struct {
	ctx context.Context
}

type GeminiClient struct {
	model *genai.GenerativeModel
	ctx   context.Context
}
