package api

import (
	"confluence-tool/content"
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
	"os"
)

func newGoConfluenceAPI() *goconfluence.API {
	domain := os.Getenv("CONFLUENCE_DOMAIN")
	username := os.Getenv("CONFLUENCE_USER")
	token := os.Getenv("CONFLUENCE_TOKEN")

	api, err := goconfluence.NewAPI(fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", domain), username, token)
	if err != nil {
		log.Fatal(err)
	}
	return api
}

type Client interface {
	GetTemplateByID(templateID string) (string, error)
	CreateContent(content content.Data) error
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetTemplateByID(templateID string) (string, error) {
	goConfluenceAPI := newGoConfluenceAPI()
	content, err := goConfluenceAPI.GetContentByID(templateID, goconfluence.ContentQuery{
		Expand: []string{"body.storage"},
	})

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	return content.Body.Storage.Value, nil
}

func (c *client) CreateContent(content content.Data) error {
	goConfluenceAPI := newGoConfluenceAPI()

	data := &goconfluence.Content{Type: "page",
		Ancestors: []goconfluence.Ancestor{
			{
				ID: content.Ancestor,
			},
		},
		Title: content.Title,
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          content.Template,
				Representation: "storage",
			},
		},
		Version: &goconfluence.Version{
			Number: 1,
		},
		Space: &goconfluence.Space{
			Key: content.Space,
		},
	}

	_, err := goConfluenceAPI.CreateContent(data)
	if err != nil {
		fmt.Printf("create content error: %+v", err)
		return err
	}
	return nil
}
