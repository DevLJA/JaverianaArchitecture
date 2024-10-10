# Javeriana - Laura Joya

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

# Go Starter App

## First steps

### Go Runtime Version

It is necessary that the golang version is installed on your computer.

For this project, it is recommended to use version `go1.23.0`

### Dependency management

The list of external dependencies used by the project is found in the `go.mod` file. 
It is important that you download the necessary dependencies. By running the following command: `go mod tidy`

### Build and run application

Once you have the dependencies installed, run the following command to build the project `go build ./...`

> If you have the Golang IDE installed, just go to the `main.go` file and click `run`


If you want to run it through a terminal, go to the folder that contains the project and verify that it has the `PATH` and `GOPATH` paths successfully.

```
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH="$HOME/go/bin
```

Finally, run the `go run .` command to run the project

## Take a look at project list:

1. [POC-circuit-breaker on Golang](CircuitBreaker.md)

2. [Chess Image Generator](ChessImage.md)

