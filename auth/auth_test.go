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
		{"/repo", nil, InvalidEndCharError},                                                                  // Leading slash with no owner
		{"owner/", nil, fmt.Errorf(InvalidRepoEntryError.Error(), "owner")},                                  // Missing repo name after slash (trailing slash is handled)
		{",owner/repo", nil, InvalidEndCharError},                                                            // Leading comma
		{"owner/repo,", nil, InvalidEndCharError},                                                            // Trailing comma
		{"/owner/repo", nil, InvalidEndCharError},                                                            // Leading slash
		{"owner1/repo1,,owner2/repo2", nil, InvalidRepoEntryError},                                           // Empty entry (two commas)
		{",,,", nil, InvalidEndCharError},                                                                    // Only commas
		{"owner/subgroup/repo", nil, fmt.Errorf(InvalidRepoEntryError.Error(), "owner/subgroup/repo")},       // Multiple slashes in one entry
		{"valid/repo,invalid,another-valid/repo", nil, fmt.Errorf(InvalidRepoEntryError.Error(), "invalid")}, // One invalid entry
		{"owner/repo,owner/repo", nil, DuplicateRepoError},                                                   // Duplicate entries

		// Invalid Char cases
		{"owner/r323>@!epo!", nil, InvalidCharError},
		{"ow29()@#ner/repo", nil, InvalidCharError},

		// Valid cases
		{"owner/repo", []string{"owner/repo"}, nil},
		{"owner/repo,owner/repo2", []string{"owner/repo", "owner/repo2"}, nil},
		{"owner/repo,owner2/repo", []string{"owner/repo", "owner2/repo"}, nil},
	}

	for _, testCase := range testCases {
		resultArray, resultError := convertExcludedReposToArray(testCase.input)

		// Check if the error matches the expectation
		if resultError != nil && testCase.expectedError != nil {
			if resultError.Error() != testCase.expectedError.Error() {
				t.Errorf("Input: %s, Expected error: %v, Got error: %v", testCase.input, testCase.expectedError, resultError)
			}
		} else if !errors.Is(resultError, testCase.expectedError) { // One is nil, the other isn't
			t.Errorf("Input: %s, Expected error: %v, Got error: %v", testCase.input, testCase.expectedError, resultError)
		}

		// If no error is expected, check if the array matches
		if testCase.expectedError == nil {
			if !reflect.DeepEqual(resultArray, testCase.expectedArray) {
				t.Errorf("Input: %s, Expected array: %v, Got array: %v", testCase.input, testCase.expectedArray, resultArray)
			}
		}
	}
}
