package parse

import (
	"io/ioutil"
	"strings"

	"github.com/LinkdxTTV/owo/docs"
)

func ParseEntry(pathToFile string) (*docs.Entry, error) {
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	entry := docs.Entry{}
	lines := strings.Split(string(data), "\n")
	var inBody bool
	var inNotes bool
	for _, line := range lines {
		if line == "" && inBody == false && inNotes == false {
			continue
		}
		keyword := checkKeyword(line)
		switch keyword {
		case &title:
			titleSplit := strings.TrimSpace(strings.Split(line, ":")[1])
			entry.Title = titleSplit
		case &command:
			cmdSplit := strings.TrimSpace(strings.Split(line, ":")[1])
			entry.Command = cmdSplit
		case &bodystart:
			inBody = true
		case &bodyend:
			inBody = false
		case &notes:
			inNotes = true
		case nil:
			if inBody {
				entry.Body = append(entry.Body, line)
			}
			if inNotes {
				entry.Notes = append(entry.Notes, line)
			}
		}
	}
	return &entry, nil
}

var (
	title     string = "Title"
	command   string = "Command"
	bodystart string = "BodyStart"
	bodyend   string = "BodyEnd"
	notes     string = "Notes"
)

func checkKeyword(line string) *string {
	if line[0:5] == title+":" {
		return &title
	} else if line[0:7] == command+":" {
		return &command
	} else if line[0:9] == bodystart+":" {
		return &bodystart
	} else if line[0:7] == bodyend+":" {
		return &bodyend
	} else if line[0:5] == notes+":" {
		return &notes
	}
	return nil
}
