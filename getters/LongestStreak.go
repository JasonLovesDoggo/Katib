package getters

import (
	"context"
	"github.com/jasonlovesdoggo/katib/auth"
	"github.com/shurcooL/githubv4"
	"time"
)

var (
	NoTime = time.Time{}
    BlankStreak = Streak{NoTime, NoTime, 0, 0}
)

// Define the GraphQL query
type contributionsQuery struct {
	User struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				Weeks []struct {
					ContributionDays []struct {
						ContributionCount int
						Date              githubv4.Date
					}
				}
			}
		} `graphql:"contributionsCollection(from: $from, to: $to)"`
	} `graphql:"user(login: $login)"`
}

type Streak struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Days         int       `json:"days"`
	TotalCommits int       `json:"total_commits"`
}

func GetStreak(client *githubv4.Client, from time.Time, to time.Time) (Streak, error) {
	var query contributionsQuery
	variables := map[string]interface{}{
		"login": githubv4.String(auth.Username),
		"from":  githubv4.DateTime{Time: from},
		"to":    githubv4.DateTime{Time: to},
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return BlankStreak, err
	}

	streak := BlankStreak
	for _, week := range query.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			if day.ContributionCount > 0 {
				streak.Days++
				streak.TotalCommits += day.ContributionCount
			} else {
				return streak, nil // Streak broken todo: find largest and also store starting date somewhere
			}
		}
	}
	return streak, nil
}

func GetLifetimeStreak(client *githubv4.Client) (Streak, error) {
	return GetStreak(client, time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()) // It is presumed that you don't have a streak before 2008

}
