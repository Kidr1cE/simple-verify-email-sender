package main

import "email_verify/verify"

func main() {
	config, err := verify.LoadEmailConfig("./config.yml")
	if err != nil {
		return
	}
	email := verify.NewVerifyEmail(config)
	email.Send([]string{"2890034671@qq.com"})
}
