![Banner](.github/assets/banner.svg)

**DNA is a dev tool for automating project tasks.** You can configure DNA to install dependencies, manage servers, create builds, run tests... with a global CLI.

## Using DNA

### No Build + Install

```zsh
curl -sf https://gobinaries.com/bimo2/DNA/_ | sh
```

### Build + Install

[Build the DNA binary](#develop) then add it to your `$PATH` variable. Copying `_` to `/usr/local/bin` is recommended:

```zsh
cp /path/to/DNA/dist/_ /usr/local/bin
```

### Configure

Start by creating a `dna.json` file for your project:

```zsh
_ init
```

#### `dna.json`

`_` looks for a `dna.json` file for task definitions. It'll search your current directory then recursively search `..` until a config file is found or a `.git` directory is reached.

Any `dna.json` config should contain a top level `_version` key to specify the DNA spec to use for parsing tasks. Project tasks are defined by name in the `scripts` object. Your config file might look like this:

```json
{
  "_version": 0,
  "_system": {
    "Banking": "/usr/local/bin/banking"
  },
  "env": {
    "BURN_ADDRESS": "X7TYFRtYHMcHtT2qNycMwgXzqPp7Pb16cH84uj5Hc7GtrsB"
  },
  "scripts": {
    "buy:xrp": {
      "info": "Buy XRP tokens",
      "commands": [
        "buy [amount=100] xrp",
        "deposit -m [address]"
      ]
    },
    "exchange": {
      "info": "Convert crypto assets",
      "commands": [
        "convert [amount] [current] -> [next]"
      ]
    },
    "burn:xrp": {
      "info": "Burn all XRP tokens",
      "commands": [
        "# SKIP: verification",
        "deposit -m &BURN_ADDRESS"
      ]
    }
  }
}
```

#### Arguments

Tasks can accept multiple arguments by passing them to the `_` command. You can define arguments by adding `[templates]` to commands. Arguments are passed to the task as a stack so each template in the execution is replaced by the next available argument in the stack. You can specify default values for undefined arguments by adding them to the `[template=0]`. Templates without defaults will resolve to an empty string.

#### Environment Variables

Environment variables are DNA constants that can be used in any command. You can set environment variables by defining them in the `env` object. Note that all environment variables should be valid JSON strings only. Variables can be referenced in commands by name prefixed with `&`.

#### System Requirements

The `_ check` command verifies if the system has the required programs and files installed to use the project scripts. You can define system requirements as key/values in the top level `_system` object. For commands that should be available in `$PATH`, use the command name (ex. `go` or `node`) as the value and for required files, like dynamic C/C++ headers, use the absolute path of the file.

#### Comments

Commands starting with `#` will be printed as comments and skipped while executing the task. This can be useful to print logs, warnings or TODOs.

### Test Drive

Using the `dna.json` file defined above, here's what would happen:

```zsh
_ buy:xrp 750.000 bimo2$balance.to
# `buy 750.000 xrp`
# `deposit -m bimo2$balance.to`

_ buy:xrp
# `buy 100 xrp`
# `deposit -m`

_ exchange 100 xrp btc
# `convert 100 xrp -> btc`

_ burn:xrp
# `echo "SKIP: verification"`
# `deposit -m X7TYFRtYHMcHtT2qNycMwgXzqPp7Pb16cH84uj5Hc7GtrsB`
```

## Develop

DNA is built with Go.

```zsh
# run source
go run ./_

# build binary
go build -o ./dist/_ ./_

# clean repo
git clean -Xfd
```

#

<sub><sup>**MIT.** Copyright &copy; 2021 Bimal Bhagrath</sup></sub>
