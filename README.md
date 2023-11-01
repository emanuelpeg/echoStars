# echoStars
A web services server monitor

## Collaboration
* [Trello](https://trello.com/c/FogW1SUW)
* Repo https://github.com/emanuelpeg/echoStars/

## Prerequisites

* Go installation.
* Text editor or your preferred IDE.

### Installation
We need to install go in the local machine. It should be done setting the environment variable GOROOT.
The GOROOT variable points where your Go SDK is located.
You can check in your terminal the installation with:
>go version

Please refer to Go Lang [installation page](https://go.dev/doc/install) for more details.


### Build and run application using CLI
Go to the directory of the source code and build:
>go build

then, you can run the application (if you are running in Windows, don't forget scape the :
>go run main.go

EchoStart is a Web application that runs on port _1323_. Please, check logs and then you can open in your preferred browser; here you have the api documentation http://localhost:1323/swagger/index.html uri.


# Write test with Mocks

## Greetings

In the folder greetings there are examples of a unit test

## Routines

Run `cd routines` and `go run .` to run the routines test code

## GoMock

Install gomock : https://github.com/uber-go/mock

>go install go.uber.org/mock/mockgen@latest

Then run mockgen with the file that has the interface that should be mock

>mockgen -source=... -destination=...

For example : 

>mockgen -source=C:\projects\goLang\echoStars\greetings\greetings.go -destination=C:\projects\goLang\echoStars\greetings\greetings.mock.go

### The Steps are: 

1. Define an interface that you wish to mock.
      type MyInterface interface {
        SomeMethod(x int64, y string)
      }
2. Use mockgen to generate a mock from the interface.
3. Use the mock in a test:

```
      func TestMyThing(t *testing.T) {
        mockCtrl := gomock.NewController(t)
        defer mockCtrl.Finish()
        mockObj := something.NewMockMyInterface(mockCtrl)
        mockObj.EXPECT().SomeMethod(4, "blah")
        // pass mockObj to a real object and play with it.
      }
```