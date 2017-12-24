package app

import (
	"github.com/minchao/go-every8d"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	client *every8d.Client

	rootCmd = &cobra.Command{
		Use:   "every8d",
		Short: "EVERY8D SMS CLI tool",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			username := viper.GetString("username")
			password := viper.GetString("password")

			client = every8d.NewClient(username, password, nil)
		},
	}
)

func init() {
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().String("username", "", "EVERY8D Username")
	rootCmd.PersistentFlags().String("password", "", "EVERY8D Password")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))

	rootCmd.AddCommand(creditCmd)
	rootCmd.AddCommand(deliveryStatusCmd)
}

func Execute() {
	rootCmd.Execute()
}
