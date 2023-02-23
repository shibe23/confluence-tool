package lib

import (
	"confluence-tool/content"
	"encoding/json"
	"fmt"
)

func Parse(variable string) []string {
	return []string{"鈴木", "佐藤", "田中", "堀井"}
}

func ExtractVariables(s string, variables *content.Data) error {
	err := json.Unmarshal([]byte(s), variables)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %v\n", err)
	}
	return nil
}
