package ParserInterface

//###########################################################//

type GlobalObj struct {
}

func New() *GlobalObj {
	obj := GlobalObj{}
	return &obj
}
