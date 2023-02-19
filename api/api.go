package api

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
	"os"
	"strings"
)

type APIClient struct {
	api *goconfluence.API
}

func CreateAPIClient() *APIClient {
	domain := os.Getenv("CONFLUENCE_DOMAIN")
	username := os.Getenv("CONFLUENCE_USER")
	token := os.Getenv("CONFLUENCE_TOKEN")

	client := NewAPIClient(domain, username, token)

	if client == nil {
		fmt.Printf("createTestAPIClient() = %v", client)
	}

	return client
}

func NewAPIClient(domain string, username string, token string) *APIClient {
	api, err := goconfluence.NewAPI(fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", domain), username, token)
	if err != nil {
		log.Fatal(err)
	}
	return &APIClient{
		api: api,
	}
}

func (c APIClient) GetTemplate(id string) (string, error) {
	content, err := c.api.GetContentByID(id, goconfluence.ContentQuery{
		Expand: []string{"body.storage"},
	})

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return content.Body.Storage.Value, nil
}

func (c APIClient) CreatePagesByTitle(template string, space string, ancestor string, title string, variables []string) (bool, error) {

	data := &goconfluence.Content{
		Type: "page",
		Ancestors: []goconfluence.Ancestor{
			{
				ID: ancestor, // ancestor-id optional if you want to create sub-pages
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          template, // your page content here
				Representation: "storage",
			},
		},
		Version: &goconfluence.Version{
			Number: 1,
		},
		Space: &goconfluence.Space{
			Key: space,
		},
	}

	for _, v := range variables {
		replacedTitle := strings.ReplaceAll(title, "${temp}", v)
		data.Title = replacedTitle
		fmt.Printf("%+v", data)
		_, err := c.api.CreateContent(data)
		if err != nil {
			fmt.Printf("create content error: %+v", err)
			return false, err
		}
	}
	return true, nil
}
