package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/LinkdxTTV/owo/config"
	"github.com/LinkdxTTV/owo/parse"
)

const (
	Edit      string = "edit"
	EditShort string = "e"
)

func editFile(cfg *config.Config, args []string) error {
	// Ensure syntax is correct
	basePath := cfg.LocalPath + "/docs/docs"
	args = args[1:] // Remove "owo"

	argPath := ""
	for idx, subdir := range args {
		if idx == len(args)-2 || idx == len(args)-1 { // skip the command and filename
			continue
		}
		argPath += "/" + subdir
	}
	fullPathDir := basePath + argPath
	filename := args[len(args)-1]
	fullPathFile := fullPathDir + "/" + filename

	// Check to see if the file exists
	_, err := ioutil.ReadFile(fullPathFile)
	// Check if its a directory
	if err != nil {
		if strings.Contains(err.Error(), "is a directory") {
			fmt.Println("Error: Cannot edit", filename, "as it is a directory")
		}
		if strings.Contains(err.Error(), "no such") {
			fmt.Println("Error: No such file", filename, "exists. Showing directory:")
			parse.NavigateAndShowDir(fullPathDir, cfg)
		}
		return nil
	}
	// Ok lets edit it

	editCmd := exec.Command(cfg.PreferredEditor, fullPathFile)
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stdout
	editCmd.Stdin = os.Stdin

	err = editCmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("editing file", filename, "in preferred editor:", cfg.PreferredEditor)

	return nil
}
