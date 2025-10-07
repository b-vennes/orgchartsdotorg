package files

import (
	"testing"

	"orgcharts.org/api/pkg/models"
)

func TestHasAllParts(t *testing.T) {
	fileName := "my-json-file"
	parts := []models.FilePart{
		models.MakeFilePart(fileName, 0, "{ \"hello\": "),
		models.MakeFilePart(fileName, 1, "100 }"),
	}

	result := HasAllParts(fileName, 2, parts)

	if !result {
		t.Error("Expected result to be 'true'.")
	}
}

func TestCombineParts(t *testing.T) {
	fileName := "my-json-file"
	parts := []models.FilePart{
		models.MakeFilePart(fileName, 2, "}"),
		models.MakeFilePart(fileName, 0, "{\"key\": "),
		models.MakeFilePart(fileName, 1, "true, \"other-key\": 1000"),
	}

	expected := "{\"key\": true, \"other-key\": 1000}"

	result := CombineParts(parts)

	if result != expected {
		t.Error("Expected", result, "to be", expected)
	}
}
