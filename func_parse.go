package ParserInterface

//###########################################################//

func parseDomainFromUrl(url string) (DomainType, error) {
	return "", ErrParseDomainNil
}

//////////

func parseAuthorFromData(data *LoadWebDataType) (PersonalityObj, error) {
	return PersonalityObj{}, ErrParseAuthorNil
}

func parseWorkFromData(data *LoadWebDataType) (WorkObj, error) {
	return WorkObj{}, ErrParseWorkNil
}

func parseChapterFromData(data *LoadWebDataType) (ChapterObj, error) {
	return ChapterObj{}, ErrParseChapterNil
}
