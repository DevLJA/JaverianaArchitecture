# POC-circuit-breaker on Golang - Laura Joya

# Go Starter App

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

POC of basic implementation of the [circuit breaker architecture pattern](https://apiumhub.com/es/tech-blog-barcelona/patron-circuit-breaker/) as an option used to detect failures and encapsulates the logic of preventing a failure from recurring constantly during the time it takes to recover the consumed service

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

### Test Description

**Test data:**
We understand `service A` to be our main service (where we have the circuit breaker) and `service B` would be a third-party service on which we depend to return a successful response.

- Number of times service B is consumed before opening the circuit breaker: `3`
- Time that the system will wait to perform a new monitoring of service B: `10 seconds`


To simulate the state of `service B`, it was determined that when `user_id` `123` is used it will always give a **successful response,** otherwise the service **throws an error**.

The application manages three states on the circuit breaker:
1. **Closed**: No abnormality is detected in `service B`. Everything works fine, so `service B` is consumed whenever necessary
2. **Open**: It was detected that `service B` is failing continuously, so the circuit is opened in such a way that we capture the error and `service B` is not consumed.
3. **Half-Open**: Monitoring status, to give the opportunity for `service B` to become available again. The call is made occasionally to verify if the service is working or not

### Test objective(s)
Carry out a proof of concept regarding the correct functioning of a circuit breaker, for decision making when consuming an external service.

> At a look at the videos found in the` test_evidence.zip` folder, functional tests of the application were carried out there

### Steps implemented to carry out the test
First of all, take a look at the `First steps` section to configure your environment and run the test locally.

> Change the value `123` to another number if you want to simulate an error state.

Service A

```
curl --location 'http://localhost:8080/status_client/123'
```

Service B

```
curl --location 'http://localhost:8080/user/123'
```

Responses:

The circuit is open, this means that we cannot generate a successful response, but the external service (`service B`) is not being degraded further (because it is not being consumed):
```
{
    "error": "Soft error: circuit breaker is OPEN"
}
```

The service is in monitoring state, so in the next request it receives, it will verify whether `service B` is working or not and from there it will update the state of the circuit breaker (open or closed)
```
{
    "error": "Soft error: circuit breaker is HALF-OPEN"
}

```
`Service B` was consumed and an error was detected:

```
{
    "error": "users-api was consumed and threw an error"
}
```

`Service B` works normally, it is possible to deliver a response with status 200
```
{
    "data": {
        "user": {
            "id": "123",
            "name": "Laura Joya",
            "email": "lauracja.123@gmail.com",
            "age": 27
        },
        "status": {
            "status": "active",
            "points": 23456543
        }
    }
}
```
### Technologies used in the test (specify languages, libraries)

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

### Results

After implementing and testing the Circuit Breaker in the proof of concept, we have observed that the system behaves robustly and efficiently, confirming that the pattern contributes significantly to the resilience and reliability of the system.

Making use of the three states of the circuit breaker that allows constant monitoring of the external services consumed and thus making decisions regarding stopping its constant consumption.

### Conclusions
The Circuit Breaker pattern is crucial when planning a system architecture because it prevents cascading failures and maintains system stability against external services or components that could fail. It acts as a protective measure, preventing repeated errors from overwhelming the system and allowing controlled recovery.