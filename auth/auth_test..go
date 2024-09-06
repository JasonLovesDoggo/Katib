package auth

import "testing"

func TestValidateExcludedRepos(t *testing.T) {
	// Test case where excludedRepos is empty
	if !validateExcludedRepos("") {
		t.Errorf("Expected validateExcludedRepos to return true for empty string")
	}

	// Test case where excludedRepos is not empty
	if !validateExcludedRepos("repo1/two,repo2/hi") {
		t.Errorf("Expected validateExcludedRepos to return true for valid string")
	}

	// Test case where excludedRepos is invalid
	if validateExcludedRepos("repo1/two,repo2/hi,") {
		t.Errorf("Expected validateExcludedRepos to return false for trailing comma")
	}

	// Test case that are just repo name not repo/owner
	if validateExcludedRepos("repo1,repo2") {
		t.Errorf("Expected validateExcludedRepos to return false for missing owner or comma instead of /")
	}
}
