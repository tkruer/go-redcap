package redcap_test

import (
	"testing"

	redcap "github.com/tkruer/go-redcap/pkg"
)

func RunUnitTests(t *testing.T, client redcap.RedCapClient) {
	t.Run("TestImportArms", func(t *testing.T) { importArmsUnitTest(t, client) })
}

func importArmsUnitTest(t *testing.T, client redcap.RedCapClient) {
	// Example arms to import
	armsToImport := []string{"Arm 1", "Arm 2"}
	// Call the ImportArms method
	importedArms, err := client.ImportArms(armsToImport)
	if err != nil {
		t.Fatalf("Failed to import arms: %v", err)
	}
	// Check if the imported arms match the expected values
	if len(importedArms) != len(armsToImport) {
		t.Fatalf("Expected %d arms, got %d", len(armsToImport), len(importedArms))
	}
	for i, arm := range importedArms {
		if arm.Name != armsToImport[i] {
			t.Fatalf("Expected arm name %s, got %s", armsToImport[i], arm.Name)
		}
	}
	t.Logf("Successfully imported arms: %+v", importedArms)
}
