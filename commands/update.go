package commands

import (
	"fmt"
	"os"
	"os/exec"
	"owo/config"
)

const Update string = "update"

func CmdUpdate() error {
	cfg, err := config.ParseConfig()
	if err != nil {
		return err
	}
	updateCmd := exec.Command("go", "get", "-u", cfg.Git.RemoteURL)
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stdout

	err = updateCmd.Run()
	if err != nil {
		return err
	}

	newSHA, err := getNewestSHA(cfg)
	if err != nil {
		return err
	}
	cfg.Git.SHA = newSHA
	config.UpdateConfigSHA(cfg)
	fmt.Println("Update Succesful. You're ready to owo :)")

	return nil
}
