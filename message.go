package every8d

// Message represents an SMS object.
type Message struct {
	// Message title. Empty titles are accepted.
	// The title will not be sent with the SMS; it is just a note.
	Subject string `form:"SB,omitempty"`

	// Message content.
	Content string `form:"MSG"`

	// Receiver's mobile number.
	// Format: +88612345678 or 0912345678
	// Separator: (,) e.g. 0912345678,0922333444
	Destination string `form:"DEST"`

	// Reservation time
	// Send immediately: No input (empty).
	// Reservation send:Please input the reservation time, using this format: yyyyMMddHHmnss, e.g. 20090131153000
	ReservationTime string `form:"ST,omitempty"`

	// SMS validity period of unit: minutes.
	// if not specified, then the platform default validity period is 1440 minutes.
	RetryTime int `form:"RETRYTIME,omitempty"`

	// Message record no.
	MessageNo string `form:"MR,omitempty"`
}
