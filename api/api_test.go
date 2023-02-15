package api

import (
	"fmt"
	"os"
	"testing"
)

func createTestAPIClient() *APIClient {
	domain := os.Getenv("CONFLUENCE_DOMAIN")
	username := os.Getenv("CONFLUENCE_USER")
	token := os.Getenv("CONFLUENCE_TOKEN")

	client := NewAPIClient(domain, username, token)

	if client == nil {
		fmt.Printf("createTestAPIClient() = %v", client)
	}
	return client
}

func TestGetTemplate(t *testing.T) {
	c := createTestAPIClient()
	id := ""

	t.Run("テンプレートを作成する", func(t *testing.T) {
		if _, err := c.GetTemplate(id); err != nil {
			t.Errorf("GetTemplate() has error: %v", err)
		}
	})
}

func TestCreatePagesByTitle(t *testing.T) {
	t.Run("テンプレートを作成する", func(t *testing.T) {
		c := createTestAPIClient()

		template := "<p>test</p>"
		space := "~1111111111"
		ancestor := "2222222222"
		title := "2023-01-30 MTG議事録 - ${temp} さん"
		variables := []string{"test 太郎", "test 健一", "test 秀子"}

		_, err := c.CreatePagesByTitle(template, space, ancestor, title, variables)
		if err != nil {
			t.Errorf("CreatePagesByTitle() has error: %v", err)
		}
	})
}
