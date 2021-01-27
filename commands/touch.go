package commands

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

const (
	Touch      string = "touch"
	TouchShort string = "t"
	New        string = "new"
	NewShort   string = "n"
)

func createNewFile(cfg *config.Config, args []string) error {
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

	// Check to see if the file already exists
	_, err := ioutil.ReadFile(fullPathFile)
	if err == nil {
		fmt.Println("Error: file already exists:", filename)
		fmt.Println("Perhaps you want to -edit")
		return nil
	}
	// Check if its a directory
	if err != nil {
		if strings.Contains(err.Error(), "is a directory") {
			fmt.Println("Error: Cannot use filename", filename, "as it is a directory")
		}
	}
	// Ok lets create it
	filecontents := []byte{}
	if filename != "meta" {
		filecontents = []byte(newFileTemplate)
	}

	err = ioutil.WriteFile(fullPathFile, filecontents, 0644)
	if err != nil {
		return err
	}
	fmt.Println(filename, "created succesfully")

	return nil
}

func showTouchHelp() {
	fmt.Println("Error: command syntax not recognized")
	fmt.Println()
	fmt.Println("-touch or -new expects a folder path before it, and a newfile name after it")
	fmt.Println("  example: owo directory1 subdirectory2 -touch newFileName")
}

func checkNewFileSyntax(args []string) {
	// Assumes owo is the 0th entry in args

}

const newFileTemplate string = `Title: 

Command: 

BodyStart:

BodyEnd:

Notes:

`
