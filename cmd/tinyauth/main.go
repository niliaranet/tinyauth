package main

import (
	"os"

	"github.com/niliaranet/tinyauth/totp"
)

func main() {
	if len(os.Args) < 2 {
		println("Please provide a valid secret code!")
		return
	}

	secretString := os.Args[1]
	secretB32 := []byte(secretString)

	code, err := totp.GenerateFromLocalTime(secretB32)
	if err != nil {
		println(err)
		return
	}

	println(code)
}
