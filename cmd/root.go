package cmd

import (
	//"fmt"

	"fmt"
	"os"

	"github.com/mikuta0407/misskey-cli/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var instanceName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "misskey-cli",
	Short: "",
	Long:  ``,
}

func Execute() {
	fmt.Printf("misskey-cli  ")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.misskey-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&instanceName, "instance", "i", "", "connect instance name(not host name)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	if cfgFile == "" {
		home, _ := os.UserHomeDir()
		cfgFile = home + "/.config/misskey-cli.toml"
	}
}

var configs config.Config
