package main

import (
	"flag"
	"fmt"
	"os"
)

var lokiVersion = "0.0.1"

// Language support: what's available on GitHub's Advisory DB as Ecosystems
var (
	supportedLanguages = map[string]string{
		// Language : PackageEcosystem
		"python":     "pip",
		"javascript": "npm",
		"dotnet":     "nuget",
		"java":       "maven",
		"php":        "composer",
	}
)

// Exit codes
const (
	exitOK = iota
	exitMissingArguments
	exitFileError
	exitLanguage
	exitMissingApiKey
)

// Startup routine
func init() {
	fmt.Print("Welcome to Loki the Open Source Security Scanner!\n\n")
	flag.Usage = help
}

// Help menu
func help() {
	h := "Scans open-source dependencies for known security advisories\n\n"

	h += "Usage:\n"
	h += "  loki [OPTIONS]\n\n"

	h += "Options:\n"
	h += "  -l, --language	Language to be scanned. Available: "
	for lang, _ := range supportedLanguages {
		h += lang + ";"
	}
	h += "\n"
	h += "  -d, --dependencies	Path to dependencies file\n"

	h += "Examples:\n"
	h += "  loki --language Python -dependencies requirements.txt\n"
	h += "  loki --language JavaScript -dependencies package-lock.json\n"

	h += "Exit codes:\n"
	h += fmt.Sprintf("  %d\t%s\n", exitOK, "OK")
	h += fmt.Sprintf("  %d\t%s\n", exitMissingArguments, "Missing arguments")
	h += fmt.Sprintf("  %d\t%s\n", exitFileError, "File not found")
	h += fmt.Sprintf("  %d\t%s\n", exitLanguage, "Language not available")

	fmt.Fprintf(os.Stderr, h)
}

func main() {
	var (
		languageFlag     string
		dependenciesFlag string
		versionFlag      bool
	)

	// Arguments
	flag.StringVar(&languageFlag, "language", "", "Language being scanned.")
	flag.StringVar(&languageFlag, "l", "", "Language being scanned.")
	flag.StringVar(&dependenciesFlag, "dependencies", "", "Path to your dependencies file")
	flag.StringVar(&dependenciesFlag, "d", "", "Path to your dependencies file")
	flag.BoolVar(&versionFlag, "version", false, "Version of this library")
	flag.Parse()

	if versionFlag {
		fmt.Printf("loki version %s\n", lokiVersion)
		os.Exit(exitOK)
	}

	// Scan info
	fmt.Println("Scan information:")
	fmt.Println("Dependencies file:", dependenciesFlag)
	fmt.Println("Language:", languageFlag)
	fmt.Println("To supress and justify (optionally) findings, simply provide a gos3ignore.yml file.")

	// Call scan routine
	code, err := StartScan(dependenciesFlag, languageFlag)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
