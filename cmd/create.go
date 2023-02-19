/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"confluence-tool/api"
	"confluence-tool/usecases"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create tempalte space ancestor, title, variables",
	Long:  `指定したテンプレートIDの内容で新規ページを作成する`,
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewAPIClient()

		templateID := args[0]

		template, err := usecases.GetTemplate(client, templateID)
		if err != nil {
			fmt.Errorf("getTemplateError: %v", err)
		}

		space := args[1]
		ancestor := args[2]
		title := args[3]
		variables := strings.Split(args[4], "\n")

		_, err = usecases.CreatePagesByTitle(client, template, space, ancestor, title, variables)
		if err != nil {
			fmt.Errorf("CreatePagesByTitle() has error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
