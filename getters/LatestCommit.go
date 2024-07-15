package getters

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"slices"
	"time"
)

var excludedRepos = []string{"JasonLovesDoggo/JasonLovesDoggo", "JasonLovesDoggo/notes"} // List of repos to exclude from the search (constant)

func GetMostRecentCommit(client *githubv4.Client) {
	var query struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name             string
					DefaultBranchRef struct {
						Target struct {
							Commit struct {
								History struct {
									Edges []struct {
										Node struct {
											Oid             githubv4.GitObjectID
											MessageHeadline string
											URL             githubv4.URI
											CommittedDate   time.Time
										}
									} `graphql:"edges"`
								} `graphql:"history(first:1)"`
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first:100,privacy:PUBLIC)"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String("JasonLovesDoggo"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err)
		return
	}

	var mostRecentCommit struct {
		Oid             githubv4.GitObjectID
		MessageHeadline string
		URL             githubv4.URI
		CommittedDate   time.Time
	}
	mostRecentCommitTime := time.Time{} // Initialize to zero time

	for _, repo := range query.User.Repositories.Nodes {
		if !slices.Contains(excludedRepos, repo.Name) {
			commit := repo.DefaultBranchRef.Target.Commit.History.Edges[0].Node
			if commit.CommittedDate.After(mostRecentCommitTime) {
				mostRecentCommit = commit
				mostRecentCommitTime = commit.CommittedDate
			}
		}
	}

	fmt.Println("Most Recent Commit:")
	fmt.Println("OID:", mostRecentCommit.Oid)
	fmt.Println("Message:", mostRecentCommit.MessageHeadline)
	fmt.Println("URL:", mostRecentCommit.URL)
	fmt.Println("Committed Date:", mostRecentCommit.CommittedDate)
	// todo: returning
}
