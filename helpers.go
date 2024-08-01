package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func fillContent(license *LicenseContent, skipPrompt bool) {
	name := getName(skipPrompt)
	year := getYear(skipPrompt)

	body := license.Body
	body = strings.ReplaceAll(body, "[year]", year)
	body = strings.ReplaceAll(body, "[yyyy]", year)
	body = strings.ReplaceAll(body, "[fullname]", name)
	body = strings.ReplaceAll(body, "[name of copyright owner]", name)
	body = strings.ReplaceAll(body, "<year>", year)
	body = strings.ReplaceAll(body, "<name of author>", name)

	err := writeFile("LICENSE", body)
	if err != nil {
		fmt.Printf("✘ An error occurred: %v\n", err)
	} else {
		fmt.Println("✔ License created successfully")
		fmt.Println("Please review the license and make any necessary changes.")
	}
}

func getGitUsername() (string, error) {
	cmd := exec.Command("git", "config", "--global", "--get", "user.name")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func getName(skipPrompt bool) string {
	gitUsername, err := getGitUsername()
	if err == nil && skipPrompt {
		return gitUsername
	}

	var name string
	if gitUsername != "" {
		fmt.Printf("Enter your name (default: %s): ", gitUsername)
	} else {
		fmt.Print("Enter your name: ")
	}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}

	if name == "" && gitUsername != "" {
		return gitUsername
	}
	return name
}

func getYear(skipPrompt bool) string {
	currentYear := time.Now().Year()
	if skipPrompt {
		return fmt.Sprintf("%d", currentYear)
	}

	fmt.Printf("Enter year (default: %d): ", currentYear)
	var year string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		year = scanner.Text()
	}

	if year == "" {
		return fmt.Sprintf("%d", currentYear)
	}
	return year
}

func writeFile(path string, content string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Print("A LICENSE file already exists. Enter a new name or press Enter to overwrite: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			newPath := scanner.Text()
			if newPath != "" {
				path = newPath
			}
		}
	}

	return os.WriteFile(path, []byte(content), 0644)
}