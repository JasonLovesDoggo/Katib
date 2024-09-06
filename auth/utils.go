package auth

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var (
	InvalidEndCharError   = errors.New("invalid format for excluded repositories")
	InvalidRepoEntryError = errors.New("invalid format for repository entry %s")
	DuplicateRepoError    = errors.New("duplicate repository entry found: %s and %s")
	InvalidCharError      = errors.New("invalid character found in repository entry. It should only contain alphanumeric characters, hyphens, and periods. If the repository is in an organization, it should be in the format 'owner/repo'")
)

var validRepoRegex, _ = regexp.Compile("(^[\\w.-]+$\n)/(^[\\w.-]+$)") // see https://stackoverflow.com/a/59082561/18516611

func convertExcludedReposToArray(excludedRepos string) ([]string, error) {
	// Handle the empty string case first
	if excludedRepos == "" {
		return []string{}, nil // No repositories specified, return an empty array
	}

	// Basic validation to ensure the string is well-formed
	if strings.HasPrefix(excludedRepos, ",") || strings.HasSuffix(excludedRepos, ",") || strings.HasPrefix(excludedRepos, "/") {
		return nil, InvalidEndCharError
	}

	// Split the string into individual repository entries
	repos := strings.Split(excludedRepos, ",")

	var validRepos []string

	// Further validation and conversion
	for _, repo := range repos {
		// Check for missing owner or repository name
		repo = strings.TrimSuffix(repo, "/") // Remove any trailing slash

		// Check if each entry has exactly one slash
		if strings.Count(repo, "/") != 1 {
			return nil, fmt.Errorf(InvalidRepoEntryError.Error(), repo)
		}

		// Trim any leading or trailing whitespace
		repo = strings.TrimSpace(repo)
		// Check for duplicate entries
		if slices.Contains(validRepos, repo) {
			return nil, fmt.Errorf(DuplicateRepoError.Error(), repo, repo)
		}

		// Check for invalid characters
		if !validRepoRegex.MatchString(repo) {
			return nil, InvalidCharError
		}

		validRepos = append(validRepos, repo)
	}

	return repos, nil // Return the array of repositories
}
