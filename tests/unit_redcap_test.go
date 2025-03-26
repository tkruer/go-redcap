package redcap_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	redcap "github.com/tkruer/go-redcap/pkg"
)

func RunUnitTests(t *testing.T, client redcap.RedCapClient) {
	t.Run("TestImportArms", func(t *testing.T) { importArmsUnitTest(t, client) })
}

func importArmsUnitTest(t *testing.T, client redcap.RedCapClient) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		if !strings.Contains(string(body), `"arm_num":"1"`) {
			t.Errorf("unexpected request body: %s", body)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"count":1}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := client.ImportArms()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `{"count":1}`
	if string(resp) != expected {
		t.Errorf("expected %s, got %s", expected, string(resp))
	}
}
