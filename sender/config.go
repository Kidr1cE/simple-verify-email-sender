package sender

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Auth struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Verify struct {
	Type string `yaml:"type"`
}

type Template struct {
	Name    string `yaml:"name"`
	Subject string `yaml:"subject"`
}

type VerifyEmailConfig struct {
	Address  string    `yaml:"address"`
	Auth     *Auth     `yaml:"auth"`
	Verify   *Verify   `yaml:"verify"`
	Template *Template `yaml:"template"`
}

func LoadConfig(filename string) (*VerifyEmailConfig, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config VerifyEmailConfig
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
