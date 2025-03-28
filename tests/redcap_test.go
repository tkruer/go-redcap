package redcap_test

import (
	"net/http"
	"testing"

	redcap "github.com/tkruer/go-redcap/pkg"
)

func TestRedcapTestSuite(t *testing.T) {
	config := LoadRedcapConfig()
	testClient := redcap.RedCapClient{
		URL:            config.RedcapApiUrl,
		Token:          config.RedcapApiToken,
		ResponseFormat: redcap.JSON,
		HTTPClient:     &http.Client{}, // Use the default HTTP client
	}
	t.Logf("Created REDCap client with URL: %s", testClient.URL)

	t.Logf("Created REDCap client with TOken  : %s", testClient.Token)

	t.Logf("Now starting functional tests...")

	t.Logf("Now starting unit tests...")
	RunUnitTests(t, testClient)
	t.Logf("Now starting integration tests...")
	RunIntegrationTests(t, testClient)
	t.Logf("Now starting functional tests...")
	RunFunctionalTests(t, testClient)
}
