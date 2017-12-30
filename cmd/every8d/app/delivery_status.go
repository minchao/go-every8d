package app

import (
	"context"

	"github.com/minchao/go-every8d"
	"github.com/spf13/cobra"
)

var (
	deliveryStatusCmd = &cobra.Command{
		Use:   "delivery-status",
		Short: "Query to retrieve the delivery status",
		Run:   deliveryStatusFunc,
	}
)

func init() {
	deliveryStatusCmd.Flags().StringP("bid", "b", "", "Batch ID")
	deliveryStatusCmd.Flags().StringP("pno", "p", "", "Paging number")
	deliveryStatusCmd.Flags().StringP("type", "t", "sms", "Message type (\"sms\"|\"mms\")")
}

func deliveryStatusFunc(cmd *cobra.Command, _ []string) {
	batchID, _ := cmd.Flags().GetString("bid")
	pageNo, _ := cmd.Flags().GetString("pno")
	messageType, _ := cmd.Flags().GetString("type")

	var resp *every8d.DeliveryStatusResponse
	var err error
	if messageType == "mms" {
		resp, err = client.GetMMSDeliveryStatus(context.Background(), batchID, pageNo)
	} else {
		resp, err = client.GetDeliveryStatus(context.Background(), batchID, pageNo)
	}
	if err != nil {
		er(err)
	}

	cmd.Printf("Count: %d\n", resp.Count)
	cmd.Println("Name\tMobile\tSendTime\tCost\tStatus\tStatusText")
	for _, record := range resp.Records {
		cmd.Printf("%s\t%s\t%s\t%.2f\t%v\t%s\n",
			record.Name,
			record.Mobile,
			record.SendTime,
			record.Cost,
			record.Status,
			record.Status.Text(),
		)
	}
}
