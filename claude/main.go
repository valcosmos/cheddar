package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic"
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic/client/native"
)

func envVariable(key string) string {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}

	return value
}
func main() {
	apiKey := envVariable("API_KEY")
	url := envVariable("URL")
	fmt.Println(apiKey, url)

	ctx := context.Background()

	client, err := native.MakeClient(native.Config{
		APIKey:  apiKey,
		BaseURL: url,
	})
	if err != nil {
		panic(err)
	}

	request := anthropic.NewMessageRequest(
		[]anthropic.MessagePartRequest{{Role: "user",
			Content: []anthropic.ContentBlock{anthropic.NewTextContentBlock("Hello, world!")}}},
		anthropic.WithModel[anthropic.MessageRequest](anthropic.Claude35Sonnet),
		anthropic.WithMaxTokens[anthropic.MessageRequest](20),
	)

	response, err := client.Message(ctx, request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Content)

}
