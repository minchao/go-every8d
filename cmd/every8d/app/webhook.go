package app

import (
	"fmt"
	"net/http"

	"github.com/minchao/go-every8d"
	"github.com/spf13/cobra"
)

var (
	webhookCmd = &cobra.Command{
		Use:   "webhook",
		Short: "Webhook to receive the sending report and reply message",
		Run:   webhookFunc,
	}
)

func init() {
	webhookCmd.Flags().IntP("port", "p", 8080, "HTTP Server Port")
}

func webhookFunc(cmd *cobra.Command, _ []string) {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		report, err := every8d.ParseReportMessage(r)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		}

		cmd.Printf("%s\t%s\t%s\t%d\t%s\t%s\t%s\n",
			report.BatchID,
			report.Destination,
			report.ReportTime,
			report.StatusCode,
			report.StatusCode.Text(),
			report.ReplyMessage,
			report.MessageNo,
		)
	})

	port, _ := cmd.Flags().GetInt("port")

	cmd.Printf("Starting HTTP server on :%d\n", port)
	cmd.Println("BatchID\tRM\tRT\tSTATUS\tSTATUS_TEXT\tSM\tMR\t")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		cmd.Printf("ListenAndServe error: %v", err)
	}
}
