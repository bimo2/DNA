![Banner](.github/assets/banner.svg)

**DNA is a developer workflow protocol for running command line scripts.** DNA can be configured to run common developer workflows like installing dependencies, starting servers, creating builds, running tests... with a simple CLI.

# Using DNA

## No Build + Install

```zsh
curl -sf https://gobinaries.com/bimo2/DNA/_ | sh
```

## Build + Install

```zsh
# requires Go ^1.15

cd path/to/DNA
sh bin/build.sh
```

For DNA to work globally, the Go binary needs to be in available in `PATH`. Adding `_` to `usr/local/bin` is recommended:

```zsh
cp /path/to/DNA/dist/_ /usr/local/bin
```

## Configure

DNA searches for a `dna.json` config file for workflow commands. It will search the current directory and recursively look up (`..`) until a config file is found or a `.git` file is reached. Get started by creating a `dna.json` template:

```zsh
_ init
```

Any `dna.json` config should contain a top level `_version` key to specify the DNA spec to use for parsing workflows. Project workflows are defined by name in the `scripts` object with the following properties:

```json
{
  "_version": 0,
  "scripts": {
    "buy:xrp": {
      "info": "Buy 750 XRP tokens",
      "commands": [
        "buy 750.000 xrp",
        "pay | capture -cc amex_2",
        "deposit bimo2$balance.to"
      ]
    }
  }
}
```

Workflows can be executed by script key. DNA will execute the specified commands synchronously relative to the `dna.json` directory. 

```zsh
_ ls
# ...
# buy:xrp       Buy 750 XRP tokens
# ...

_ buy:xrp
# 0 `buy 750.000 xrp`
# 1 `pay | capture -cc amex_2`
# 2 `deposit bimo2$balance.to`
```

# Developers

```zsh
# run from source
go run ./_

# build binary
sh bin/build.sh

# clean repository
sh bin/clean.sh
```

#

###### MIT License. Copyright &copy; 2021 Bimal Bhagrath
