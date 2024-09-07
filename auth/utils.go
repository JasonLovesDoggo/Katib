package auth

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var (
	InvalidRepoEntryError = errors.New("invalid format for repository entry %s")
	DuplicateRepoError    = errors.New("duplicate repository entry found: %s cannot be contained twice")
	InvalidCharError      = errors.New("invalid character found in repository entry. It should only contain alphanumeric characters, hyphens, and periods. No spaces. If the repository is in an organization, it should be in the format 'owner/repo'")
	// Regex to validate repository entries
	validRepoRegex, _ = regexp.Compile(`^[\w.-]+/[\w.-]+$`)
)

/*
Converts a comma-separated string of repositories to an array of strings. The string should be in the format 'owner/repo'.
Returns an error if the string is not well-formed or contains invalid characters.
*/
func convertExcludedReposToArray(excludedRepos string) ([]string, error) {
	// Handle the empty string case first
	excludedRepos = strings.TrimSpace(excludedRepos) // Remove **leading and trailing** whitespace
	if excludedRepos == "" {
		return []string{}, nil // No repositories specified, return an empty array
	}

	repos := strings.Split(excludedRepos, ",")
	validRepos := make([]string, 0, len(repos))

	for _, repo := range repos {
		repo = strings.TrimSpace(repo) // Remove leading and trailing whitespace
		if slices.Contains(validRepos, repo) {
			return nil, fmt.Errorf(DuplicateRepoError.Error(), repo)
		}
		if repo == "" {
			return nil, fmt.Errorf(InvalidRepoEntryError.Error(), repo+" (empty entry, check commas)")
		}
		// If there is whitespace in the repository name, it is invalid
		if strings.Contains(repo, " ") {
			return nil, InvalidCharError
		}

		if !validRepoRegex.MatchString(repo) {
			return nil, InvalidCharError
		}

		validRepos = append(validRepos, repo)
	}

	// Basic validation to ensure the string is well-formed

	return validRepos, nil // Return the array of repositories if no errors are found
}
