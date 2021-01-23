package docs

import "fmt"

type Entry struct {
	Command string
	Title   string
	Body    []string
	Notes   []string
}

func (e Entry) Print() {
	fmt.Println(e.Title)
	fmt.Println()
	fmt.Println("  Command")
	fmt.Println(e.Command)
	fmt.Println()
	for _, bodyLine := range e.Body {
		fmt.Println(bodyLine)
	}
	fmt.Println()
	fmt.Println("  Notes:")
	for _, note := range e.Notes {
		fmt.Println(note)
	}
}

type Folder struct {
	Name       string
	Entries    []Entry
	SubFolders []Folder
	Meta       string // Meta gets displayed above the folder contents
}

func (f Folder) Print() {
	fmt.Println(f.Name + "{")
	fmt.Println("   " + f.Meta)
	fmt.Println()
	if len(f.SubFolders) != 0 {
		fmt.Println("   Folders:")
		for _, folder := range f.SubFolders {
			fmt.Println("     " + folder.Name)
		}
	}
	if len(f.Entries) != 0 {
		fmt.Println("   Entries:")
		for _, entry := range f.Entries {
			fmt.Println("     " + entry.Title)
		}
	}
}
