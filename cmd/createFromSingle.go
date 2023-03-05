/*
Copyright © 2023 shibe23 <shibe23.info@gmail.com>
*/
package cmd

import (
	"confluence-tool/api"
	"confluence-tool/content"
	"confluence-tool/lib"
	"confluence-tool/usecases"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var (
	pageInfoFilePath string
	variables        string
)

// createCmd represents the create command
var createFromSingleCmd = &cobra.Command{
	Use:   "create-from-s",
	Short: "1つのテンプレートから新規ページを作成する",
	Long:  `1つのテンプレートから新規ページを作成する。タイトルの一部を任意の文字列に変更できる`,
	Run: func(cmd *cobra.Command, args []string) {
		// get target path from file.
		params := content.Parameter{}
		parseJSON(pageInfoFilePath, &params)

		// get variable data from file.
		text, err := cmd.PersistentFlags().GetString("variables")
		if err != nil {
			fmt.Println("Invalid values.")
			os.Exit(1)
		}

		variables := parseTextWithNewLine(text)

		client := api.NewClient()
		template, err := client.GetTemplateByID(params.TemplateID)
		if err != nil {
			fmt.Printf("templateID is invalid. error: %v", err)
			os.Exit(1)
		}

		data := content.ConfluencePageInfo{
			Template: template,
			Space:    params.Space,
			Ancestor: params.Ancestor,
			Title:    params.Title,
		}

		err = usecases.CreatePagesByTitle(client, data, variables)
		if err != nil {
			fmt.Printf("createPagesByTitle is invalid. error: %v", err)
		}
	},
}

func parseJSON(path string, t interface{}) bool {
	// read file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("can't open file in %v. error: %v", path, err)
		return true
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(t)
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

func init() {
	rootCmd.AddCommand(createFromSingleCmd)
	createFromSingleCmd.PersistentFlags().StringVar(&pageInfoFilePath, "page-info-file", "", "対象となるConfluenceについての情報をまとめたファイルへのパス")
	createFromSingleCmd.PersistentFlags().StringVar(&variables, "variables-file", "", "ページごとに差し替える文字列.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
