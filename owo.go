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

	cfg, err := config.ParseConfig("config.json")
	if err != nil {
		cfg, err = config.ParseConfig("default.json")
		if err != nil {
			log.Fatal(err)
		}
	}
	if !cfg.Initialized {
		cfg = commands.FirstTimeSetup(cfg)
		err := config.UpdateConfig(cfg)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if len(args) == 1 {
		fmt.Println("owo: command line knowledge source 1")
		fmt.Println("------------------------------------")
		fmt.Println("  owo -help || -about || -checkup || -update || -config\n")
		parse.NavigateAndShowDir(cfg.LocalPath+"/docs/docs", cfg)
		os.Exit(0)
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
		case commands.Config:
			err := commands.CmdConfig(cfg)
			if err != nil {
				log.Fatal(err)
			}
		case commands.Help:
			commands.CmdHelp()
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
