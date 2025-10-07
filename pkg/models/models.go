package models

import (
	"slices"
)

type Empty struct{}

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

func MakePerson(id int, name string, title string) Person {
	return Person{
		ID:    id,
		Name:  name,
		Title: title,
	}
}

type Link struct {
	RootID int `json:"rootID"`
	SubID  int `json:"subID"`
}

func MakeLink(root int, sub int) Link {
	return Link{
		RootID: root,
		SubID:  sub,
	}
}

type Chart struct {
	Ref    ChartRef `json:"ref"`
	Links  []Link   `json:"links"`
	People []Person `json:"people"`
}

func MakeChart(ref ChartRef, links []Link, people []Person) Chart {
	return Chart{
		Ref:    ref,
		Links:  links,
		People: people,
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

type UnparsedChart struct {
	ID      string
	Name    string
	Content string
}

func MakeUnparsedChart(id string, name string, content string) UnparsedChart {
	return UnparsedChart{
		ID:      id,
		Name:    name,
		Content: content,
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

type FileState struct {
	ActiveUploads map[PartialFileRef][]FilePart
}

func (a *FileState) StartUpload(start PartialFileRef) {
	a.ActiveUploads[start] = []FilePart{}
}

func (a *FileState) AddPart(part FilePart) *PartialFileRef {
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
			func(first FilePart, second FilePart) int {
				return first.Piece - second.Piece
			},
		)

		slices.Reverse(updatedParts)

		updatedParts = slices.CompactFunc(
			updatedParts,
			func(first FilePart, second FilePart) bool {
				return first.Piece == second.Piece
			},
		)

		activeUploads[ref] = updatedParts

		a.ActiveUploads = activeUploads

		return &ref
	}

	return nil
}

func EmptyFileState() FileState {
	return FileState{
		ActiveUploads: make(map[PartialFileRef][]FilePart),
	}
}

type ChartState struct {
	Charts []Chart
}

func EmptyChartState() ChartState {
	return ChartState{
		Charts: []Chart{},
	}
}
