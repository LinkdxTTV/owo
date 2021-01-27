package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/LinkdxTTV/owo/config"
)

const Update string = "update"

func CmdUpdate(cfg *config.Config) error {

	needsUpdate, err := CmdCheckup(cfg)
	if err != nil {
		return err
	}

	if !needsUpdate {
		fmt.Println("owo you are already up to date :)")
		return nil
	}

	os.Setenv("GO111MODULE", "off")
	updateCmd := exec.Command("go", "get", "-u", cfg.Git.RemoteURL)
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stdout

	err = updateCmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	newSHA, err := getNewestSHA(cfg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cfg.Git.SHA = newSHA
	config.UpdateConfig(cfg)
	fmt.Println("Update Successful. You are ready to owo :)")

	return nil
}
