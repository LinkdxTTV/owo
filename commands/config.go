package commands

import (
	"bufio"
	"fmt"
	"go/build"
	"log"
	"os"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

const Config string = "config"

func CmdConfig(cfg *config.Config) error {
	cfg = FirstTimeSetup(cfg)
	err := config.UpdateConfig(cfg)
	if err != nil {
		return err
	}
	return nil
}

func FirstTimeSetup(cfg *config.Config) *config.Config {

	fmt.Println("Woah this is your FIRST TIME owo, I need to ask you a few questions!")

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	localPath := gopath + "/src/" + cfg.Git.RemoteURL
	cfg.LocalPath = localPath

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your preferred editor? ex. vim, emacs, nano, code (vscode)")
	fmt.Println("Note this application will invoke this command verbatim when you choose to edit a file")

	editorChoice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	cfg.PreferredEditor = strings.TrimSpace(editorChoice)
	fmt.Println("cool :o I also love", strings.TrimSpace(editorChoice))

RemoteLoop:
	for {
		fmt.Println("Just making sure, but you know we're working out of", cfg.Git.RemoteURL, "right? [yes]/[no]")
		remoteConfirm, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		switch remoteConfirm {
		case "yes\n":
			fmt.Println("owo cool :)")
			break RemoteLoop
		case "no\n":
			fmt.Println("oh, where should I point it? Leave blank to cancel")
			fmt.Println("format: gitURL.com/author/repository")
			newRemote, err := reader.ReadString('\n')
			newRemote = strings.TrimSpace(newRemote)
			if err != nil {
				log.Fatal(err)
			}
			if newRemote != "" {
				cfg.Git.RemoteURL = newRemote
			}
		}
	}

	cfg.Initialized = true
	fmt.Println("Config Saved! Thanks! You should probably run 'owo -update' just in case :)")
	return cfg
}
