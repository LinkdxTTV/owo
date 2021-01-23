package commands

import (
	"fmt"
	"os/exec"
	"owo/config"
	"strings"
)

const Checkup string = "checkup"

// CheckForUpdate checks for updates
func CheckForUpdate() error {
	cfg, err := config.ParseConfig()
	if err != nil {
		return err
	}

	headSHA, err := getNewestSHA(cfg)

	fmt.Println("Current build: ", cfg.Git.SHA)
	fmt.Println("Remote  build: ", headSHA)

	if headSHA == cfg.Git.SHA {
		fmt.Println("owo you're up to date :)")
	} else {
		fmt.Println("Please run owo update")
	}
	return nil
}

func getNewestSHA(cfg *config.Config) (string, error) {
	gitCmd := exec.Command("git", "ls-remote", cfg.Git.RemoteURL)

	outBytes, err := gitCmd.Output()
	if err != nil {
		return "", err
	}

	// TODO maybe improve this parsing. I could use bash commands to do this but it might break functionality between OS'es..
	var headSHA string
	shaLines := strings.Split(string(outBytes), "\n")
	for _, line := range shaLines {
		if strings.Contains(line, "HEAD") {
			headSHA = strings.TrimSpace(strings.Split(line, "HEAD")[0])
			break
		}
	}
	return headSHA, nil
}
