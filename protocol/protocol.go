package protocol

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bimo2/DNA/console"
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
func Load(filename string) (*DNAFile, error) {
	var config *DNAFile
	var err error

	for i, directory, end := 0, ".", false; config == nil && !end && i < depth; i, directory = i+1, directory+"/.." {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			break
		}

		for _, file := range files {
			if file.Name() == filename {
				config, err = parse(file)
			}

			if file.Name() == gitfile {
				end = true
			}
		}
	}

	return config, err
}

func parse(file os.FileInfo) (*DNAFile, error) {
	var output DNAFile
	content, err := ioutil.ReadFile(file.Name())

	if err != nil {
		console.Error("Failed to parse")
		return nil, err
	}

	json.Unmarshal(content, &output)
	return &output, nil
}
