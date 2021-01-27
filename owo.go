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
		fmt.Println("owo: command line knowledge source")
		fmt.Println("----------------------------------")
		fmt.Println("  owo -[h]elp || -[a]bout || -[ch]eckup || -[u]pdate || -[c]onfig")
		fmt.Println()
		parse.NavigateAndShowDir(cfg.LocalPath+"/docs/docs", cfg)
		os.Exit(0)
	}

	if commands.IsBaseCommand(args) { // Base Command
		commands.HandleBaseCommand(cfg, args)
	} else if commands.IsFileCommand(args) {
		commands.HandleFileCommands(cfg, args)
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
