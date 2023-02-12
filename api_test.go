package confluence_tool

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
