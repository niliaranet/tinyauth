package server

import (
	"github.com/niliaranet/tinyauth/totp"
	totpCounter "github.com/niliaranet/tinyauth/totp/counter"
)

func ValidateKey(inputKey uint, secret []byte, margin uint) (bool, error) {
	currentCount := totpCounter.GetCurrent()
	counters := []uint{totpCounter.GetCurrent()}

	for i := uint(1); i <= margin; i++ {
		counters = append(counters, currentCount+i)
		counters = append(counters, currentCount-i)
	}

	for _, counter := range counters {
		matches, err := validateSingle(inputKey, secret, counter)
		if err != nil {
			return false, err
		}

		if matches {
			return true, nil
		}
	}

	return false, nil
}

func validateSingle(inputKey uint, secret []byte, count uint) (bool, error) {
	real, err := totp.GenerateKey(secret, count)
	if err != nil {
		return false, err
	}

	if real == inputKey {
		return true, err
	}

	return false, nil
}
