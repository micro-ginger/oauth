package verify

const otpType = "MOBILE_OTP"

type config struct {
	Debug                   bool
	NotificationMessageType string
}

func (c *config) initialize() {
	if c.NotificationMessageType == "" {
		c.NotificationMessageType = "LOGIN_MOBILE_VERIFY"
	}
}
