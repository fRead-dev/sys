package temp

type EnglishGreeter struct{}

func (g EnglishGreeter) Greet(name string) string {
	return "Hello, " + name
}
