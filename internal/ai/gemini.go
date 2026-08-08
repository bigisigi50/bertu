package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GenerateContent uses the Gemini API to generate a response for a given prompt.
func GenerateContent(ctx context.Context, prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY environment variable is not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	// Use the recommended model for general text tasks
	model := client.GenerativeModel("gemini-1.5-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response received from Gemini")
	}

	// Extract the text from the response
	var result string
	for _, part := range resp.Candidates[0].Content.Parts {
		if text, ok := part.(genai.Text); ok {
			result += string(text)
		}
	}

	return result, nil
}

// GeminiEmbeddingFunc returns a chromem.EmbeddingFunc using Gemini text-embedding-004
func GeminiEmbeddingFunc() func(ctx context.Context, text string) ([]float32, error) {
	return func(ctx context.Context, text string) ([]float32, error) {
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("GEMINI_API_KEY not set")
		}
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			return nil, err
		}
		defer client.Close()

		em := client.EmbeddingModel("text-embedding-004")
		res, err := em.EmbedContent(ctx, genai.Text(text))
		if err != nil {
			return nil, err
		}
		if res == nil || res.Embedding == nil || len(res.Embedding.Values) == 0 {
			return nil, fmt.Errorf("no embedding returned")
		}
		return res.Embedding.Values, nil
	}
}
