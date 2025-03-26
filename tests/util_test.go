package redcap_test

import (
	"encoding/csv"
	"log"
	"net/http"
	"strings"
)

type TestSuiteConfig struct {
	RedcapApiUrl   string
	RedcapApiToken string
}

func LoadRedcapConfig() TestSuiteConfig {
	const csvURL = "https://raw.githubusercontent.com/redcap-tools/redcap-test-datasets/master/connections/oklahoma-bbmc.csv"

	resp, err := http.Get(csvURL)
	if err != nil {
		log.Fatalf("failed to fetch CSV: %v", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV records: %v", err)
	}

	var redcapTestURL string
	var redcapTestAPIToken string

	for _, record := range records[1:] {
		if strings.Contains(strings.ToLower(record[7]), "read & write") {
			redcapTestURL = record[2]
			redcapTestAPIToken = record[1]
			break // stop at the first match
		}
	}

	return TestSuiteConfig{
		RedcapApiUrl:   redcapTestURL,
		RedcapApiToken: redcapTestAPIToken,
	}
}
