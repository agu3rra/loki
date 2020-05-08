package main

import (
	"errors"
	"testing"
)

func TestLoki(t *testing.T) {
	cases := []struct {
		dependencies     string
		language         string
		expectedExitCode int
		expectedError    error
	}{
		{"requirements.txt", "Python", exitOK, nil},
		{"dontExist.txt", "Python", exitOK, nil},
		{"", "", exitMissingArguments, errors.New("Missing arguments")},
		{"dontExist.txt", "Python", exitFileNotFound, errors.New("File not found")},
		{"requirements.txt", "SomeRandomLanguage", exitLanguage, errors.New("Language not available")},
	}

	for _, testCase := range cases {
		code, err := StartScan(testCase.dependencies, testCase.language)
		if testCase.expectedExitCode != code {
			t.Error("Unexpected exit code")
		}
		if testCase.expectedError != err {
			t.Error("Unexpected error")
		}
	}
}
