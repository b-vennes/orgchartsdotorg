package files

import (
  "slices"

  "orgcharts.org/api/pkg/models"
)

func HasAllParts(id string, needed int, parts []models.FilePart) bool {
  partNumbers := []int{}

  for _, p := range parts {
    if p.ID == id && p.Piece < needed {
      partNumbers = append(partNumbers, p.Piece)
    }
  }

  slices.Sort(partNumbers)
  partNumbers = slices.Compact(partNumbers)

  return len(partNumbers) == needed
}

func CombineParts(parts []models.FilePart) string {
  slices.SortFunc(parts, func(a models.FilePart, b models.FilePart) int {
    return a.Piece - b.Piece
  })

  combined := ""

  for _, p := range parts {
    combined = combined + p.Content
  }

  return combined
}
