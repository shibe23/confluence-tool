/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"confluence-tool/api"
	"confluence-tool/content"
	"confluence-tool/lib"
	"confluence-tool/usecases"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create template space ancestor, title, variables",
	Long:  `指定したテンプレートIDの内容で新規ページを作成する`,
	Run: func(cmd *cobra.Command, args []string) {
		param := content.Parameter{}
		lib.ExtractVariables(args[0], &param)

		client := api.NewClient()

		templateID := args[0]
		template, _ := usecases.GetTemplate(client, templateID)

		data := content.Data{
			Template: template,
			Space:    args[1],
			Ancestor: args[2],
			Title:    args[3],
		}
		variables := strings.Split(args[4], "\n")

		usecases.CreatePagesByTitle(client, data, variables)
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
