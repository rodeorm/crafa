package cookie

import "time"

type CookieManager struct {
	key      string
	liveTime time.Duration // Время жизни в миллисекундах
}
