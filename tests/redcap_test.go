package redcap_test

import (
	"testing"

	redcap "github.com/tkruer/go-redcap/pkg"
)

func TestRedcapTestSuite(t *testing.T) {
	config := LoadRedcapConfig()
	testClient := redcap.RedCapClient{
		URL:            config.RedcapApiUrl,
		Token:          config.RedcapApiToken,
		ResponseFormat: redcap.JSON,
	}
	t.Logf("Created REDCap client with URL: %s", testClient.URL)

	t.Logf("Now starting functional tests...")

	t.Logf("Now starting unit tests...")
	RunUnitTests(t, testClient)
	t.Logf("Now starting integration tests...")
	RunIntegrationTests(t, testClient)
	t.Logf("Now starting functional tests...")
	RunFunctionalTests(t, testClient)
}
