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

		configs, err := config.ParseToml(cfgFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		if deleteId == "" && replyId == "" {
			if len(args) > 1 {
				fmt.Println("too many args")
				return
			}
			if len(args) == 0 {
				fmt.Println("Please write note")
				return
			}
			misskey.CreateNote(configs, instanceName, args[0])
		} else if deleteId != "" && replyId == "" {
			misskey.DeleteNote(configs, instanceName, deleteId)
		} else if deleteId == "" && replyId != "" {
			if len(args) > 1 {
				fmt.Println("too many args")
				return
			}
			if len(args) == 0 {
				fmt.Println("Please write note")
				return
			}
			misskey.ReplyNote(configs, instanceName, replyId, args[0])
		} else {
			fmt.Println("Please one Option")
			return
		}

	},
}

var (
	delete   string
	children string
	replyId  string
	deleteId string
)

func init() {
	rootCmd.AddCommand(noteCmd)

	// 削除
	noteCmd.Flags().StringVarP(&deleteId, "delete", "d", "", "Delete notes (id)")

	// リプライ
	noteCmd.Flags().StringVarP(&replyId, "reply", "r", "", "reply note")

	// 公開範囲の話
	//noteCmd.Flags().StringVarP(&reply, "", "", "", "")

}
