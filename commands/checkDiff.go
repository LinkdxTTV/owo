package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

const (
	Diff string = "diff"
)

func CheckDiff(cfg *config.Config) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)
	err = os.Chdir(cfg.LocalPath)
	if err != nil {
		return err
	}

	c1 := exec.Command("git", "status", "--short")
	c2 := exec.Command("wc", "-l")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	err = c1.Start()
	if err != nil {
		return err
	}
	err = c2.Start()
	if err != nil {
		return err
	}
	err = c1.Wait()
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	err = c2.Wait()
	if err != nil {
		return err
	}

	filesChanged := strings.TrimSpace(string(b2.Bytes()))
	numFilesChanged, err := strconv.Atoi(filesChanged)
	if err != nil {
		return err
	}

	if numFilesChanged == 0 {
		return nil
	}
	fmt.Println()
	fmt.Println("You have", numFilesChanged, "unsynced change(s). Perhaps you want to owo -sync or owo -reset ?")

	return nil
}
