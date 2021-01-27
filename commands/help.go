package commands

import "fmt"

const Help = "help"

func CmdHelp() {

	fmt.Println("owo")
	fmt.Println()
	fmt.Println("owo -help       | show this page")
	fmt.Println("owo -about      | cool information about owo")
	fmt.Println("owo -config     | rerun that cool first time setup")
	fmt.Println("owo -checkup    | check remote upstream if there are updates")
	fmt.Println("owo -update     | attempt to update owo")
}
