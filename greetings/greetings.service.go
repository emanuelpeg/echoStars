package greetings

type GreeterService interface {
	Hello(name string) (message string)
	HelloWold() (message string)
}

type GreeterServiceImp struct {
	greeter Greeter
}

func NewGreeterService(greeter Greeter) GreeterService {
	svc := GreeterServiceImp{
		greeter: greeter,
	}
	return svc
}

// Hello returns a greeting for the named person.
func (g GreeterServiceImp) Hello(name string) (message string) {
	// Call the model method
	return g.greeter.Hello(name)
}

// HelloWold returns a greeting for the wold.
func (g GreeterServiceImp) HelloWold() (message string) {
	// Call the model method with parameter Wold
	return g.greeter.Hello("Wold")
}
