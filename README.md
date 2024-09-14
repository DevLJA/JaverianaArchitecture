# poc-circuit-breaker on Golang

# Go Starter App

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

Poc of basic implementation of the circuit breaker architecture pattern as an option used to detect failures and encapsulates the logic of preventing a failure from recurring constantly during the time it takes to recover the consumed service

## Available Endpoints

* **_POST /status_client/:user_ui_**
* **_POST /user/:user_ui_**

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

## Tests and operation of the circuit breaker pattern
