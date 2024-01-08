# echoStars

A web services server monitor

## Collaboration

* [Trello](https://trello.com/c/FogW1SUW)
* [Repo](https://github.com/emanuelpeg/echoStars/)
* Authors: <emanuelpeg@gmail.com>, <axel.salmeron@gmail.com>, <juanpabloalvis@gmail.com>, <sergio.delatorre@unmsm.edu.pe>

## Prerequisites

* Go installation:
  
  Install go in the local machine. It should be done setting the environment variable *GOROOT*.
  The *GOROOT* variable points where your Go SDK is located.

  Please refer to Go Lang [installation page](https://go.dev/doc/install) for more details.

  You can check in your terminal the installation with:

  ```bash
    go build
  ```

* Text editor or your preferred IDE.

## Build and run application using CLI

Go to the directory of the source code and build:

```bash
  go build
```

Then, you can run the application

```bash
$env:APP_PROFILE="dev"; go run main.go
```

the `AP_PROFILE` environment variable will indicate which `.property` file the app will use

**EchoStars** is a Web application that runs on port *1323*. Please, check logs and then you can open in your preferred browser; here you have the [api documentation uri](http://localhost:1323/swagger/index.html).

## Write test with Mocks

### Runing all tests

```bash
go test -v ./...
```

### Test examples "Greetings"

In the folder *greetings* there are a few examples of unit tests

### GoMock

[Install gomock](https://github.com/uber-go/mock)

```bash
go install go.uber.org/mock/mockgen@latest
```

Then run *mockgen* with the file that has the interface that should be mock `mockgen -source=<sourceFilePath> -destination=<destinationFilePath>`

For example :

```bash
mockgen -source=C:\projects\goLang\echoStars\greetings\greetings.go -destination=C:\projects\goLang\echoStars\greetings\greetings.mock.go
```

#### The Steps are

1. Define an interface that you wish to mock.

    ```Go
    type MyInterface interface {
      SomeMethod(x int64, y string)
    }
    ```

2. Use *mockgen* to generate a mock from the interface.

3. Use the mock in a test:

    ```Go
          func TestMyThing(t *testing.T) {
            mockCtrl := gomock.NewController(t)
            defer mockCtrl.Finish()
            mockObj := something.NewMockMyInterface(mockCtrl)
            mockObj.EXPECT().SomeMethod(4, "blah")
            // pass mockObj to a real object and play with it.
          }
    ```

## Routines

Go to the *routines* subdirectory  and use `go run .` to run the routines test code.

## Express test server

*ExpressJs server* for testing the monitor service with 3 endpoints.

* <http://localhost:3000/healthCheck/success>: always succeeds
* <http://localhost:3000/healthCheck/error>  always fails
* <http://localhost:3000/healthCheck/sketchy> fails 10% of the time

When running the monitor service these endpoints will be checked periodically even if the test server is down.

To run the test server go to the *express-test-server* subdirectory and run:

```bash
npm install
npm start
```

## Docker

To run the app in docker use `docker-compose up`. Add the `--build` option if you want to force a build of the images before staring the containers. Add the `-d` option if you want docker-compose to run in dettached mode

To stop the containers while running in dettached mode run `docker-compose down`. Use `-v` if you want to remove the volumes

To check the logs use `docker-compose logs [service-name]` where *service-name* is optional

Use `docker-compose ps` to list the currently running containers

## GRpc
### Install

Install with following commands:
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go: downloading google.golang.org/protobuf v1.28.1
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go: downloading google.golang.org/grpc v1.2.1
```
### Develop

To develop a GRpc services should follow these steps:

1. Write a .proto, for example :
```Go
syntax = "proto3";

package greetings;
option go_package = "echoStarts/greetings";

message HelloRequest {
  string message = 1;
}

message HelloResponse {
  string message = 1;
}

message HelloWoldRequest {
}

message HelloWoldResponse {
  string message = 1;
}

//  Greetings service.
service GreeterServiceGrpc {

  rpc Hello(HelloRequest) returns (HelloResponse) {}
  rpc HelloWold(HelloWoldRequest) returns (stream HelloWoldResponse) {}

}
```
2. Generate code from the .proto. With this command:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./greetings/greetings.proto
```

3. Write a server and a client. see [server](greetings/grpc_server/main.go) and [client](greetings/grpc_client/main.go)