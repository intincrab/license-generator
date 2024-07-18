package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

type LicenseContent struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
	Conditions  []string `json:"conditions"`
	Limitations []string `json:"limitations"`
	Body        string   `json:"body"`
}

func fetchLicenses() []License {
	resp, err := http.Get("https://api.github.com/licenses")
	if err != nil {
		panic(fmt.Sprintf("Error fetching licenses: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Error reading response body: %v", err))
	}

	var licenses []License
	err = json.Unmarshal(body, &licenses)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON: %v", err))
	}

	return licenses
}

func getLicenseNames(licenses []License) []string {
	var names []string
	for _, license := range licenses {
		names = append(names, license.Name)
	}
	return names
}

func getLicenseKeys(licenses []License) []string {
	var keys []string
	for _, license := range licenses {
		keys = append(keys, license.Key)
	}
	return keys
}

func getLicenseFromName(licenses []License, name string) LicenseContent {
	for _, license := range licenses {
		if license.Name == name {
			return fetchLicenseContent(license.URL)
		}
	}
	panic(fmt.Sprintf("License not found: %s", name))
}

func getLicenseFromKey(licenses []License, key string) LicenseContent {
	for _, license := range licenses {
		if license.Key == key {
			return fetchLicenseContent(license.URL)
		}
	}
	panic(fmt.Sprintf("License not found: %s", key))
}

func fetchLicenseContent(url string) LicenseContent {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Error fetching license content: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Error reading response body: %v", err))
	}

	var licenseContent LicenseContent
	err = json.Unmarshal(body, &licenseContent)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON: %v", err))
	}

	return licenseContent
}