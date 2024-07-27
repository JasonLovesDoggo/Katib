package getters

import (
	"context"
	"fmt"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/shurcooL/githubv4"
	"slices"
	"time"
)

type MostRecentCommit struct {
	Repo            string         `json:"repo"`
	Additions       int            `json:"additions"`
	Deletions       int            `json:"deletions"`
	CommitUrl       string         `json:"commitUrl"`
	CommittedDate   time.Time      `json:"committedDate"`
	Oid             string         `json:"oid"`
	MessageHeadline string         `json:"messageHeadline"`
	MessageBody     string         `json:"messageBody"`
	Languages       []Language     `json:"languages"`
	ParentCommits   []parentCommit `json:"parentCommits"`
}

type parentCommit struct {
	Additions       int       `json:"additions"`
	Deletions       int       `json:"deletions"`
	CommitUrl       string    `json:"commitUrl"`
	CommittedDate   time.Time `json:"committedDate"`
	MessageHeadline string    `json:"messageHeadline"`
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
					NameWithOwner string
					Languages     struct {
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
								} `graphql:"history(first: 5)"`
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first: 10, privacy: PUBLIC, orderBy: {field: UPDATED_AT, direction: DESC})"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String(auth.USERNAME),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return MostRecentCommit{}, fmt.Errorf("GraphQL query error: %v", err)
	}

	mostRecentCommit := MostRecentCommit{MessageHeadline: "Something went wrong", MessageBody: "Please try again later", Languages: []Language{}, CommittedDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)} // Set to a date in the past

	for _, repo := range query.User.Repositories.Nodes {
		if slices.Contains(auth.ExcludedRepos, repo.NameWithOwner) {
			continue // Skip excluded repositories
		}

		for _, edge := range repo.DefaultBranchRef.Target.Commit.History.Edges {
			commit := edge.Node
			if repo.NameWithOwner == mostRecentCommit.Repo && commit.CommittedDate.Before(mostRecentCommit.CommittedDate) {
				mostRecentCommit.ParentCommits = append(mostRecentCommit.ParentCommits, parentCommit{
					Additions:       commit.Additions,
					Deletions:       commit.Deletions,
					CommitUrl:       commit.CommitUrl.String(), // Convert to string
					CommittedDate:   commit.CommittedDate,
					MessageHeadline: commit.MessageHeadline,
				})
			}

			if commit.Additions > 5 && commit.CommittedDate.After(mostRecentCommit.CommittedDate) { // Get the most recent non-bs commit
				fmt.Print(" - Replacing commit: " + mostRecentCommit.CommittedDate.String())
				languages := make([]Language, len(repo.Languages.Edges))
				for i, languageEdge := range repo.Languages.Edges {
					languages[i] = Language{
						Size:  languageEdge.Size,
						Name:  languageEdge.Node.Name,
						Color: languageEdge.Node.Color,
					}
				}
				mostRecentCommit = MostRecentCommit{
					Repo:            repo.NameWithOwner,
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
	}

	return mostRecentCommit, nil
}
