package usecases

import (
	"confluence-tool/api"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	client := api.NewAPIClient()
	id := ""

	t.Run("テンプレートを作成する", func(t *testing.T) {
		if _, err := GetTemplate(client, id); err != nil {
			t.Errorf("GetTemplate() has error: %v", err)
		}
	})
}

func TestCreatePagesByTitle(t *testing.T) {
	t.Run("テンプレートを作成する", func(t *testing.T) {
		client := api.NewAPIClient()

		template := "<p>test</p>"
		space := "~1111111111"
		ancestor := "2222222222"
		title := "2023-01-30 MTG議事録 - ${temp} さん"
		variables := []string{"test 太郎", "test 健一", "test 秀子"}

		_, err := CreatePagesByTitle(client, template, space, ancestor, title, variables)
		if err != nil {
			t.Errorf("CreatePagesByTitle() has error: %v", err)
		}
	})
}
