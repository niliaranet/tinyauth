package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niliaranet/tinyauth/pkg/totp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid secret code!")
		return
	}

	secretString := os.Args[1]
	secretB32 := []byte(secretString)

	code, err := totp.GenerateCurrentTOTP(secretB32)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(code)
}
