package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

const (
	Delete      string = "delete"
	DeleteShort string = "d"
	Remove      string = "remove"
	RemoveShort string = "rm"
)

func deleteFile(cfg *config.Config, args []string) error {
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
			fmt.Println("Error: Cannot use", filename, "as it is a directory")
		}
		if strings.Contains(err.Error(), "no such") {
			fmt.Println("Error: No such file", filename, "exists")
		}
		return nil
	}
	// Ok lets delete it

	err = os.Remove(fullPathFile)

	if err != nil {
		return err
	}
	fmt.Println(filename, "deleted succesfully")

	return nil
}

func showDeleteHelp() {
	fmt.Println("Error: command syntax not recognized")
	fmt.Println()
	fmt.Println("-delete or -remove expects a folder path before it, and a file name after it")
	fmt.Println("  example: owo directory1 subdirectory2 -delete newFileName")
}
