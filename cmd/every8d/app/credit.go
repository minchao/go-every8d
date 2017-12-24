package app

import (
	"context"

	"github.com/spf13/cobra"
)

var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "Query credit",
	Long:  "Query to retrieve your account balance",
	Run:   creditFunc,
}

func creditFunc(cmd *cobra.Command, args []string) {
	credit, err := client.GetCredit(context.Background())
	if err != nil {
		er(err)
	}

	cmd.Printf("Credit: %.2f\n", credit)
}
