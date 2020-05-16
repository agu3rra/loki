package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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
		return exitFileError, errors.New("Unable to open dependencies file")
	}
	defer file.Close()

	// Verify language availability
	language = strings.ToLower(language)
	ecosystem := supportedLanguages[language]
	if ecosystem == "" {
		msg := "Language not available"
		fmt.Println(msg)
		return exitLanguage, errors.New(msg)
	}

	// Parse dependencies file
	data, err := ioutil.ReadFile(dependencies)
	if err != nil {
		msg := "Error while reading dependencies"
		fmt.Println(msg)
		return exitFileError, errors.New(msg)
	}
	dependencies := string(data)
	strings.Split(dependencies, "\n")
	fmt.Println(string(data))

	return exitOK, nil
}
