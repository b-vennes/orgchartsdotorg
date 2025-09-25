package models

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
