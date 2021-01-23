package commands

import "fmt"

const About = "about"

func ShowAbout() {
	fmt.Println("OwO is a command line executable that allows easy access of small amounts of data organized by a filestructure in the git repository")
	fmt.Println("It allows for easy editing of this data to a source of truth hosted on a github server")
	fmt.Println("Handy for people working often on command lines who need access to handy commands and/or documentation")
}
