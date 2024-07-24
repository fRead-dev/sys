package ParserInterface

//###########################################################//

func New(url string) error {
	domain, err := Methods.Parse.Domain(url)
	if err != nil {
		return err
	}

	//protectParser = &ParserObj{domain: domain}

	return nil
}
