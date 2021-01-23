package main

import (
	"fmt"
	"os"
	"owo/commands"
)

func main() {
	args := os.Args
	// fmt.Println(args)

	if len(args) == 1 {
		fmt.Println("welcome to owo")
		fmt.Println("--------------")
		fmt.Println("try:")
		fmt.Println("owo about")
		fmt.Println("owo checkup")
		fmt.Println("owo update")
		os.Exit(0)
	}

	switch args[1] {
	case commands.Checkup:
		commands.CheckForUpdate()
	case commands.About:
		commands.ShowAbout()
	case commands.Update:
		err := commands.CmdUpdate()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return
}
