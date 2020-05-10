package main

import (
	"testing"
)

func TestScan(t *testing.T) {
	cases := []struct {
		dependencies     string
		language         string
		expectedExitCode int
	}{
		{"testdata/pip_requirements.txt", "Python", exitOK},
		{"testdata/pip_requirements.txt", "SomeRandomLanguage", exitLanguage},
		{"dontExist.txt", "Python", exitFileError},
		{"", "", exitMissingArguments},
		{"dontExist.txt", "Python", exitFileError},
	}

	for _, testCase := range cases {
		code, _ := StartScan(testCase.dependencies, testCase.language)
		if testCase.expectedExitCode != code {
			t.Error("Unexpected exit code")
		}
	}
}
