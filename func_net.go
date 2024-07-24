package ParserInterface

//###########################################################//

func netPing(url string) error {
	return ErrNetPingNil
}

func netLoadContent(url string) (*[]LoadWebDataType, error) {
	return nil, ErrNetLoadNil
}

//###########################################################//

func (net *MethodsNetObject) Author(domain DomainType, author UIDType) (*[]LoadWebDataType, error) {
	return nil, nil
}
