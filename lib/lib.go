package lib

import (
	"confluence-tool/content"
	"encoding/json"
	"fmt"
	"strings"
)

func Parse(variables string) []string {
	return strings.Split(variables, "\n")
}

func ExtractVariables(s string, variables *content.Parameter) error {
	err := json.Unmarshal([]byte(s), variables)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %v\n", err)
	}
	return nil
}
