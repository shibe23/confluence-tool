package usecases

import (
	"confluence-tool/content"
	"confluence-tool/mock"
	"testing"
)

func TestCreatePagesByCustomTitle(t *testing.T) {
	t.Run("CreatePagesByCustomTitleはタイトルの${value}とついた部分を指定した変数に置き換えて複数ページを作成する", func(t *testing.T) {
		client := mock.NewClient()

		data := content.Request{
			TemplateID: "12345678",
			Space:      "~1111111111",
			Ancestor:   "2222222222",
			Title:      "2023-01-30 MTG議事録 - ${value} さん",
		}

		err := CreatePagesByCustomTitle(client, data)
		if err != nil {
			t.Errorf("CreatePagesByCustomTitle() has error: %v", err)
		}
	})
}
