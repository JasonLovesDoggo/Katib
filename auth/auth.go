package auth

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

var (
	Client         *githubv4.Client
	Username       string
	ExcludedRepos  []string
	RequiredTokens = [...]string{"GITHUB_TOKEN", "USERNAME"}
)

func init() {
	err := godotenv.Load("../.env", ".env")
	if err != nil { // If the .env file is not found
		log.Fatalf("Error loading .env file: %v", err)
	}
	for _, token := range RequiredTokens {
		if os.Getenv(token) == "" {
			log.Fatalf("Error: %s is not set. Please set it in your .env file", token)
		}
	}

	Username = os.Getenv("USERNAME")

	fmt.Println("Initializing client...")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	Client = githubv4.NewClient(httpClient)
}
