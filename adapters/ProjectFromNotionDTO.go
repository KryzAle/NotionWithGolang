package adapters

type ProjectFromNotionDTO struct {
	Results []Result `json:"results"`
}

type Result struct {
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	ProjectStatus ProjectStatus `json:"Project status"`
	Client        Client        `json:"Client"`
	ReportingTo   ReportingTo   `json:"Reporting to"`
	Name          Name          `json:"Name"`
}
type ProjectStatus struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Select Select `json:"select"`
}
type Client struct {
	Relation []Relation `json:"relation"`
}

type Relation struct {
	Id string `json:"id"`
}

type ReportingTo struct {
	RichText []RichText `json:"rich_text"`
}
type Name struct {
	Title []Title `json:"title"`
}

type RichText struct {
	Text      Text   `json:"text"`
	PlainText string `json:"plain_text"`
}
type Title struct {
	Text Text `json:"text"`
}
type Text struct {
	Content string `json:"content"`
}

type Select struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
