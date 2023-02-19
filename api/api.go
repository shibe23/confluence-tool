package api

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
	"os"
)

type API struct{}

func NewAPIClient() *goconfluence.API {
	domain := os.Getenv("CONFLUENCE_DOMAIN")
	username := os.Getenv("CONFLUENCE_USER")
	token := os.Getenv("CONFLUENCE_TOKEN")

	api, err := goconfluence.NewAPI(fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api", domain), username, token)
	if err != nil {
		log.Fatal(err)
	}
	return api
}
