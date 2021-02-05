package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	gitfile = ".git"
	depth   = 50
)

// DNAFile : user defined DNA config
type DNAFile struct {
	DNA struct {
		Version int
		Spec    string
	}
	Scripts map[string]DNAScript
}

// DNAScript : user defined script
type DNAScript struct {
	Info     string
	Commands []string
}

// Load : recursively find a DNAFile in the current repo
func Load(filename string) *DNAFile {
	var config *DNAFile

	for i, directory, end := 0, ".", false; config == nil && !end && i < depth; i, directory = i+1, directory+"/.." {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			break
		}

		for _, file := range files {
			if file.Name() == filename {
				config = parse(file)
			}

			if file.Name() == gitfile {
				end = true
			}
		}
	}

	return config
}

func parse(file os.FileInfo) *DNAFile {
	var output DNAFile
	content, err := ioutil.ReadFile(file.Name())

	if err != nil {
		return nil
	}

	json.Unmarshal(content, &output)

	fmt.Println(output)

	return &output
}
