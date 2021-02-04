package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	// TARGET : DNA file name
	TARGET = "dna.json"

	// GITFILE : repo top directory identifier
	GITFILE = ".git"

	// DEPTH : max recursion depth
	DEPTH = 50
)

// DNAFile : user defined DNA config
type DNAFile struct {
	Version int
	Require []string
	Scripts map[string]DNAScript
}

// DNAScript : user defined script
type DNAScript struct{}

// Find : recursively find a DNAFile in the current repo
func Find() *DNAFile {
	var config *DNAFile

	for i, directory, end := 0, ".", false; config == nil && !end && i < DEPTH; i, directory = i+1, directory+"/.." {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			break
		}

		for _, file := range files {
			if file.Name() == TARGET {
				config = parse(file)
			}

			if file.Name() == GITFILE {
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

	return &output
}
