package models

import (
	"slices"
)

type Empty struct{}

type Person struct {
	ID    int
	Name  string
	Title string
}

func MakePerson(id int, name string, title string) Person {
	return Person{
		ID:    id,
		Name:  name,
		Title: title,
	}
}

type Link struct {
	RootID int
	SubID  int
}

func MakeLink(root int, sub int) Link {
	return Link{
		RootID: root,
		SubID:  sub,
	}
}

type Chart struct {
	ID    string
	Links []Link
}

func MakeChart(links []Link) Chart {
	return Chart{
		Links: links,
	}
}

type ChartRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func MakeChartRef(id string, name string) ChartRef {
	return ChartRef{
		ID:   id,
		Name: name,
	}
}

type FilePart struct {
	ID      string `json:"id"`
	Piece   int    `json:"piece"`
	Content string `json:"content"`
}

type PartialFileRef struct {
	FileID string
	Parts  int
	Name   string
}

func MakePartialFileRef(fileID string, parts int, name string) PartialFileRef {
	return PartialFileRef{
		FileID: fileID,
		Parts:  parts,
		Name:   name,
	}
}

type FileRef struct {
	FileID string `json:"id"`
	Name   string `json:"name"`
}

func MakeFilePart(id string, piece int, content string) FilePart {
	return FilePart{
		ID:      id,
		Piece:   piece,
		Content: content,
	}
}

type AppState struct {
	ActiveUploads map[PartialFileRef][]FilePart
	Charts        []ChartRef
}

func (a *AppState) StartUpload(start PartialFileRef) {
	a.ActiveUploads[start] = []FilePart{}
}

func (a *AppState) AddPart(part FilePart) *PartialFileRef {
	activeUploads := a.ActiveUploads
	for ref, parts := range activeUploads {
		if ref.FileID != part.ID {
			continue
		}

		if ref.Parts <= part.Piece {
			return nil
		}

		updatedParts := append(parts, part)

		slices.SortStableFunc(
			updatedParts,
			func (first FilePart, second FilePart) int {
				return first.Piece - second.Piece
			},
		)

		slices.Reverse(updatedParts)

		updatedParts = slices.CompactFunc(
			updatedParts,
			func (first FilePart, second FilePart) bool {
				return first.Piece == second.Piece
			},
		)

		activeUploads[ref] = updatedParts

		a.ActiveUploads = activeUploads

		return &ref
	}

	return nil
}

func EmptyAppState() AppState {
	return AppState{
		ActiveUploads: make(map[PartialFileRef][]FilePart),
		Charts:        []ChartRef{},
	}
}
