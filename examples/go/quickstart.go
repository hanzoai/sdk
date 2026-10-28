package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hanzoai/go-sdk"
)

func main() {
	apiKey := os.Getenv("HANZO_API_KEY")
	if apiKey == "" {
		apiKey = "sk-test-key"
	}

	client := hanzo.NewClient(hanzo.WithAPIKey(apiKey))
	ctx := context.Background()

	// 1. List Models
	fmt.Println("--- 1. List Models ---")
	models, err := client.Models.List(ctx)
	if err == nil && len(models.Data) > 0 {
		fmt.Printf("Model count: %d, first: %s\n", len(models.Data), models.Data[0].ID)
	}

	// 2. Chat Completion
	fmt.Println("\n--- 2. Chat Completion ---")
	resp, err := client.Chat.CreateCompletion(ctx, hanzo.ChatCompletionRequest{
		Model: "zen5",
		Messages: []hanzo.ChatMessage{
			{Role: "user", Content: "Hello from Go!"},
		},
	})
	if err == nil && len(resp.Choices) > 0 {
		fmt.Printf("Reply: %s\n", resp.Choices[0].Message.Content)
	}

	// 3. Kai Decision
	fmt.Println("\n--- 3. Kai Decision Evaluation ---")
	decision, err := client.Decisions.Evaluate(ctx, hanzo.KaiDecisionRequest{
		AgentID:  "agent-go-01",
		PolicyID: "pol-safety-core",
		Action:   map[string]any{"type": "exec", "payload": map[string]any{"cmd": "status"}},
	})
	if err == nil {
		fmt.Printf("Decision: %s (Score: %.2f)\n", decision.Verdict, decision.Confidence)
	}
}
