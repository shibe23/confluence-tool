package mock

import (
	"confluence-tool/content"
	"fmt"
)

type Client interface {
	GetTemplateByID(templateID string) (string, error)
	CreateContent(content content.ConfluencePageInfo) error
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetTemplateByID(templateID string) (string, error) {
	fmt.Printf("mock templateID: %v\n", templateID)
	return "<p>this is mock template.</p>", nil
}

func (c *client) CreateContent(content content.ConfluencePageInfo) error {
	fmt.Printf("mock content: %v\n", content)
	return nil
}
