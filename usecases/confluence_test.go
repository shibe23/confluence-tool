package usecases

import (
	"confluence-tool/content"
	"confluence-tool/mock"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	client := mock.NewClient()
	id := ""

	t.Run("GetTemplateは指定したIDのテンプレートを取得する", func(t *testing.T) {
		if _, err := GetTemplate(client, id); err != nil {
			t.Errorf("GetTemplate() has error: %v", err)
		}
	})
}

func TestCreatePagesByTitle(t *testing.T) {
	t.Run("CreatePagesByTitleはタイトルの${temp}とついた部分を指定した変数に置き換えて複数ページを作成する", func(t *testing.T) {
		client := mock.NewClient()

		data := content.ConfluencePageInfo{
			Template: "<p>test</p>",
			Space:    "~1111111111",
			Ancestor: "2222222222",
			Title:    "2023-01-30 MTG議事録 - ${temp} さん",
		}
		variables := []string{"test 太郎", "test 健一", "test 秀子"}

		err := CreatePagesByTitle(client, data, variables)
		if err != nil {
			t.Errorf("CreatePagesByTitle() has error: %v", err)
		}
	})
}
