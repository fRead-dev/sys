package sys

//###########################################################//

func netPing(url string) error {
	return ErrNetPingNil
}

func netLoadAuthor(domain DomainType, author UIDType) (LoadWebDataType, error) {
	return nil, ErrNetAuthorNil
}
func (net *MethodsNetObject) AuthorData(domain DomainType, author UIDType) (PersonalityObj, error) {
	data, err := net.Author(domain, author)
	if err != nil {
		return PersonalityObj{}, err
	}
	return Methods.Parse.Author(&data)
}
func (net *MethodsNetObject) AuthorWorksData(domain DomainType, author UIDType) ([]PersonalityObj, error) {
	data, err := net.Author(domain, author)
	if err != nil {
		return []PersonalityObj{}, err
	}
	return Methods.Parse.AuthorWorks(&data)
}

func netLoadWork(domain DomainType, work UIDType) (LoadWebDataType, error) {
	return nil, ErrNetWorkNil
}
func (net *MethodsNetObject) WorkData(domain DomainType, work UIDType) (WorkObj, error) {
	data, err := net.Work(domain, work)
	if err != nil {
		return WorkObj{}, err
	}
	return Methods.Parse.Work(&data)
}
func (net *MethodsNetObject) WorkChaptersData(domain DomainType, author UIDType) ([]PersonalityObj, error) {
	data, err := net.Work(domain, author)
	if err != nil {
		return []PersonalityObj{}, err
	}
	return Methods.Parse.WorkChapters(&data)
}

func netLoadChapter(domain DomainType, work UIDType, chapter UIDType) (LoadWebDataType, error) {
	return nil, ErrNetChapterNil
}
func (net *MethodsNetObject) ChapterData(domain DomainType, work UIDType, chapter UIDType) (ChapterObj, error) {
	data, err := net.Chapter(domain, work, chapter)
	if err != nil {
		return ChapterObj{}, err
	}
	return Methods.Parse.Chapter(&data)
}

//###########################################################//
