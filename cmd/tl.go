package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mikuta0407/misskey-cli/misskey"
)

// tlCmd represents the tl command
var tlCmd = &cobra.Command{
	Use:   "tl",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := misskey.NewClient(instanceName, cfgFile)
		if err := client.GetTimeline(limit, mode); err != nil {
			fmt.Println(err)
		}
	},
}

var limit int
var mode string

func init() {
	rootCmd.AddCommand(tlCmd)

	tlCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit display items(default: 10)")
	tlCmd.Flags().StringVarP(&mode, "mode", "m", "local", "TimeLine mode(local(default)/home/global)")

}
