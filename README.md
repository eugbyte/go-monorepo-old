# go monorepo
## About
monorepo in golang

The crucial feature of go 1.18 that enables monorepo is the use of go workspaces, or `go.work`

## Installation
install `azure functions core tools` [here](https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=v4%2Cwindows%2Cpowershell%2Cazurecli%2Cbash&source=docs#install-the-azure-functions-core-tools)

```
pip install https://github.com/joh/when-changed/archive/master.zip
sudo apt-get install docker-ce docker-ce-cli containerd.io | choco install docker-desktop (windows)
apt-get install make | choco install make (windows)
make test-install-gotest
make lint-install
```

## Development
Full list of commands are listed in Makefile
If you are on windows, you need to have `git bash` cli installed to run the commands

## Start azure functions core server
`make run`

## Watch files and recompile whenever they change
Open another terminal
`make watch`

## Gotchas
1. `function.json` cannot be in nested directories, see [here](https://github.com/Azure/azure-functions-host/issues/5373)
2. Azure functions core tools does not have a watch option. Thus, need to keep manually rebuilding the executable file.
3. When running `go mod tidy`, packages specified in the [`go.work` will not be ignored](https://github.com/golang/go/issues/50750). So, do `go mod tidy -e` instead. The `-e` flag causes go mod tidy to attempt to proceed despite errors encountered while loading packages.
