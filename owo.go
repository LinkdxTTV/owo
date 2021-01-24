package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LinkdxTTV/owo/commands"
	"github.com/LinkdxTTV/owo/config"
	"github.com/LinkdxTTV/owo/parse"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("owo: command line knowledge source 1")
		fmt.Println("------------------------------------")
		fmt.Println("  owo -about")
		fmt.Println("  owo -checkup")
		fmt.Println("  owo -update")
		os.Exit(0)
	}

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(args[1], "-") { // Command
		switch args[1] {
		case commands.Checkup:
			needsUpdate, err := commands.CheckForUpdate(cfg)
			if err != nil {
				log.Fatal(err)
			}
			if !needsUpdate {
				fmt.Println("owo you're up to date :)")
			} else {
				fmt.Println("Please run: owo update")
			}
		case commands.About:
			commands.ShowAbout()
		case commands.Update:
			err := commands.CmdUpdate(cfg)
			if err != nil {
				log.Fatal(err)
			}
		// case "test":

		// 	entry, err := parse.ParseEntry(cfg.LocalPath + "/docs/text/testfile")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	entry.Print()
		// case "nav":
		// 	parse.NavigateAndShow([]string{"testfile"}, cfg)
		default:
			fmt.Println("command", args[1], "not recognized")
		}
	} else {
		// Likely a directory or whatnot
		err := parse.NavigateAndShow(args[1:], cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
	return
}
