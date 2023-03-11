package content

type Parameter struct {
	Keys []Request
}

type Request struct {
	Title      string `json:"title"`
	Value      string `json:"value"`
	Space      string `json:"space"`
	Ancestor   string `json:"ancestor"`
	TemplateID string `json:"templateID"`
}

type ConfluencePageInfo struct {
	Title    string
	Space    string
	Ancestor string
	Template string
}
