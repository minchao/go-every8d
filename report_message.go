package every8d

import (
	"net/http"
	"strconv"
)

// ReportMessage represents sending report or reply message.
type ReportMessage struct {
	// Batch ID.
	BatchID string `url:"BATCHID"`

	// Receive's mobile number.
	// Format: +88612345678
	Destination string `url:"RM"`

	// Report time.
	ReportTime string `url:"RT"`

	// Sending status.
	StatusCode StatusCode `url:"STATUS"`

	// Reply message.
	ReplyMessage string `url:"SM,omitempty"`

	// Message record no.
	MessageNo string `url:"MR,omitempty"`
}

// ParseReportMessage parses an incoming EVERY8D callback request and return the ReportMessage.
func ParseReportMessage(r *http.Request) (*ReportMessage, error) {
	values := r.URL.Query()

	code, _ := strconv.Atoi(values.Get("STATUS"))

	reportMessage := new(ReportMessage)
	reportMessage.BatchID = values.Get("BATCHID")
	reportMessage.Destination = values.Get("RM")
	reportMessage.ReportTime = values.Get("RT")
	reportMessage.StatusCode = StatusCode(code)
	reportMessage.ReplyMessage = values.Get("SM")
	reportMessage.MessageNo = values.Get("MR")

	return reportMessage, nil
}
