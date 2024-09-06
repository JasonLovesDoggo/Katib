package auth

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var (
	Client *githubv4.Client
	Username string
	ExcludedRepos []string
	RequiredTokens = [...]string{"GITHUB_TOKEN", "USERNAME"}
)



func init() {
	err := godotenv.Load()
	if err != nil { // If the .env file is not found
		log.Fatal("Error loading .env file")
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


func validateExcludedRepos(excludedRepos string) bool{
	// ExcludedRepos is a comma separated list of repositories to exclude
	// from the stats. This is useful for excluding forks or other repositories
	// that you don't want to include in your stats (such as automation repositories).

	if excludedRepos == "" {
		return true // Valid
	}

	if (strings.Count(excludedRepos, "/") != strings.Count(excludedRepos, ",") - 1) {
		return false // Invalid
	}
	return true // Valid
}
