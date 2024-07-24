package ParserInterface

//###########################################################//

func parseDomainFromUrl(url string) (DomainType, error) {
	return "", ErrParseUrlDomainNil
}

func parseAuthorFromUrl(url string) (UIDType, error) {
	return "", ErrParseUrlAuthorNil
}

func parseWorkFromUrl(url string) (UIDType, error) {
	return "", ErrParseUrlWorkNil
}

func parseChapterFromUrl(url string) (UIDType, error) {
	return "", ErrParseUrlChapterNil
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

////

func parseWorkFromAuthorData(data *LoadWebDataType) ([]PersonalityObj, error) {
	return []PersonalityObj{}, ErrParseAuthorWorksNil
}

func parseChapterFromWorkData(data *LoadWebDataType) ([]PersonalityObj, error) {
	return []PersonalityObj{}, ErrParseWorkChaptersNil
}
