package totp

import (
	"crypto/hmac"
	"crypto/sha1"

	"github.com/niliaranet/tinyauth/pkg/totp/counter"
)

func GenerateCurrentTOTP(secret []byte) (uint, error) {
	return GenerateTOTP(secret, counter.GetCurrentCounter())
}

func GenerateTOTP(secret []byte, counter uint) (uint, error) {
	counterByte := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		counterByte[i] = byte(counter & 0xff)
		counter >>= 8
	}

	hash := hmac.New(sha1.New, secret)
	if _, err := hash.Write(counterByte); err != nil {
		return 0, err
	}

	hmacBytes := hash.Sum(nil)
	offset := hmacBytes[len(hmacBytes)-1] & 0xf

	code := uint(hmacBytes[offset]&0x7f)<<24 |
		(uint(hmacBytes[offset+1])&0xff)<<16 |
		(uint(hmacBytes[offset+2])&0xff)<<8 |
		(uint(hmacBytes[offset+3]) & 0xff)

	code %= 1000000

	return code, nil
}
