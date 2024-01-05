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

### Greetings

In the folder *greetings* there are examples of unit tests

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
