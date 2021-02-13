package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/bimo2/DNA/console"
)

const template = `
{
	"_version": 0,
	"scripts": {
		"install": {
			"info": "Install dependencies",
			"commands": [
				"# YOUR INSTALL COMMAND"
			]
		},
		"build": {
			"info": "Package for distribution",
			"commands": [
				"# YOUR BUILD COMMAND"
			]
		},
		"dev": {
			"info": "Start development server",
			"commands": [
				"# YOUR DEV COMMAND"
			]
		},
		"test": {
			"info": "Run test suites",
			"commands": [
				"# YOUR UNIT TEST COMMAND",
				"# YOUR INTEGRATION TEST COMMAND"
			]
		}
	}
}
`

// Init : create dna.json template
func Init(filename string) {
	var data bytes.Buffer

	if err := json.Indent(&data, []byte(template), "", "  "); err != nil {
		console.Error("Failed to build template")
	}

	path := "./" + filename
	ioutil.WriteFile(path, data.Bytes(), 0777)
	message := "Created new `" + filename + "` template!"
	console.Message(message, nil)
}
