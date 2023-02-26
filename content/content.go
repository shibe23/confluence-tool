package content

type Parameter struct {
	Space      string `json:"space"`
	Ancestor   string `json:"ancestor"`
	TemplateID string `json:"templateID"`
	Title      string `json:"title"`
}

type ConfluencePageInfo struct {
	Space    string
	Ancestor string
	Template string
	Title    string
}
