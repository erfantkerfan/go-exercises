package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`
	RateLimits struct {
		RequestLimit uint `yaml:"ip_requests_per_sec"`
		OTPLimit     uint `yaml:"otp_sms_interval_sec"`
	} `yaml:"rate_limits"`
}

func main() {
	fmt.Println(Read("config.yaml"))
}

func Read(yaml string) (config Config, err error) {
	rawData, err := ioutil.ReadFile(yaml)
	var data Config

	if err != nil {
		return data, err
	}

	if err := data.parseAndValidate(rawData); err != nil {
		return data, err
	} else {
		return data, err
	}
}

func (output *Config) parseAndValidate(data []byte) error {
	if err := yaml.Unmarshal(data, output); err != nil {
		return err
	}

	if output.Server.Host == "" && output.Server.Port == 0 {
		return errors.New("Server field must be present")
	}

	if output.Server.Host == "" {
		output.Server.Host = "localhost"
	}

	if output.Server.Port == 0 {
		return errors.New("Port field must be present")
	}

	if output.Server.Port < 1 || output.Server.Port > 65535 {
		return errors.New("Port field must be a valid port between 1 and 65535")
	}

	if output.RateLimits.RequestLimit == 0 && output.RateLimits.OTPLimit == 0 {
		output.RateLimits.RequestLimit = 100
		output.RateLimits.OTPLimit = 100
	}

	if output.RateLimits.RequestLimit == 0 {
		return errors.New("RequestLimit field must be present")
	}

	if output.RateLimits.RequestLimit <= 60 || output.RateLimits.RequestLimit >= 1000 {
		return errors.New("RequestLimit field must be between >60 and <1000")
	}

	if output.RateLimits.OTPLimit == 0 {
		output.RateLimits.OTPLimit = 100
	}

	if output.RateLimits.OTPLimit < 60 || output.RateLimits.OTPLimit >= 300 {
		return errors.New("OTPLimit field must be between >=60 and <300")
	}

	return nil
}
