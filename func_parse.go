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

//////////

func parseFieldChapterPersonalityFromData(data *LoadWebDataType) (PersonalityObj, error) {
	return PersonalityObj{}, ErrParseFieldChapterPersonalityNil
}

func parseFieldChapterTimestampFromData(data *LoadWebDataType) (TimestampObj, error) {
	return TimestampObj{}, ErrParseFieldChapterTimestampNil
}

func parseFieldChapterDataFromData(data *LoadWebDataType) (NormalisedTextType, error) {
	return []byte{}, ErrParseFieldChapterDataNil
}

////

func parseFieldWorkPersonalityFromData(data *LoadWebDataType) (PersonalityObj, error) {
	return PersonalityObj{}, ErrParseFieldWorkPersonalityNil
}

func parseFieldWorkLanguageFromData(data *LoadWebDataType) (LanguageType, error) {
	return "", ErrParseFieldWorkLanguageNil
}

func parseFieldWorkTimestampFromData(data *LoadWebDataType) (TimestampObj, error) {
	return TimestampObj{}, ErrParseFieldWorkTimestampNil
}

//

func parseFieldWorkTagsFromData(data *LoadWebDataType) ([]TagType, error) {
	return []TagType{}, ErrParseFieldWorkTagsNil
}

func parseFieldWorkFandomsFromData(data *LoadWebDataType) ([]FandomType, error) {
	return []FandomType{}, ErrParseFieldWorkFandomsNil
}

func parseFieldWorkPersonagesFromData(data *LoadWebDataType) ([]PersonType, error) {
	return []PersonType{}, ErrParseFieldWorkPersonagesNil
}

//

func parseFieldWorkStatusFromData(data *LoadWebDataType) (StatusTag, error) {
	return 0, ErrParseFieldWorkStatusNil
}

func parseFieldWorkRatingFromData(data *LoadWebDataType) (RatingTag, error) {
	return 0, ErrParseFieldWorkRatingNil
}

func parseFieldWorkFocusFromData(data *LoadWebDataType) (FocusTag, error) {
	return 0, ErrParseFieldWorkFocusNil
}

//

func parseFieldWorkDescriptionFromData(data *LoadWebDataType) (NormalisedTextType, error) {
	return []byte{}, ErrParseFieldWorkDescriptionNil
}
