package getters

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"slices"
	"time"
)

var excludedRepos = []string{"JasonLovesDoggo/JasonLovesDoggo", "JasonLovesDoggo/notes", "JasonLovesDoggo/status"} // List of repos to exclude from the search (constant)

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

type CommitItem struct {
	Repo            string    `json:"repo"`
	Additions       int       `json:"additions"`
	Deletions       int       `json:"deletions"`
	CommitUrl       string    `json:"commitUrl"`
	CommittedDate   time.Time `json:"committedDate"`
	Oid             string    `json:"oid"`
	MessageHeadline string    `json:"messageHeadline"`
	MessageBody     string    `json:"messageBody"`
}

type CommitsListResponse struct {
	Commits   []CommitItem `json:"commits"`
	Languages []Language   `json:"languages"`
	Stats     struct {
		TotalAdditions int `json:"totalAdditions"`
		TotalDeletions int `json:"totalDeletions"`
		TotalCommits   int `json:"totalCommits"`
	} `json:"stats"`
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
											Author          struct {
												User struct {
													Login string
												}
											}
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
		"username": githubv4.String("JasonLovesDoggo"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return MostRecentCommit{}, fmt.Errorf("GraphQL query error: %v", err)
	}

	mostRecentCommit := MostRecentCommit{MessageHeadline: "Something went wrong", MessageBody: "Please try again later", Languages: []Language{}, CommittedDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)} // Set to a date in the past

	for _, repo := range query.User.Repositories.Nodes {
		if slices.Contains(excludedRepos, repo.NameWithOwner) {
			continue // Skip excluded repositories
		}

		for _, edge := range repo.DefaultBranchRef.Target.Commit.History.Edges {
			commit := edge.Node

			// Skip commits not authored by JasonLovesDoggo
			if commit.Author.User.Login != "JasonLovesDoggo" {
				continue
			}

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

func GetCommitsList(client *githubv4.Client, limit int) (CommitsListResponse, error) {
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
					} `graphql:"languages(first: 10)"`
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
											Author          struct {
												User struct {
													Login string
												}
											}
										}
									} `graphql:"edges"`
								} `graphql:"history(first: 10)"`
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first: 20, privacy: PUBLIC, orderBy: {field: UPDATED_AT, direction: DESC})"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String("JasonLovesDoggo"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return CommitsListResponse{}, fmt.Errorf("GraphQL query error: %v", err)
	}

	var allCommits []CommitItem
	languageMap := make(map[string]Language)
	totalAdditions := 0
	totalDeletions := 0

	for _, repo := range query.User.Repositories.Nodes {
		if slices.Contains(excludedRepos, repo.NameWithOwner) {
			continue
		}

		// Collect language data from active repos
		hasValidCommit := false
		for _, edge := range repo.DefaultBranchRef.Target.Commit.History.Edges {
			if edge.Node.Author.User.Login == "JasonLovesDoggo" && edge.Node.Additions > 5 {
				hasValidCommit = true
				break
			}
		}

		if hasValidCommit {
			for _, langEdge := range repo.Languages.Edges {
				key := langEdge.Node.Name
				if existing, ok := languageMap[key]; ok {
					languageMap[key] = Language{
						Size:  existing.Size + langEdge.Size,
						Name:  langEdge.Node.Name,
						Color: langEdge.Node.Color,
					}
				} else {
					languageMap[key] = Language{
						Size:  langEdge.Size,
						Name:  langEdge.Node.Name,
						Color: langEdge.Node.Color,
					}
				}
			}
		}

		for _, edge := range repo.DefaultBranchRef.Target.Commit.History.Edges {
			commit := edge.Node

			// Skip commits not authored by JasonLovesDoggo
			if commit.Author.User.Login != "JasonLovesDoggo" {
				continue
			}

			// Skip tiny commits
			if commit.Additions <= 5 {
				continue
			}

			totalAdditions += commit.Additions
			totalDeletions += commit.Deletions

			allCommits = append(allCommits, CommitItem{
				Repo:            repo.NameWithOwner,
				Additions:       commit.Additions,
				Deletions:       commit.Deletions,
				CommitUrl:       commit.CommitUrl.String(),
				CommittedDate:   commit.CommittedDate,
				Oid:             commit.AbbreviatedOid,
				MessageHeadline: commit.MessageHeadline,
				MessageBody:     commit.MessageBody,
			})
		}
	}

	// Sort by date descending
	for i := 0; i < len(allCommits)-1; i++ {
		for j := i + 1; j < len(allCommits); j++ {
			if allCommits[j].CommittedDate.After(allCommits[i].CommittedDate) {
				allCommits[i], allCommits[j] = allCommits[j], allCommits[i]
			}
		}
	}

	// Convert language map to sorted slice
	var languages []Language
	for _, lang := range languageMap {
		languages = append(languages, lang)
	}
	// Sort languages by size descending
	for i := 0; i < len(languages)-1; i++ {
		for j := i + 1; j < len(languages); j++ {
			if languages[j].Size > languages[i].Size {
				languages[i], languages[j] = languages[j], languages[i]
			}
		}
	}

	// Apply limit
	limitedCommits := allCommits
	if limit > 0 && limit < len(allCommits) {
		limitedCommits = allCommits[:limit]
	}

	response := CommitsListResponse{
		Commits:   limitedCommits,
		Languages: languages,
	}
	response.Stats.TotalAdditions = totalAdditions
	response.Stats.TotalDeletions = totalDeletions
	response.Stats.TotalCommits = len(allCommits)

	return response, nil
}
