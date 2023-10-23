package sender_test

import (
	"email_verify/sender"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig(t *testing.T) {
	assert := assert.New(t)
	config, err := sender.LoadConfig("../config.yml")
	assert.Nil(err)

	var sender sender.VerifyEmailSender
	err = sender.ReadConfig(config)
	assert.Nil(err)
}
