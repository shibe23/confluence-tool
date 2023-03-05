/*
Copyright © 2023 shibe23 <shibe23.info@gmail.com>
*/
package cmd

import (
	"confluence-tool/api"
	"confluence-tool/content"
	"confluence-tool/usecases"
	"fmt"
	"github.com/spf13/cobra"
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

func init() {
	rootCmd.AddCommand(createFromSingleCmd)
	createFromSingleCmd.PersistentFlags().StringVar(&pageInfoFilePath, "page-info-file", "", "対象となるConfluenceについての情報をまとめたファイルへのパス")
	createFromSingleCmd.PersistentFlags().StringVar(&variables, "variables-file", "", "ページごとに差し替える文字列.")
}
