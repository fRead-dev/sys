package temp

type EnglishGreeter struct {
	Ax string
	Bx string
}

type Greeter interface {
	Greet(name string) string
}

func (g EnglishGreeter) Greet(name string) string {
	return "Hello, " + name
}
