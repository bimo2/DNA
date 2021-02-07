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
		"test": {
			"info": "Run unit + integration test suites",
			"commands": [
				"echo \"YOUR UNIT TEST COMMAND\"",
				"echo \"YOUR INTEGRATION TEST COMMAND\""
			]
		},
		"dev": {
			"info": "Start development server",
			"commands": [
				"echo \"YOUR DEV COMMAND\""
			]
		}
	}
}
`

// Initialize : create template dna.json
func Initialize(filename string) {
	var data bytes.Buffer
	err := json.Indent(&data, []byte(template), "", "  ")

	if err != nil {
		console.Error("Failed to build template")
	}

	ioutil.WriteFile("./"+filename, data.Bytes(), 0777)
	console.Message("Created `"+filename+"` template!", nil)
}
