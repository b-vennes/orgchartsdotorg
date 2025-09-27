package models

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

func (a *AppState) AddPart(part FilePart) {
	activeUploads := a.ActiveUploads
	for ref, parts := range activeUploads {
		if ref.FileID == part.ID {
			updatedParts := append(parts, part)
			activeUploads[ref] = updatedParts

			a.ActiveUploads = activeUploads

			return
		}
	}
}

func EmptyAppState() AppState {
	return AppState{
		ActiveUploads: make(map[PartialFileRef][]FilePart),
		Charts:        []ChartRef{},
	}
}
