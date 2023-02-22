package usecases

import (
	"confluence-tool/api"
	"confluence-tool/content"
	"fmt"
	"strings"
)

func CreatePagesByTitle(api api.Client, data content.Data, variables []string) error {
	for _, v := range variables {
		replacedTitle := strings.ReplaceAll(data.Title, "${temp}", v)

		err := api.CreateContent(content.Data{
			Template: data.Template,
			Space:    data.Space,
			Ancestor: data.Ancestor,
			Title:    replacedTitle,
		})
		if err != nil {
			fmt.Printf("api.CreateContentError: %v", err)
			return err
		}
	}
	return nil
}

func GetTemplate(api api.Client, id string) (string, error) {
	template, err := api.GetTemplateByID(id)
	if err != nil {
		fmt.Errorf("Can't get Template %v\n", err)
	}
	return template, nil
}
