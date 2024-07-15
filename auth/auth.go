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

var Client *githubv4.Client

func init() {
	err := godotenv.Load()
	if err != nil && os.Getenv("GITHUB_TOKEN") == "" { // If the .env file is not found and the GITHUB_TOKEN is not set
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Initializing client...")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	Client = githubv4.NewClient(httpClient)
}
