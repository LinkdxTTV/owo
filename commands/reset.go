package commands

import (
	"os"
	"os/exec"

	"github.com/LinkdxTTV/owo/config"
)

const (
	Reset      string = "reset"
	ResetShort string = "r"
)

func reset(cfg *config.Config) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)
	err = os.Chdir(cfg.LocalPath)
	if err != nil {
		return err
	}
	return exec.Command("git", "reset", "--hard").Run()
}
