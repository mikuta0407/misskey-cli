/*
Copyright © 2022 mikuta0407
*/
package cmd

import (
	"fmt"

	"github.com/mikuta0407/misskey-cli/misskey"
	"github.com/spf13/cobra"
)

// renoteCmd represents the renote command
var renoteCmd = &cobra.Command{
	Use:     "renote",
	Short:   "Renote note by noteId",
	Long:    `Renote note by noteId`,
	Example: `renote -i hoge 90ab12cd ねこはいます`,
	Run: func(cmd *cobra.Command, args []string) {

		client := misskey.NewClient(instanceName, cfgFile)

		if len(args) > 1 {
			fmt.Println("too many args")
			return
		}
		if len(args) == 0 {
			fmt.Println("Please give noteID")
			return
		}
		client.RenoteNote(args[0])
	},
}

func init() {
	rootCmd.AddCommand(renoteCmd)

}
