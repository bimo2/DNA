package lib

import (
	"encoding/json"
	"io/ioutil"
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
	version int
	require []string
	tasks   json.RawMessage
}

// Find : recursively find a DNAFile in the current repo
func Find() *DNAFile {
	var target *DNAFile

	for i, directory, stop := 0, ".", false; target == nil && !stop && i < DEPTH; i, directory = i+1, directory+"/.." {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			break
		}

		for _, file := range files {
			if file.Name() == TARGET {
				target = &DNAFile{}
			}

			if file.Name() == GITFILE {
				stop = true
			}
		}
	}

	return target
}
