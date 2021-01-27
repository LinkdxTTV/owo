package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

func SanitizeCommand(command string) string {
	// Get rid of weird capitals
	command = strings.ToLower(command)
	// Allow one or two dashes
	if len(command) < 2 {
		return command
	}
	if command[0:2] == "--" {
		return command[2:]
	}
	if command[0:1] == "-" {
		return command[1:]
	}
	return command
}

func HandleCommand(cfg *config.Config, command string) {
	switch SanitizeCommand(command) {
	case Checkup:
		needsUpdate, err := CmdCheckup(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if !needsUpdate {
			fmt.Println("owo you're up to date :)")
		} else {
			fmt.Println("Please run: owo -update")
		}
	case About:
		CmdAbout()
	case Update:
		err := CmdUpdate(cfg)
		if err != nil {
			log.Fatal(err)
		}
	case Config:
		err := CmdConfig(cfg)
		if err != nil {
			log.Fatal(err)
		}
	case Help:
		CmdHelp()
	default:
		fmt.Println("command", command, "not recognized")
	}
}
