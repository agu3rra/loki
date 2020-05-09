package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func StartScan(dependencies string, language string) (int, error) {
	// Required inputs check
	if dependencies == "" || language == "" {
		fmt.Println("Missing required flags. See --help below:")
		help()
		return exitMissingArguments, errors.New("Missing arguments")
	}

	// Reads provided dependencies file
	file, err := os.Open(dependencies)
	if err != nil {
		fmt.Println(err)
		os.Exit(exitFileNotFound)
	}

	// Verify language availability
	language = strings.ToLower(language)
	ecosystem := supportedLanguages[language]
	if ecosystem == "" {
		fmt.Println("Language not available")
		os.Exit(exitLanguage)
	}

	// Parse dependencies file
	fmt.Println(file)

	return exitOK, nil
}
