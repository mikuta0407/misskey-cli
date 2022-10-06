package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"github.com/go-playground/validator"
	"github.com/mikuta0407/misskey-cli/config"
	"github.com/mikuta0407/misskey-cli/misskey"
)

// tlCmd represents the tl command
var tlCmd = &cobra.Command{
	Use:   "tl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tl called")
		tlMain()
	},
}

func init() {
	rootCmd.AddCommand(tlCmd)

	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 10, "Limit display items(default: 10)")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "local", "TimeLine mode(local(default)/home/global)")
	//rootCmd.PersistentFlags().BoolVar(&local, "local", false, "local only (default: false")

}

var limit int
var mode string

func tlMain() {
	configs, err := config.ParseToml(cfgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	misskey.GetTl(configs, instanceName, limit, mode)

}
