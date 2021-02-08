package protocol

import (
	"encoding/json"
	"errors"
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

// Find : recursively find a DNAFile in the current repo
func Find(filename string) (*DNAFile, *string, error) {
	var config *DNAFile
	var path string
	var err error

	for i, directory, end := 0, ".", false; config == nil && !end && i < depth; i, directory = i+1, directory+"/.." {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			err = errors.New("Failed to read directory")
			break
		}

		for _, file := range files {
			if file.Name() == filename {
				config, err = parse(directory, file)
				path = directory
			}

			if file.Name() == gitfile {
				end = true
			}
		}
	}

	return config, &path, err
}

func parse(path string, file os.FileInfo) (*DNAFile, error) {
	var output DNAFile
	content, err := ioutil.ReadFile(path + "/" + file.Name())

	if err != nil {
		return nil, errors.New("Failed to parse: " + err.Error())
	}

	json.Unmarshal(content, &output)
	return &output, nil
}
