package confluence_tool

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
)

type APIClient struct{}

func (c APIClient) GetAuthedAPI(domain string, username string, token string) *goconfluence.API {
	api, err := goconfluence.NewAPI(fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", domain), username, token)
	if err != nil {
		log.Fatal(err)
	}
	return api
}

func (c APIClient) GetTemplate() error {
	error := 1
	if error != 1 {
		return fmt.Errorf("api call error: %v", error)
	}
	return nil
}
