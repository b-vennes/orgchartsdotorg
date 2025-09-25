package files

import "slices"

type FilePart struct {
  FileID string
  Piece int
  Content string
}

func MakeFilePart(id string, piece int, content string) FilePart {
  return FilePart{
    FileID: id,
    Piece: piece,
    Content: content,
  }
}

func HasAllParts(id string, needed int, parts []FilePart) bool {
  partNumbers := []int{}

  for _, p := range parts {
    if p.FileID == id && p.Piece < needed {
      partNumbers = append(partNumbers, p.Piece)
    }
  }

  slices.Sort(partNumbers)
  partNumbers = slices.Compact(partNumbers)

  return len(partNumbers) == needed
}

func CombineParts(parts []FilePart) string {
  slices.SortFunc(parts, func(a FilePart, b FilePart) int {
    return a.Piece - b.Piece
  })

  combined := ""

  for _, p := range parts {
    combined = combined + p.Content
  }

  return combined
}
