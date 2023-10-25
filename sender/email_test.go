package sender_test

import (
	"email_verify/sender"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Template(t *testing.T) {
	assert := assert.New(t)
	file, err := os.ReadFile("../config.yml")
	config, err := sender.LoadConfig(file)
	assert.Nil(err)

	var sender sender.VerifyEmailSender
	err = sender.ReadConfig(config)
	assert.Nil(err)
	log.Panicln(err)

}
