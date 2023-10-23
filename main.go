package main

import (
	verify "email_verify/sender"
	"fmt"
)

func main() {
	config, err := verify.LoadConfig("./config.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Address)
	fmt.Println(config.Auth.Host)
	var sender verify.VerifyEmailSender

	err = sender.ReadConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	sender.SendTo()
}
