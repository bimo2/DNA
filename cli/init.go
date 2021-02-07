package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/bimo2/DNA/console"
)

const template = `
{
	"DNA": {
		"version": 0,
		"spec": "https://github.com/bimo2/DNA"
	},
	"scripts": {
		"install": {
			"info": "Install dependencies",
			"commands": [
				"echo \"YOUR INSTALL COMMAND\""
			]
		},
		"build": {
			"info": "Package for distribution",
			"commands": [
				"echo \"YOUR BUILD COMMAND\""
			]
		},
		"dev": {
			"info": "Start development server",
			"commands": [
				"echo \"YOUR DEV COMMAND\""
			]
		},
		"test": {
			"info": "Run test suites",
			"commands": [
				"echo \"YOUR UNIT TEST COMMAND\"",
				"echo \"YOUR INTEGRATION TEST COMMAND\""
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
