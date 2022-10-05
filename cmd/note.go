/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mikuta0407/misskey-cli/config"
	"github.com/mikuta0407/misskey-cli/misskey"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("note called")
		if len(args) > 1 {
			fmt.Println("too many args")
			return
		}
		if len(args) == 0 {
			fmt.Println("Please write note")
			return
		}
		noteMain(args[0])
	},
}

func init() {
	rootCmd.AddCommand(noteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func noteMain(text string) {
	configs, err := config.ParseToml(cfgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	misskey.PostNote(configs, instanceName, text)
}
