package ParserInterface

//###########################################################//

func parseDomainFromUrl(url string) (DomainType, error) {
	return "", ErrParseDomainNil
}

func parseAuthorFromData(data *[]LoadWebDataType) (PersonalityObj, error) {
	return PersonalityObj{}, ErrParseAuthorNil
}
