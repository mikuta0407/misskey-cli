package cmd

import (
	//"fmt"
	"os"

	"github.com/mikuta0407/misskey-cli/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var instanceName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "misskey-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.misskey-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&instanceName, "instance", "i", "", "connect instance name(not host name)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile == "" {
	// 	home, err := os.UserHomeDir()
	// 	cobra.CheckErr(err)

	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName(".misskey-cli")

	// }

	// viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }

	// // 設定ファイルの内容を構造体にコピーする
	// if err := viper.Unmarshal(&config); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}

var configs config.Config
