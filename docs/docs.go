package docs

import "fmt"

type Entry struct {
	Command string
	Title   string
	Body    []string
	Notes   []string
}

func (e Entry) Print() {
	fmt.Println("-----------------------------")
	if e.Title != "" {
		fmt.Println(e.Title)
		fmt.Println()
	}
	if e.Command != "" {
		fmt.Println("Command:")
		fmt.Println("  ", e.Command)
		fmt.Println()
	}
	if len(e.Body) != 0 {
		for _, bodyLine := range e.Body {
			fmt.Println(bodyLine)
		}
		fmt.Println()
	}
	if len(e.Notes) != 0 {
		fmt.Println("Notes:")
		for _, note := range e.Notes {
			fmt.Println("", note)
		}
	}
	fmt.Println("-----------------------------")
}

type Folder struct {
	Path       []string
	Entries    []string
	SubFolders []string
	Meta       string // Meta gets displayed above the folder contents
}

func (f Folder) Print() {
	fmt.Println(f.Path)
	if f.Meta != "" {
		fmt.Println("          " + f.Meta)
	}

	if len(f.SubFolders) == 0 && len(f.Entries) == 0 {
		fmt.Println("  folder is empty")
		return
	}

	if len(f.SubFolders) != 0 {
		for _, folder := range f.SubFolders {
			fmt.Println("  [" + folder + "]")
		}
	}
	if len(f.Entries) != 0 {
		for _, entry := range f.Entries {
			fmt.Println("  " + entry)
		}
	}
}
