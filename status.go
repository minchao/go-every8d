package every8d

// StatusCode of EVERY8D API.
type StatusCode int

// List of EVERY8D API status codes.
const (
	StatusContentIsEmpty                 = StatusCode(-24)
	StatusWrongUsername                  = StatusCode(-100)
	StatusWrongPassword                  = StatusCode(-101)
	StatusUsernameAndPasswordAreRequired = StatusCode(-300)
)
