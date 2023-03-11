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
)

var (
	path      string
	variables string
)

// createCmd represents the create command
var createFromSingleCmd = &cobra.Command{
	Use:   "create-from-s",
	Short: "1つのテンプレートから新規ページを作成する",
	Long:  `1つのテンプレートから新規ページを作成する。タイトルの一部を任意の文字列に変更できる`,
	Run: func(cmd *cobra.Command, args []string) {
		// get target path from file.
		params := content.Parameter{}
		parseJSON(path, &params)

		client := api.NewClient()

		for _, v := range params.Keys {
			err := usecases.CreatePagesByTitle(client, v)
			if err != nil {
				fmt.Printf("createPagesByTitle is invalid. error: %v", err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(createFromSingleCmd)
	createFromSingleCmd.PersistentFlags().StringVar(&path, "path", "", "対象となるConfluenceについての情報をまとめたファイルへのパス")
}
