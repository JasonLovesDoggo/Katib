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
	Additions       int
	Deletions       int
	CommitUrl       githubv4.URI
	CommittedDate   time.Time
	MessageHeadline string
	MessageBody     string
	URL             githubv4.URI
	Languages       []Language
	AbbreviatedOid  string
}

type Language struct {
	Size int
	Node struct {
		Name  string
		Color string
	}
}

func GetMostRecentCommit(client *githubv4.Client) (MostRecentCommit, error) {
	// language=GraphQL
	var query struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name           string
					StargazerCount int
					Languages      struct {
						Edges []struct { // Add Edges array
							Size int
							Node struct { // Node struct inside Edges
								Name  string
								Color string
							}
						}
					} `graphql:"languages(first:5)"`
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
								} `graphql:"history(first:5)"`
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first:20,privacy:PUBLIC)"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String("JasonLovesDoggo"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err)
		return MostRecentCommit{}, err
	}

	//mostRecentCommitTime := time.Time{} // Initialize to zero time
	mostRecentCommit := MostRecentCommit{}
	fmt.Println("Repositories:")
	for _, repo := range query.User.Repositories.Nodes {
		if !slices.Contains(excludedRepos, repo.Name) {
			for _, commit := range repo.DefaultBranchRef.Target.Commit.History.Edges {
				if commit.Node.Additions > 5 && commit.Node.Deletions > 5 {
					if commit.Node.CommittedDate.After(mostRecentCommit.CommittedDate) { // Stack approach
						var languages []Language
						fmt.Println("Languages", repo.Languages)
						if repo.Languages.Edges != nil {
							languages := make([]Language, len(repo.Languages.Edges))
							for i, languageEdge := range repo.Languages.Edges {
								languages[i] = Language{
									Size: languageEdge.Size,
									Node: languageEdge.Node,
								}
							}
						}

						fmt.Println("Commit:", commit.Node)
						mostRecentCommit = MostRecentCommit{
							CommittedDate:   commit.Node.CommittedDate,
							Additions:       commit.Node.Additions,
							Deletions:       commit.Node.Deletions,
							CommitUrl:       commit.Node.CommitUrl,
							MessageHeadline: commit.Node.MessageHeadline,
							MessageBody:     commit.Node.MessageBody,
							Languages:       languages,
							AbbreviatedOid:  commit.Node.AbbreviatedOid,
						}
					}
				}
			}
		}
	}

	return mostRecentCommit, nil
}
