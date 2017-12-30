package app

import (
	"context"
	"encoding/base64"
	"io/ioutil"

	"github.com/minchao/go-every8d"
	"github.com/spf13/cobra"
)

var (
	sendMMSCmd = &cobra.Command{
		Use:   "send-mms",
		Short: "Send an MMS",
		Run:   sendMMSFunc,
	}
)

func init() {
	sendMMSCmd.Flags().StringP("sb", "s", "", "Message title. Empty titles are accepted")
	sendMMSCmd.Flags().StringP("msg", "m", "", "Message content")
	sendMMSCmd.Flags().StringP("dest", "d", "", "Receiver's mobile number")
	sendMMSCmd.Flags().StringP("st", "R", "", "Reservation time")
	sendMMSCmd.Flags().IntP("retryTime", "r", 0, "SMS validity period of unit: minutes")
	sendMMSCmd.Flags().StringP("image", "i", "", "Image file, binary base64 encoded")
	sendMMSCmd.Flags().StringP("attachment", "a", "", "Image file path")
	sendMMSCmd.Flags().StringP("type", "t", "", "Image file extension, support jpg/jpeg/png/git")
}

func sendMMSFunc(cmd *cobra.Command, _ []string) {
	message := every8d.MMS{}
	message.Subject, _ = cmd.Flags().GetString("sb")
	message.Content, _ = cmd.Flags().GetString("msg")
	message.Destination, _ = cmd.Flags().GetString("dest")
	message.ReservationTime, _ = cmd.Flags().GetString("st")
	message.RetryTime, _ = cmd.Flags().GetInt("retryTime")
	message.Attachment, _ = cmd.Flags().GetString("image")
	message.Type, _ = cmd.Flags().GetString("type")

	if attachment, _ := cmd.Flags().GetString("attachment"); attachment != "" {
		f, err := ioutil.ReadFile(attachment)
		if err != nil {
			er(err)
		}
		message.Attachment = base64.StdEncoding.EncodeToString(f)
	}

	resp, err := client.SendMMS(context.Background(), message)
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
