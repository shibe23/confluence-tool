package lib

import (
	"confluence-tool/content"
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		variables string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Parseは改行区切りで受け取った名前をひとつずつ分ける", args{variables: "鈴木\n佐藤\n田中\n堀井"}, []string{"鈴木", "佐藤", "田中", "堀井"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.variables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractVariables(t *testing.T) {
	s := `{"title":"2023-01-30 MTG議事録 - ${temp} さん", "space": "~1111111111", "ancestor": "2222222222", "templateID": "12345678"}`
	param := content.Parameter{}
	t.Run("ExtractVariablesはJSONをVariable Typeに変換する", func(t *testing.T) {
		err := ExtractVariables(s, &param)
		if err != nil {
			t.Errorf("ExtractVariables() failed: %v", err)
		}

		want := content.Parameter{
			TemplateID: "12345678",
			Space:      "~1111111111",
			Ancestor:   "2222222222",
			Title:      "2023-01-30 MTG議事録 - ${temp} さん",
		}

		isEqual := reflect.DeepEqual(param, want)
		if !isEqual {
			t.Errorf("content.Data{} is not equal.\nwant: %+v\ngot : %+v\n", want, param)
		}
		fmt.Printf("%+v\n", param)
	})
}
