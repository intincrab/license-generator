package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	licenseName := flag.String("license", "", "License name")
	skipPrompt := flag.Bool("skip-prompt", false, "Skip prompts and use defaults")
	flag.Parse()

	licenses := fetchLicenses()
	validLicenseNames := getLicenseKeys(licenses)

	var licenseContent LicenseContent
	if *licenseName != "" && contains(validLicenseNames, strings.ToLower(*licenseName)) {
		licenseContent = getLicenseFromKey(licenses, strings.ToLower(*licenseName))
	} else {
		license := select_(getLicenseNames(licenses))
		licenseContent = getLicenseFromName(licenses, license)
	}

	fillContent(&licenseContent, *skipPrompt)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func select_(options []string) string {
	fmt.Println("Choose a license:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid selection. Exiting.")
		os.Exit(1)
	}

	return options[choice-1]
}