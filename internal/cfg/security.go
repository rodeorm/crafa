package cfg

import "time"

type SecurityConfig struct {
	TokenLiveTime time.Duration
	OTPLiveTime   time.Duration
	JWTKey        string
}
