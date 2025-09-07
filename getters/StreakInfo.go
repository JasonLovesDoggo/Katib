package getters

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"time"
)

type StreakInfo struct {
	CurrentStreak int  `json:"currentStreak"`
	HighestStreak int  `json:"highestStreak"`
	Active        bool `json:"active"`
}

func GetStreakInfo(client *githubv4.Client) (StreakInfo, error) {
	var query struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					Weeks []struct {
						ContributionDays []struct {
							ContributionCount int
							Date              string
						}
					}
				}
			}
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String("JasonLovesDoggo"),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return StreakInfo{}, fmt.Errorf("GraphQL query error: %v", err)
	}

	var contributionDays []struct {
		ContributionCount int
		Date              string
	}

	for _, week := range query.User.ContributionsCollection.ContributionCalendar.Weeks {
		contributionDays = append(contributionDays, week.ContributionDays...)
	}

	currentStreak := calculateCurrentStreak(contributionDays)
	highestStreak := calculateHighestStreak(contributionDays)
	active := isActive(contributionDays)

	return StreakInfo{
		CurrentStreak: currentStreak,
		HighestStreak: highestStreak,
		Active:        active,
	}, nil
}

func calculateCurrentStreak(contributionDays []struct {
	ContributionCount int
	Date              string
}) int {
	if len(contributionDays) == 0 {
		return 0
	}

	now := time.Now()
	today := now.Format("2006-01-02")
	yesterday := now.AddDate(0, 0, -1).Format("2006-01-02")

	// Check if there's activity today or yesterday
	hasRecentActivity := false
	for i := len(contributionDays) - 1; i >= 0; i-- {
		day := contributionDays[i]
		if (day.Date == today || day.Date == yesterday) && day.ContributionCount > 0 {
			hasRecentActivity = true
			break
		}
	}

	// If no recent activity, streak is 0
	if !hasRecentActivity {
		return 0
	}

	// Count consecutive days with contributions from the end
	currentStreak := 0
	skippedToday := false

	for i := len(contributionDays) - 1; i >= 0; i-- {
		day := contributionDays[i]
		dayDate, err := time.Parse("2006-01-02", day.Date)
		if err != nil {
			continue
		}

		// Skip future dates
		if dayDate.After(now) {
			continue
		}

		// If it's today and no contributions, skip it once
		if day.Date == today && day.ContributionCount == 0 && !skippedToday {
			skippedToday = true
			continue
		}

		if day.ContributionCount > 0 {
			currentStreak++
		} else {
			break
		}
	}

	return currentStreak
}

func calculateHighestStreak(contributionDays []struct {
	ContributionCount int
	Date              string
}) int {
	if len(contributionDays) == 0 {
		return 0
	}

	highestStreak := 0
	currentStreak := 0

	for _, day := range contributionDays {
		if day.ContributionCount > 0 {
			currentStreak++
			if currentStreak > highestStreak {
				highestStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}
	}

	return highestStreak
}

func isActive(contributionDays []struct {
	ContributionCount int
	Date              string
}) bool {
	if len(contributionDays) == 0 {
		return false
	}

	now := time.Now()
	today := now.Format("2006-01-02")
	yesterday := now.AddDate(0, 0, -1).Format("2006-01-02")

	for i := len(contributionDays) - 1; i >= 0; i-- {
		day := contributionDays[i]
		if (day.Date == today || day.Date == yesterday) && day.ContributionCount > 0 {
			return true
		}
	}

	return false
}
