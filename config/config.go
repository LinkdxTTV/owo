package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config defines a config
type Config struct {
	Git struct {
		RemoteURL string `json:"remoteURL"`
		SSHURL    string `json:"SSHURL"`
		SHA       string `json:"sha"`
	} `json:"git"`
}

// ParseConfig parses the config file
func ParseConfig() (*Config, error) {
	f, err := os.Open("./config/config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func UpdateConfigSHA(cfg *Config) error {
	byteJSON, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./config/config.json", byteJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}
