package main

import (
	"fmt"
	"go/build"
	"log"
	"os"

	"github.com/LinkdxTTV/owo/commands"
	"github.com/LinkdxTTV/owo/config"
	"github.com/LinkdxTTV/owo/parse"
)

func main() {
	args := os.Args
	// fmt.Println(args)

	if len(args) == 1 {
		fmt.Println("owo: command line knowledge source 5")
		fmt.Println("--------------------------------")
		fmt.Println("  owo about")
		fmt.Println("  owo checkup")
		fmt.Println("  owo update")
		os.Exit(0)
	}

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	switch args[1] {
	case commands.Checkup:
		commands.CheckForUpdate(cfg)
	case commands.About:
		commands.ShowAbout()
	case commands.Update:
		err := commands.CmdUpdate(cfg)
		if err != nil {
			log.Fatal(err)
		}
	case "test":
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}

		gopath += "/src/github.com/LinkdxTTV/owo/docs/text/testfile"
		entry, err := parse.ParseEntry(gopath)
		if err != nil {
			log.Fatal(err)
		}
		entry.Print()
	}

	return
}
