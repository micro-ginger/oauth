package validator

import "time"

type Config struct {
	// KeyPrefix is prefix of key going to be saved into database
	KeyPrefix string
	// RequestRetryLimitDuration is minimum time needs to pass after last request
	RequestRetryLimitDuration time.Duration
	// RequestExpiration duration limit to verify after request
	RequestExpiration time.Duration

	Validators []struct {
		HistoryCycleDuration time.Duration
		MaximumRequestsCount int
		MaximumVerifiesCount int
	}
}

func (c *Config) Initialize() {
	for _, c := range c.Validators {
		if c.HistoryCycleDuration == 0 {
			panic("HistoryCycleDuration is 0")
		}
	}
}
