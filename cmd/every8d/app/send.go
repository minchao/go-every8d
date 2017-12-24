package app

import (
	"context"

	"github.com/minchao/go-every8d"
	"github.com/spf13/cobra"
)

var (
	sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send an SMS",
		Run:   sendFunc,
	}
)

func init() {
	sendCmd.Flags().StringP("sb", "s", "", "Message title. Empty titles are accepted")
	sendCmd.Flags().StringP("msg", "m", "", "Message content")
	sendCmd.Flags().StringP("dest", "d", "", "Receiver's mobile number")
	sendCmd.Flags().StringP("st", "R", "", "Reservation time")
	sendCmd.Flags().IntP("retryTime", "r", 0, "SMS validity period of unit: minutes")
}

func sendFunc(cmd *cobra.Command, _ []string) {
	message := every8d.Message{}
	message.Subject, _ = cmd.Flags().GetString("sb")
	message.Content, _ = cmd.Flags().GetString("msg")
	message.Destination, _ = cmd.Flags().GetString("dest")
	message.ReservationTime, _ = cmd.Flags().GetString("st")
	message.RetryTime, _ = cmd.Flags().GetInt("retryTime")

	resp, err := client.Send(context.Background(), message)
	if err != nil {
		er(err)
	}

	cmd.Printf("Credit: %.2f\nSent: %d\nCost: %.2f\nUnsent: %d\nBatchID: %s\n",
		resp.Credit,
		resp.Sent,
		resp.Cost,
		resp.Unsent,
		resp.BatchID,
	)
}
