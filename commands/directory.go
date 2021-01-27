package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

const (
	Directory     string = "dir"
	DirectoryUnix string = "mkdir"
)

func createNewDirectory(cfg *config.Config, args []string) error {
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
	dirName := args[len(args)-1]
	fullPathFile := fullPathDir + "/" + dirName

	// Check to see if the file already exists
	_, err := ioutil.ReadFile(fullPathFile)
	if err == nil {
		fmt.Println("Error:", dirName, "already exists as file")
		return nil
	}
	// Check if its a directory
	if err != nil {
		if strings.Contains(err.Error(), "is a directory") {
			fmt.Println("Error:", dirName, "already is a directory")
		}
	}
	// Ok lets create it
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)
	err = os.Chdir(basePath)
	if err != nil {
		return err
	}

	err = os.Mkdir(dirName, 0755)
	if err != nil {
		return err
	}
	fmt.Println(dirName, "created succesfully")

	return nil
}
