package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Loki the Open Source Security Scanner!")

	// Arguments
	dependenciesPtr := flag.String("dependencies", "", "Path to your dependencies file.")
	languagePtr := flag.String("language", "", "JavaScript, Python, Ruby, Java, PHP, DOTNET")
	flag.Parse()

	// Startup messages
	fmt.Println("Dependencies file:", *dependenciesPtr)
	fmt.Println("Language:", *languagePtr)
	fmt.Println("To supress and justify (optionally) findings, simply provide a gos3ignore.yml file.")

}
