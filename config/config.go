package config

import (
	"encoding/json"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
)

// Config defines a config
type Config struct {
	Git struct {
		RemoteURL string `json:"remoteURL"`
		SSHURL    string `json:"SSHURL"`
		SHA       string `json:"sha"`
		LocalPath string `json:"localpath"`
	} `json:"git"`
}

// ParseConfig parses the config file
func ParseConfig() (*Config, error) {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	gopath += "/src/github.com/LinkdxTTV/owo/config/config.json"
	fmt.Println(gopath)

	f, err := os.Open(gopath)
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

func UpdateConfig(cfg *Config) error {
	byteJSON, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	gopath += "/src/github.com/LinkdxTTV/owo/config/config.json"

	err = ioutil.WriteFile(gopath, byteJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}
