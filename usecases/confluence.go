package usecases

import (
	"confluence-tool/api"
	"confluence-tool/content"
	"fmt"
	"strings"
)

func CreatePagesByTitle(api api.Client, data content.Request) error {
	template, err := api.GetTemplateByID(data.TemplateID)
	if err != nil {
		fmt.Printf("templateID is invalid. error: %v", err)
		return err
	}

	replacedTitle := strings.ReplaceAll(data.Title, "${temp}", data.Value)

	err = api.CreateContent(content.ConfluencePageInfo{
		Title:    replacedTitle,
		Template: template,
		Space:    data.Space,
		Ancestor: data.Ancestor,
	})
	if err != nil {
		fmt.Printf("api.CreateContentError: %v", err)
		return err
	}
	return nil
}
