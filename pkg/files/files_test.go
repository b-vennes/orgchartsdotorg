package files

import (
    "testing"
)

func TestHasAllParts(t *testing.T) {
  fileName := "my-json-file"
  parts := []FilePart{
    MakeFilePart(fileName, 0, "{ \"hello\": "),
    MakeFilePart(fileName, 1, "100 }"),
  }

  result := HasAllParts(fileName, 2, parts)

  if !result {
    t.Error("Expected result to be 'true'.")
  }
}

func TestCombineParts(t *testing.T) {
  fileName := "my-json-file"
  parts := []FilePart{
    MakeFilePart(fileName, 2, "}"),
    MakeFilePart(fileName, 0, "{\"key\": "),
    MakeFilePart(fileName, 1, "true, \"other-key\": 1000"),
  }

  expected := "{\"key\": true, \"other-key\": 1000}"

  result := CombineParts(parts)

  if result != expected {
    t.Error("Expected", result, "to be", expected)
  }
}
