/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mikuta0407/misskey-cli/misskey"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "streaming timeline",
	Long:  `Streaming timeline like UserStream`,
	Run: func(cmd *cobra.Command, args []string) {
		client := misskey.NewClient(instanceName, cfgFile)
		if err := client.GetStream(mode); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)

	streamCmd.Flags().StringVarP(&mode, "mode", "m", "local", "TimeLine mode(local/home/global)")
}
