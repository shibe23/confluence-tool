package create_subpage

import (
	"reflect"
	"testing"
)

func TestParseValues(t *testing.T) {
	c := CreateSubpage{}

	type args struct {
		variables string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"改行区切りで受け取った名前をひとつずつ分ける", args{variables: "鈴木\n佐藤\n田中\n堀井"}, []string{"鈴木", "佐藤", "田中", "堀井"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Parse(tt.args.variables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
