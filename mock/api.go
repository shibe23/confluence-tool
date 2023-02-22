package mock

import (
	"confluence-tool/content"
	"fmt"
)

type Client interface {
	GetTemplateByID(templateID string) (string, error)
	CreateContent(content content.Data) error
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetTemplateByID(templateID string) (string, error) {
	fmt.Printf("mock templateID: %v\n", templateID)
	return "", nil
}

func (c *client) CreateContent(content content.Data) error {
	fmt.Printf("mock content: %v\n", content)
	return nil
}
