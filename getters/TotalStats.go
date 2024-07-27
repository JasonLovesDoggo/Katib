package getters

import (
	"context"
	"fmt"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/shurcooL/githubv4"
)

/*
	todo: remove

// GraphQL query to fetch the total commit count for each repository

	query ($username: String!, $cursor: String) {
	  rateLimit {
	    cost
	  }
	  user(login: $username) {
	    repositoriesContributedTo(first: 100, after: $cursor, includeUserRepositories: true, contributionTypes: [COMMIT, ISSUE, PULL_REQUEST, REPOSITORY]) {
	      totalCount
	      pageInfo {
	        endCursor
	        hasNextPage
	      }
	      nodes {
	        nameWithOwner
	        defaultBranchRef {
	          target {
	            ... on Commit {
	              history {
	                totalCount # Get total commit count for the repository
	              }
	            }
	          }
	        }
	      }
	    }
	  }
	}
*/
type Repository struct {
	NameWithOwner string
}

type Repositories struct {
	TotalCount int
	PageInfo   struct {
		EndCursor   string
		HasNextPage bool
	}
	Nodes []Repository
}

type ContributedTo struct {
	RepositoriesContributedTo Repositories
}

type User struct {
	ContributedTo ContributedTo
}

type StatsData struct {
	TotalCommits      uint32
	TotalAdditions    uint32
	TotalDeletions    uint32
	TotalRepositories uint16
}

func GetTotalStats(client *githubv4.Client) (StatsData, error) {
	var query struct {
		User struct {
			Repositories struct {
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
				Nodes []struct {
					NameWithOwner    string
					DefaultBranchRef struct {
						Target struct {
							Commit struct {
								History struct {
									Nodes []struct {
										Author struct {
											User struct {
												Login githubv4.String
											}
										}
										Additions int
										Deletions int
									} `graphql:"nodes"`
									TotalCount int
								} `graphql:"history(first: 100)"` // Fetch up to 100 commits per page
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first: 100, privacy: PUBLIC, after: $cursor)"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String(auth.USERNAME),
		"cursor":   (*githubv4.String)(nil), // Start from the beginning
	}

	var totalCommits, totalAdditions, totalDeletions uint32
	var totalRepositories uint16

	for {
		err := client.Query(context.Background(), &query, variables)
		if err != nil {
			return StatsData{}, fmt.Errorf("GraphQL query error: %v", err)
		}

		for _, repo := range query.User.Repositories.Nodes {
			//if slices.Contains(excludedRepos, repo.NameWithOwner){
			//continue // Skip excluded repositories

			for _, commit := range repo.DefaultBranchRef.Target.Commit.History.Nodes {
				if commit.Author.User.Login == auth.USERNAME { // Filter by the target user
					totalCommits++
					totalAdditions += uint32(commit.Additions)
					totalDeletions += uint32(commit.Deletions)
				}
			}
		}

		if !query.User.Repositories.PageInfo.HasNextPage {
			break // No more pages to fetch
		}
		variables["cursor"] = query.User.Repositories.PageInfo.EndCursor
	}

	return StatsData{totalCommits, totalAdditions, totalDeletions, totalRepositories}, nil
}
