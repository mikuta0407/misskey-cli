package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and build revision.",
	Long:  `Show version and build revision.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
