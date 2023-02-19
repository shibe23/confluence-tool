package usecases

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"strings"
)

func CreatePagesByTitle(api *goconfluence.API, template string, space string, ancestor string, title string, variables []string) (bool, error) {

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
		_, err := api.CreateContent(data)
		if err != nil {
			fmt.Printf("create content error: %+v", err)
			return false, err
		}
	}
	return true, nil
}

func GetTemplate(api *goconfluence.API, id string) (string, error) {
	content, err := api.GetContentByID(id, goconfluence.ContentQuery{
		Expand: []string{"body.storage"},
	})

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return content.Body.Storage.Value, nil
}
