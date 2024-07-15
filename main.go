package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
)

var Client *githubv4.Client

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	Client = githubv4.NewClient(httpClient)
	// Use client...
	r := gin.Default()
	r.GET("/commits/latest", LatestCommit)
	r.Run()
}
