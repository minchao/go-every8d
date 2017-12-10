package every8d

// StatusCode of EVERY8D API.
type StatusCode int

// List of EVERY8D API status codes.
const (
	StatusTheContentIsEmpt               = StatusCode(-24)
	StatusNoMobile                       = StatusCode(-41)
	StatusWrongUsername                  = StatusCode(-100)
	StatusWrongPassword                  = StatusCode(-101)
	StatusUsernameAndPasswordAreRequired = StatusCode(-300)
)
