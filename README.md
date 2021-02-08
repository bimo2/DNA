![Banner](.github/assets/banner.svg)

# Using DNA

**DNA is a developer workflow protocol for running command line scripts.** DNA can be configured to run common developer workflows like installing dependencies, starting servers, creating builds, running tests... with a simple CLI.

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

For DNA to work globally, the `_` binary needs to be in available in `PATH`. Adding `_` to `usr/local/bin` is recommended:

```zsh
cp /path/to/DNA/dist/_ /usr/local/bin
```

## Configure

DNA searches for a `dna.json` config file for workflow commands. It will search the current directory and recursively look up (`..`) until a config file is found or a `.git` file is reached. Get started by creating a `dna.json` template:

```zsh
_ init
```

Any `dna.json` config contains a top level `DNA` object with the DNA spec `version` for parsing workflows. Workflows are defined in the top level `scripts` object:

```json
{
  "DNA": {
    "version": 0,
    "spec": "https://github.com/bimo2/DNA"
  },
  "scripts": {
    "buy:xrp": {
      "info": "Buy 750 XRP tokens",
      "commands": [
        "buy 750.000 xrp",
        "capture -cc AMEX_1",
        "deposit bimo2$balance.to"
      ]
    }
  }
}
```

Workflows can be called by their keys and DNA will execute the specified commands synchronously relative to the `dna.json` directory.

```zsh
_ ls
# ...
# buy:xrp       Buy 750 XRP tokens
# ...

_ buy:xrp
# 0 `buy 750.000 xrp`
# 1 `capture -cc AMEX_1`
# 2 `deposit bimo2$balance.to`
```

# Develop

```zsh
# run from source
go run ./_

# build binary
sh bin/build.sh
```

#

MIT. Copyright &copy; 2021 Bimal Bhagrath
