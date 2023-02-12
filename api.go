package confluence_tool

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
)

type APIClient struct {
	api *goconfluence.API
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
