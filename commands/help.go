package commands

import "fmt"

const (
	Help      string = "help"
	HelpShort string = "h"
)

func CmdHelp() {

	fmt.Println("owo")
	fmt.Println()
	fmt.Println("owo base commands:")
	fmt.Println("owo -[h]elp        | show this page")
	fmt.Println("owo -[a]bout       | cool information about owo")
	fmt.Println("owo -[c]onfig      | rerun that cool first time setup")
	fmt.Println("owo -[ch]eckup     | check remote upstream if there are updates")
	fmt.Println("owo -[u]pdate      | attempt to update owo")
	fmt.Println("owo -[s]ync        | push changes to docs upstream")
	fmt.Println("owo -[r]eset       | destroy local changes to docs")
	fmt.Println("owo -diff          | list files that are different from remote")
	fmt.Println("owo fileactions    | see file action documentation from inside owo")
}
