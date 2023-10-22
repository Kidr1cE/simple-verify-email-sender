package verify

import (
	"os"

	"gopkg.in/yaml.v2"
)

type EmailConfig struct {
	VerifyEmail VerifyEmailConfig `yaml:"verify-email"`
}

type VerifyEmailConfig struct {
	Auth    AuthConfig    `yaml:"auth"`
	Content ContentConfig `yaml:"content"`
}

type AuthConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ContentConfig struct {
	Name    string `yaml:"name"`
	Subject string `yaml:"subject"`
	Path    string `yaml:"path"`
}

func LoadEmailConfig(filename string) (*EmailConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config EmailConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
