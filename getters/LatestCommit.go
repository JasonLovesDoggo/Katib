package getters

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"slices"
	"time"
)

var excludedRepos = []string{"JasonLovesDoggo", "notes"} // List of repos to exclude from the search (constant)

type MostRecentCommit struct {
	Additions       int        `json:"additions"`
	Deletions       int        `json:"deletions"`
	CommitUrl       string     `json:"commitUrl"`
	CommittedDate   time.Time  `json:"committedDate"`
	Oid             string     `json:"oid"`
	MessageHeadline string     `json:"messageHeadline"`
	MessageBody     string     `json:"messageBody"`
	Languages       []Language `json:"languages"`
}

type Language struct {
	Size  int    `json:"size"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func GetMostRecentCommit(client *githubv4.Client) (MostRecentCommit, error) {
	var query struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name      string
					Languages struct {
						Edges []struct {
							Size int
							Node struct {
								Name  string
								Color string
							}
						}
					} `graphql:"languages(first: 5)"`
					DefaultBranchRef struct {
						Target struct {
							Commit struct {
								History struct {
									Edges []struct {
										Node struct {
											AbbreviatedOid  string
											Additions       int
											Deletions       int
											CommitUrl       githubv4.URI
											CommittedDate   time.Time
											MessageHeadline string
											MessageBody     string
										}
									} `graphql:"edges"`
								} `graphql:"history(first: 1)"` // Fetch only the most recent commit
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first: 20, privacy: PUBLIC)"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String("JasonLovesDoggo"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return MostRecentCommit{}, fmt.Errorf("GraphQL query error: %v", err)
	}

	mostRecentCommit := MostRecentCommit{}

	for _, repo := range query.User.Repositories.Nodes {
		if slices.Contains(excludedRepos, repo.Name) {
			continue // Skip excluded repositories
		}

		commit := repo.DefaultBranchRef.Target.Commit.History.Edges[0].Node

		if commit.Additions > 5 && commit.Deletions > 5 && commit.CommittedDate.After(mostRecentCommit.CommittedDate) {
			languages := make([]Language, len(repo.Languages.Edges))
			for i, languageEdge := range repo.Languages.Edges {
				languages[i] = Language{
					Size:  languageEdge.Size,
					Name:  languageEdge.Node.Name,
					Color: languageEdge.Node.Color,
				}
			}

			mostRecentCommit = MostRecentCommit{
				Additions:       commit.Additions,
				Deletions:       commit.Deletions,
				CommitUrl:       commit.CommitUrl.String(), // Convert to string
				CommittedDate:   commit.CommittedDate,
				Oid:             commit.AbbreviatedOid,
				MessageHeadline: commit.MessageHeadline,
				MessageBody:     commit.MessageBody,
				Languages:       languages,
			}
		}
	}

	return mostRecentCommit, nil
}
