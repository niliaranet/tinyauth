package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"time"
)

func GenerateTOTP(secret string) any {
	counter := uint(time.Now().Unix() / 30)

	counterByte := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		counterByte[i] = byte(counter & 0xff)
		counter >>= 8
	}

	secretByte, _ := base32.StdEncoding.DecodeString(secret)
	hash := hmac.New(sha1.New, secretByte)
	_, err := hash.Write(counterByte)
	if err != nil {
		panic(err)
	}
	hmacBytes := hash.Sum(nil)

	offset := hmacBytes[len(hmacBytes)-1] & 0xf

	code := uint(hmacBytes[offset]&0x7f)<<24 |
		(uint(hmacBytes[offset+1])&0xff)<<16 |
		(uint(hmacBytes[offset+2])&0xff)<<8 |
		(uint(hmacBytes[offset+3]) & 0xff)

	code = code % 1000000

	return code
}
