package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	API_KEY = os.Getenv("GITHUB_APIKEY")
	API_URL = "https://api.github.com/graphql"
)

// Retrieves advisories from Github's Advisory DB
// Important: set GITHUB_APIKEY as an OS variable before using it
func GetAdvisories(library string, ecosystem string) (string, error) {
	if API_KEY == "" {
		return "", errors.New("Missing Github API KEY")
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", API_URL, nil)

	if err != nil {
		return "", err
	}

	token := fmt.Sprintf("bearer %s", API_KEY)
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Sprintf("Response status code: %s", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Print(body)
	return "The advisory data goes here", nil
}
