package cmd

import (
	"confluence-tool/lib"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func parseJSON(path string, target interface{}) bool {
	// read file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("can't open file in %v. error: %v", path, err)
		return true
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(target)
	if err != nil {
		fmt.Printf("can't decode json. error:%v", err)
	}
	return false
}

func parseTextWithNewLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("can't open file in %v. error: %v", path, err)
		return nil
	}

	buffer, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("can't read variables in this flie. error:%v", err)
		return nil
	}

	variables := lib.Parse(string(buffer))
	return variables
}
