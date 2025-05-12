package main

import (
	"fmt"
	"os"

	"github.com/niliaranet/tinyauth/pkg/totp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid secret code!")
		return
	}

	secret := os.Args[1]
	fmt.Println(secret)
	code := totp.GenerateTOTP(secret)
	fmt.Println(code)
}
