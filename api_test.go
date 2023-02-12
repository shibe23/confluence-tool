package confluence_tool

import (
	"fmt"
	"os"
	"testing"
)

func createTestAPIClient() APIClient {
	c := APIClient{}

	domain := os.Getenv("CONFLUENCE_DOMAIN")
	username := os.Getenv("CONFLUENCE_USER")
	token := os.Getenv("CONFLUENCE_TOKEN")

	api := c.GetAuthedAPI(domain, username, token)

	if api == nil {
		fmt.Printf("createTestAPIClient() = %v", api)
	}
	return c
}

func TestGetTemplate(t *testing.T) {
	c := createTestAPIClient()

	t.Run("テンプレートを作成する", func(t *testing.T) {
		if _, err := c.GetTemplate(); err != nil {
			t.Errorf("GetTemplate() hass error: %v", err)
		}
	})
}
