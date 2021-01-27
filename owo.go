package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LinkdxTTV/owo/commands"
	"github.com/LinkdxTTV/owo/config"
	"github.com/LinkdxTTV/owo/parse"
)

func main() {
	args := os.Args

	// Config
	cfg, err := config.ParseConfig("config.json")
	if err != nil {
		cfg, err = config.ParseConfig("default.json")
		if err != nil {
			log.Fatal(err)
		}
	}
	// First time?
	if !cfg.Initialized {
		cfg = commands.FirstTimeSetup(cfg)
		err := config.UpdateConfig(cfg)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	// No args
	if len(args) == 1 {
		showBaseMessage(cfg)
		os.Exit(0)
	}

	// Commands and filestructure
	if commands.IsBaseCommand(args) { // Base Command
		err = commands.HandleBaseCommand(cfg, args)
		if err != nil {
			log.Fatal(err)
		}
	} else if commands.IsFileCommand(args) {
		err = commands.HandleFileCommands(cfg, args)
		if err != nil {
			log.Fatal(err)
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

func showBaseMessage(cfg *config.Config) {
	fmt.Println("owo: command line knowledge source")
	fmt.Println("----------------------------------")
	fmt.Println("  owo -[h]elp || -[a]bout || -[ch]eckup || -[u]pdate || -[c]onfig")
	fmt.Println()
	parse.NavigateAndShowDir(cfg.LocalPath+"/docs/docs", cfg)
}
