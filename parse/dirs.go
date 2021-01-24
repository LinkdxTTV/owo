package parse

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/LinkdxTTV/owo/config"
	"github.com/LinkdxTTV/owo/docs"
)

func NavigateAndShow(argPath []string, cfg *config.Config) error {
	fullpath := cfg.LocalPath + "/docs/text"
	for _, subdir := range argPath {
		fullpath += "/" + subdir
	}

	// Try to see if its a file
	entry, err := ParseEntry(fullpath)
	if err != nil {
		if strings.Contains(err.Error(), "is a directory") { // SUPER HACKY NO SHAME PLS
			NavigateAndShowDir(fullpath, cfg)
		} else if strings.Contains(err.Error(), "no such") {
			fmt.Println("no such file or folder. showing directory...")
			NavigateAndShowDir(path.Dir(fullpath), cfg)
		} else {
			return err
		}
	} else {
		entry.Print()
	}
	return nil
}

func NavigateAndShowDir(path string, cfg *config.Config) error {

	directory, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	dir := strings.Split(path, "/")

	folder := docs.Folder{
		Name:       dir[len(dir)-1],
		Entries:    []string{},
		SubFolders: []string{},
	}

	for _, item := range directory {
		if item.IsDir() {
			folder.SubFolders = append(folder.SubFolders, item.Name())
		} else {
			if item.Name() == "meta" {
				metadata, err := ioutil.ReadFile(path + "/meta")
				if err != nil {
					return err
				}
				folder.Meta = string(metadata)
			} else {
				folder.Entries = append(folder.Entries, item.Name())
			}
		}
	}

	folder.Print()

	return nil
}
