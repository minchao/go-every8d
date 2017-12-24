package app

import (
	"context"

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
}

func deliveryStatusFunc(cmd *cobra.Command, _ []string) {
	batchID, _ := cmd.Flags().GetString("bid")
	pageNo, _ := cmd.Flags().GetString("pno")

	resp, err := client.GetDeliveryStatus(context.Background(), batchID, pageNo)
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
