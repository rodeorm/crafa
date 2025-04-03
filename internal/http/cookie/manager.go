package cookie

import "time"

type CookieManager struct {
	key      string
	liveTime time.Duration // Время жизни в миллисекундах
}

func NewCookieManager(key string, sessionLiveTime time.Duration) *CookieManager {
	return &CookieManager{key: key, liveTime: sessionLiveTime}
}
