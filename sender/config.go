package sender

import (
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

func LoadConfig(file []byte) (*VerifyEmailConfig, error) {
	var config VerifyEmailConfig
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
