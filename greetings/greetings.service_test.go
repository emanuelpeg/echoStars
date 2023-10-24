package greetings

import (
	"go.uber.org/mock/gomock"
	"regexp"
	"testing"
)

// TestNewGreeterService call NewGreeterService and check that it doesn't return nil
func TestNewGreeterService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := NewMockGreeter(mockCtrl)

	g := NewGreeterService(mockObj)
	if g == nil {
		t.Fatal("Error: the NewGreeterService should return a service")
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGreeterServiceImp_HelloName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := NewMockGreeter(mockCtrl)

	g := NewGreeterService(mockObj)

	name := "Gladys"
	mockObj.EXPECT().Hello(name).Return("Hi, Gladys. Welcome!")
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := g.Hello("Gladys")
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, %v`, msg, want)
	}
}

// TestServiceHelloEmpty calls greetings.Hello with an empty string.
func TestGreeterServiceImp_HelloEmpty(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := NewMockGreeter(mockCtrl)

	g := NewGreeterService(mockObj)

	mockObj.EXPECT().Hello("").Return("Hi!")
	msg := g.Hello("")
	if msg != "Hi!" {
		t.Fatalf(`Hello("") = %q, want "Hi!"`, msg)
	}
}

// TestServiceHelloEmpty calls greetings.Hello with an empty string.
func TestGreeterServiceImp_HelloWold(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := NewMockGreeter(mockCtrl)

	g := NewGreeterService(mockObj)

	mockObj.EXPECT().Hello("Wold").Return("Hi, Wold. Welcome!")
	msg := g.HelloWold()
	if msg != "Hi, Wold. Welcome!" {
		t.Fatalf(`HelloWold() != Hi, Wold. Welcome!`)
	}
}
