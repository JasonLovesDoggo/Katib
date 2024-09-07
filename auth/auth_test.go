package auth

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestConvertExcludedReposToArray(t *testing.T) {
	testCases := []struct {
		input         string
		expectedArray []string
		expectedError error
	}{
		// Empty and whitespace cases
		{"", []string{}, nil},
		{" ", []string{}, nil},
		{"  ,  ", nil, fmt.Errorf(InvalidRepoEntryError.Error(), " (empty entry, check commas)")}, // Only commas (all empty entries)

		// Single repository cases
		{"owner/repo", []string{"owner/repo"}, nil},
		{"owner/Repo", []string{"owner/Repo"}, nil},       // Case sensitivity (allowed by the regex)
		{"owner/repo ", []string{"owner/repo"}, nil},      // Trailing whitespace (trimmed)
		{" owner/repo", []string{"owner/repo"}, nil},      // Leading whitespace (trimmed)
		{"   owner/repo   ", []string{"owner/repo"}, nil}, // Leading and trailing whitespace (trimmed)

		// Multiple repository cases
		{"owner1/repo1,owner2/repo2", []string{"owner1/repo1", "owner2/repo2"}, nil},
		{"owner1/repo1, owner2/repo2", []string{"owner1/repo1", "owner2/repo2"}, nil}, // Whitespace around commas
		{"owner1/repo1 ,owner2/repo2 , owner3/repo3", []string{"owner1/repo1", "owner2/repo2", "owner3/repo3"}, nil},

		// Invalid character cases
		{"owner/r323>@!epo!", nil, InvalidCharError},
		{"ow29()@#ner/repo", nil, InvalidCharError},
		{"owner/repo,owner/invalid*repo", nil, InvalidCharError},

		// Invalid format cases
		{"/repo", nil, InvalidCharError},  // Leading slash with no owner (invalid char)
		{"owner/", nil, InvalidCharError}, // Missing repo name after slash
		{",owner/repo", nil, fmt.Errorf(InvalidRepoEntryError.Error(), " (empty entry, check commas)")}, // Leading comma results in an empty entry
		{"owner/repo,", nil, fmt.Errorf(InvalidRepoEntryError.Error(), " (empty entry, check commas)")}, // Trailing comma results in an empty entry
		{"/owner/repo", nil, InvalidCharError}, // Leading slash (invalid char)
		{"owner1/repo1,,owner2/repo2", nil, fmt.Errorf(InvalidRepoEntryError.Error(), " (empty entry, check commas)")}, // Empty entry (two commas)
		{",,,", nil, fmt.Errorf(InvalidRepoEntryError.Error(), " (empty entry, check commas)")},                        // Only commas (all empty entries)
		{"owner/subgroup/repo", nil, InvalidCharError},                                                                 // Multiple slashes in one entry
		{"valid/repo,invalid,another-valid/repo", nil, InvalidCharError},                                               // One invalid entry

		// Duplicate entry cases
		{"owner/repo,owner/repo,owner/other_repo", nil, fmt.Errorf(DuplicateRepoError.Error(), "owner/repo")}, // Duplicate entry

		// Edge cases
		{"owner/repo with spaces", nil, InvalidCharError}, // Spaces in repo name,
	}

	for _, testCase := range testCases {
		resultArray, resultError := convertExcludedReposToArray(testCase.input)

		// Check if the error matches the expectation
		if resultError != nil && testCase.expectedError != nil {
			if resultError.Error() != testCase.expectedError.Error() {
				t.Errorf("Input: `%s` Expected error: %v, Got error: %v", testCase.input, testCase.expectedError, resultError)
			}
		} else if !errors.Is(resultError, testCase.expectedError) { // One is nil, the other isn't
			t.Errorf("Input: `%s` Expected error: %v, Got error: %v", testCase.input, testCase.expectedError, resultError)
		}

		// If no error is expected, check if the array matches
		if testCase.expectedError == nil {
			if !reflect.DeepEqual(resultArray, testCase.expectedArray) {
				t.Errorf("Input: `%s` Expected array: %v, Got array: %v", testCase.input, testCase.expectedArray, resultArray)
			}
		}
	}
}
