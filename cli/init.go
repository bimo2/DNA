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
	"env": {
		"PORT": "default_port"
	},
	"scripts": {
		"install": {
			"info": "Install dependencies",
			"commands": [
				"# YOUR INSTALL COMMAND"
			]
		},
		"dev": {
			"info": "Start development server",
			"commands": [
				"# YOUR DEV COMMAND &PORT"
			]
		},
		"build": {
			"info": "Create production build",
			"commands": [
				"# YOUR BUILD COMMAND [beta=0]"
			]
		},
		"test": {
			"info": "Run test suites",
			"commands": [
				"# YOUR TEST COMMAND [file]"
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
