package every8d

// StatusCode of EVERY8D API.
type StatusCode int

// List of EVERY8D API status codes.
const (
	StatusInvalidMobileNumber                  = StatusCode(-3)
	StatusDTFormatErrorOrPassedMoreThan24Hours = StatusCode(-4)
	StatusTheContentIsEmpt                     = StatusCode(-24)
	StatusNoMobile                             = StatusCode(-41)
	StatusWrongUsername                        = StatusCode(-100)
	StatusWrongPassword                        = StatusCode(-101)
	StatusUsernameAndPasswordAreRequired       = StatusCode(-300)
	StatusSent                                 = StatusCode(0)
	StatusMessageReceived                      = StatusCode(100)
	StatusDeliveryFailureDueMobile             = StatusCode(101)
	StatusDeliveryFailureDueTelecom102         = StatusCode(102)
	StatusMobileNumberNotExist                 = StatusCode(103)
	StatusDeliveryFailureDueTelecom104         = StatusCode(104)
	StatusDeliveryFailureDueTelecom105         = StatusCode(105)
	StatusDeliveryFailureDueTelecom106         = StatusCode(106)
	StatusReceivedAfterDeadline                = StatusCode(107)
	StatusReservationSMS                       = StatusCode(300)
	StatusNoCredit                             = StatusCode(301)
	StatusCanceled                             = StatusCode(303)
	StatusInternationalSMSNotConfigured        = StatusCode(500)
	StatusSMSSent                              = StatusCode(700)
	StatusTestingMode                          = StatusCode(888)
	StatusReplayContent                        = StatusCode(999)
)
