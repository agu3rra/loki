package main

import (
	"errors"
	"fmt"
)

func StartScan(dependencies string, language string) (int, error) {
	fmt.Println(dependencies)
	fmt.Println(language)

	// Required inputs check
	if dependencies == "" || language == "" {
		fmt.Println("Missing required flags. See --help below:")
		help()
		return exitMissingArguments, errors.New("Missing arguments")
	}

	return exitOK, nil
}
