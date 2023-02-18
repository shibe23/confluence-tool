/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"confluence-tool/api"
	"confluence-tool/cmd"
	"fmt"
	"os"
)

func createAPIClient() *api.APIClient {
	domain := os.Getenv("CONFLUENCE_DOMAIN")
	username := os.Getenv("CONFLUENCE_USER")
	token := os.Getenv("CONFLUENCE_TOKEN")

	client := api.NewAPIClient(domain, username, token)

	if client == nil {
		fmt.Printf("createTestAPIClient() = %v", client)
	}

	return client
}

func runCreatePageByTitle() {
	c := createAPIClient()

	templateID := "1111111111"

	template, err := c.GetTemplate(templateID)
	if err != nil {
		fmt.Errorf("getTemplateError: %v", err)
	}

	space := "~1111111111"
	ancestor := "2222222222"
	title := "2023-01-30 MTG議事録 - ${temp} さん"
	variables := []string{"田中 太郎", "鈴木 健一", "小林 秀子"}

	_, err = c.CreatePagesByTitle(template, space, ancestor, title, variables)
	if err != nil {
		fmt.Errorf("CreatePagesByTitle() has error: %v", err)
	}
}

func main() {
	cmd.Execute()
}
