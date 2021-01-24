package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/LinkdxTTV/owo/config"
)

const Checkup string = "-checkup"

// CheckForUpdate checks for updates
func CheckForUpdate(cfg *config.Config) (bool, error) {

	headSHA, err := getNewestSHA(cfg)
	if err != nil {
		return false, err
	}

	fmt.Println("Current build: ", cfg.Git.SHA)
	fmt.Println("Remote  build: ", headSHA)

	return headSHA != cfg.Git.SHA, nil
}

func getNewestSHA(cfg *config.Config) (string, error) {
	gitCmd := exec.Command("git", "ls-remote", cfg.Git.SSHURL)

	outBytes, err := gitCmd.Output()
	if err != nil {
		return "", err
	}

	// TODO maybe improve this parsing. I could use bash commands to do this but it might break functionality between OS'es.. string maniupalation always feels dirty though
	var headSHA string
	shaLines := strings.Split(string(outBytes), "\n")
	for _, line := range shaLines {
		if strings.Contains(line, "HEAD") { // should always be first
			headSHA = strings.TrimSpace(strings.Split(line, "HEAD")[0])
			break
		}
	}
	return headSHA, nil
}
