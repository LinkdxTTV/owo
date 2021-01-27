package commands

import "fmt"

const About = "about"

func CmdAbout() {
	fmt.Println(`
    ██████  ██     ██  ██████  
   ██    ██ ██     ██ ██    ██ 
   ██    ██ ██  █  ██ ██    ██
   ██    ██ ██ ███ ██ ██    ██ 
    ██████   ███ ███   ██████    
   `)
	fmt.Println("\033[1m                         whats this? \033[0m")

	fmt.Println("  motivation:")
	fmt.Println("owo was born out of frustration at using the confluence web wiki. It is very hard to quickly access the information I want, and editing the info with new info is a huge pain. For instance, I want to ensure that the health endpoint of one of my services is responding correctly, but I dont remember the port. Should I open the repo and look it up? Oh, I can just check which ports are listening on this host, but what was that command again? What about the command to restart the service as managed by daemon tools? Is it on the wiki? I normally have tons of terminals open so this seemed like a good compromise. Similarly, I wanted some exercise in writing a command line interface handling all the operations myself, including io and code generation. As a result, owo has no external dependencies and only requires go and git be installed. Finally, owo is called owo because memes control everything around me UwU")
	fmt.Println()
	fmt.Println("  description:")
	fmt.Println("owo is a command line executable that allows easy access of small amounts of documentation organized by a standard filestructure in a git repository. owo makes heavy use of its own source code, located in your GOPATH when it was installed. this allows for easy editing of this data to a source of truth hosted in a github repo. owo is handy for people working often on command lines who need access to handy commands and/or documentation. Additional works on github actions and integrations serve to expand functionality and serve as access points for non developers. owo helps you organize and update your data by allowing for easy updates of information facilitated through git pull requests auto generated by owo.")
	fmt.Println()
	fmt.Println("thx for using :) owo only gets better as more people use it - @nelson")
	fmt.Println()
}
