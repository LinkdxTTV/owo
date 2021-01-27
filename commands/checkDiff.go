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

func DeferDiffCheck(cfg *config.Config) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)
	err = os.Chdir(cfg.LocalPath)
	if err != nil {
		return err
	}

	// Inside Go pipe shenanigans
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
	fmt.Println("You have", numFilesChanged, "unsynced change(s). owo -diff to see. Perhaps you want to owo -[s]ync or owo -[r]eset?")

	return nil
}

func checkDiff(cfg *config.Config) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)
	err = os.Chdir(cfg.LocalPath + "/docs/docs")
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "status", "--short")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
