package main

import (
	verify "email_verify/sender"
	"fmt"
	"os"
)

func main() {
	// read config
	file, err := os.ReadFile("./config.yml")
	config, err := verify.LoadConfig(file)

	// inint email sender
	var sender verify.VerifyEmailSender
	err = sender.ReadConfig(config)
	if err != nil {
		fmt.Println(err)
	}

	// send email
	sender.SendTo("784312513@qq.com", "https://yourwebsite/verity/?token=114514")
}
