package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	redcap "github.com/tkruer/go-redcap/pkg"
)

func main() {
	// Define flags for the command line interface
	apiURL := flag.String("url", "https://api.redcap.com", "URL of the REDCap API")
	apiToken := flag.String("token", "", "API token for REDCap")
	format := flag.String("format", "json", "Response format (json or xml)")
	dagsInput := flag.String("dags", "", "Comma-separated list of DAGs to delete")

	// Parse the flags
	flag.Parse()

	// Simple input validation
	if *apiToken == "" {
		log.Fatal("API token is required")
	}
	if *dagsInput == "" {
		log.Fatal("At least one DAG is required")
	}

	// Split the comma-separated list of DAGs into a slice
	dags := strings.Split(*dagsInput, ",")

	// Create a new RedCapClient
	client := redcap.RedCapClient{
		Token:          *apiToken,
		URL:            *apiURL,
		ResponseFormat: redcap.ResponseFormat(*format),
	}

	// Delete the DAGs using the RedCapClient
	response, err := client.DeleteDags(dags)
	if err != nil {
		log.Fatalf("Error deleting DAGs: %v", err)
	}

	// Print the response
	fmt.Println(string(response))
}