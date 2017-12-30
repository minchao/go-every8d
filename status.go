package every8d

// StatusCode of EVERY8D API.
type StatusCode int

// List of EVERY8D API status codes.
const (
	StatusInvalidMobileNumber                  = StatusCode(-3)
	StatusDTFormatErrorOrPassedMoreThan24Hours = StatusCode(-4)
	StatusTheContentIsEmpty                    = StatusCode(-24)
	StatusNoMobile                             = StatusCode(-41)
	StatusServerSiteError                      = StatusCode(-99)
	StatusWrongUsername                        = StatusCode(-100)
	StatusWrongPassword                        = StatusCode(-101)
	StatusUsernameAndPasswordAreRequired       = StatusCode(-300)
	StatusSubjectRequired                      = StatusCode(-303)
	StatusImageRequired                        = StatusCode(-304)
	StatusImageTypeRequired                    = StatusCode(-305)
	StatusImageTooLarge                        = StatusCode(-314)
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

var statusText = map[StatusCode]string{
	StatusInvalidMobileNumber:                  "無效門號",
	StatusDTFormatErrorOrPassedMoreThan24Hours: "DT 格式錯誤或預計發送時間已過去 24小時以上",
	StatusTheContentIsEmpty:                    "The content is empty.",
	StatusNoMobile:                             "no mobile.",
	StatusServerSiteError:                      "主機端發生不明錯誤，請與廠商窗口聯繫。",
	StatusWrongUsername:                        "無此帳號。",
	StatusWrongPassword:                        "密碼錯誤。",
	StatusUsernameAndPasswordAreRequired:       "帳號密碼不得為空值。",
	StatusSubjectRequired:                      "主旨不得為空",
	StatusImageRequired:                        "未上傳圖片",
	StatusImageTypeRequired:                    "檔案副檔名類型不得為空",
	StatusImageTooLarge:                        "圖文簡訊大小超過50K",
	StatusSent:                                 "已發送",
	StatusMessageReceived:                      "發送成功",
	StatusDeliveryFailureDueMobile:             "手機端因素未能送達",
	StatusDeliveryFailureDueTelecom102:         "電信終端設備異常未能送達",
	StatusMobileNumberNotExist:                 "無此手機號碼",
	StatusDeliveryFailureDueTelecom104:         "電信終端設備異常未能送達",
	StatusDeliveryFailureDueTelecom105:         "電信終端設備異常未能送達",
	StatusDeliveryFailureDueTelecom106:         "電信終端設備異常未能送達",
	StatusReceivedAfterDeadline:                "逾時收訊",
	StatusReservationSMS:                       "預約簡訊",
	StatusNoCredit:                             "無額度(或額度不足)無法發送",
	StatusCanceled:                             "取消簡訊",
	StatusInternationalSMSNotConfigured:        "未開通國際簡訊，無法發送國際簡訊",
	StatusSMSSent:                              "MMS 已發送",
	StatusTestingMode:                          "測試模式",
	StatusReplayContent:                        "代表此呼叫為回覆簡訊之內容",
}

// Text returns status code text.
func (c StatusCode) Text() string {
	if str, ok := statusText[c]; ok {
		return str
	}
	return "Unknown StatusCode"
}
