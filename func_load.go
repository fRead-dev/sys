package ParserInterface

//###########################################################//

func LoadWorksFromAuthor() (*[]PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if !IsAuthor() {
		return nil, ErrGlobalIsNotAuthor
	}

	arr, err := Methods.Net.AuthorWorksData(protectParser.domain, protectParser.author.UID)
	return &arr, err
}

func LoadChapterFromWork() (*[]PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if !IsWork() {
		return nil, ErrGlobalIsNotAuthor
	}

	arr, err := Methods.Net.AuthorWorksData(protectParser.domain, protectParser.work.UID)
	return &arr, err
}
