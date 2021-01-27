package commands

import (
	"fmt"

	"github.com/LinkdxTTV/owo/config"
)

func IsFileCommand(args []string) bool {
	if len(args) < 2 {
		return false
	}
	if string(args[len(args)-2][0]) == "-" {
		return true
	}
	lastArg := string(args[len(args)-1][0])
	if lastArg == "-" {
		showFileCommandHelp(lastArg)
	}
	return false
}

func showFileCommandHelp(command string) {
	fmt.Println("Error: No target for file command:", command)
	fmt.Println("Generally file commands have a directory before them, and a target file after")
	fmt.Println("  example: owo folder1 subfolder2 -new newFileName")
}

func HandleFileCommands(cfg *config.Config, args []string) error {
	if len(args) < 2 {
		return nil
	}
	command := sanitizeCommand(args[len(args)-2])
	switch command {
	case Touch, TouchShort, New, NewShort:
		err := createNewFile(cfg, args)
		if err != nil {
			return err
		}
	default:
		fmt.Println("Unrecognized file command: ", command)
	}

	return nil
}
