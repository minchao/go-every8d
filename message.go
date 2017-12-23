package every8d

// Message represents an SMS object.
type Message struct {
	// Message title. Empty titles are accepted.
	// The title will not be sent with the SMS; it is just a note.
	Subject string `url:"SB,omitempty"`

	// Message content.
	Content string `url:"MSG"`

	// Receiver's mobile number.
	// Format: +88612345678 or 0912345678
	// Separator: (,) e.g. 0912345678,0922333444
	Destination string `url:"DEST"`

	// Reservation time
	// Send immediately: No input (empty).
	// Reservation send:Please input the reservation time, using this format: yyyyMMddHHmnss, e.g. 20090131153000
	ReservationTime string `url:"ST,omitempty"`

	// SMS validity period of unit: minutes.
	// if not specified, then the platform default validity period is 1440 minutes.
	RetryTime int `url:"RETRYTIME,omitempty"`

	// Message record no.
	MessageNo string `url:"MR,omitempty"`

	// Callback URL to receive the delivery status or reply report.
	StatusReportURL string `url:"StatusReportURL,omitempty"`
}
