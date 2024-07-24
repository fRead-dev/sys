package ParserInterface

//###########################################################//

func netPing(url string) error {
	return ErrNetPingNil
}

func netLoadAuthor(domain DomainType, author UIDType) (*[]LoadWebDataType, error) {
	return nil, ErrNetAuthorNil
}

func netLoadWork(domain DomainType, work UIDType) (*[]LoadWebDataType, error) {
	return nil, ErrNetWorkNil
}

func netLoadChapter(domain DomainType, work UIDType, chapter UIDType) (*[]LoadWebDataType, error) {
	return nil, ErrNetChapterNil
}

//###########################################################//
