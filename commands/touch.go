package commands

import "fmt"

const Touch string = "touch"

func ShowTouchHelp() {
	fmt.Println("Error: command -touch syntax not recognized")
	fmt.Println()
	fmt.Println("-touch expects a folder path before it, and a newfile name after it")
	fmt.Println("  example: owo directory1 subdirectory2 -touch newFileName")
}
