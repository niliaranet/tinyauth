package counter

import "time"

func GetCurrentCounter() uint {
	return uint(time.Now().Unix() / 30)
}
