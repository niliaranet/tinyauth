package counter

import "time"

func GetCurrent() uint {
	return uint(time.Now().Unix() / 30)
}
