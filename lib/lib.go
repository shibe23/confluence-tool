package lib

import (
	"confluence-tool/content"
	"encoding/json"
)

func Parse(variable string) []string {
	return []string{"鈴木", "佐藤", "田中", "堀井"}
}

func ExtractVariables(s []byte, variables *content.Data) error {
	return json.Unmarshal(s, variables)
}
