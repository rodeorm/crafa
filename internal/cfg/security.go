package cfg

import "time"

type SecurityConfig struct {
	TokeLiveTime time.Duration
	JWTKey       string
}
