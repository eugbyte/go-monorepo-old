# go monorepo
## About
monorepo in golang

The service/greet module can reference the libs/util modules

The crucial feature of go 1.18 that enables monorepo is the use of go workspaces, or `go.work`

## Installation
```
pip install https://github.com/joh/when-changed/archive/master.zip
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
1. When running `go mod tidy`, packages specified in the [`go.work` will not be ignored](https://github.com/golang/go/issues/50750). So, do `go mod tidy -e` instead. The `-e` flag causes go mod tidy to attempt to proceed despite errors encountered while loading packages.
