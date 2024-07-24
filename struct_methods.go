package ParserInterface

//###########################################################//

type ParserObj struct {
	domain DomainType
	work   *PersonalityObj
	author *PersonalityObj

	load *LoadWebDataType
}

//###########################################################//

type MethodsObject struct {
	SYS MethodsSysObject

	Parse MethodsParseObject
	Net   MethodsNetObject
	Gen   MethodsGenObject
}

type MethodsSysObject struct {
}

////////////

type MethodsParseObject struct {
	URL MethodsParseUrlObject

	Author  func(data *LoadWebDataType) (PersonalityObj, error)
	Work    func(data *LoadWebDataType) (WorkObj, error)
	Chapter func(data *LoadWebDataType) (ChapterObj, error)

	AuthorWorks  func(data *LoadWebDataType) ([]PersonalityObj, error)
	WorkChapters func(data *LoadWebDataType) ([]PersonalityObj, error)
}

type MethodsParseUrlObject struct {
	Domain  func(url string) (DomainType, error)
	Author  func(url string) (UIDType, error)
	Work    func(url string) (UIDType, error)
	Chapter func(url string) (UIDType, error)
}

////////

type MethodsNetObject struct {
	Ping func(url string) error

	Author  func(domain DomainType, author UIDType) (LoadWebDataType, error)
	Work    func(domain DomainType, work UIDType) (LoadWebDataType, error)
	Chapter func(domain DomainType, work UIDType, chapter UIDType) (LoadWebDataType, error)
}

////////

type MethodsGenObject struct {
	URL MethodsGenUrlObject
}

type MethodsGenUrlObject struct {
	Domain  func(domain DomainType) string
	Author  func(domain DomainType, author UIDType) string
	Work    func(domain DomainType, work UIDType) string
	Chapter func(domain DomainType, work UIDType, chapter UIDType) string
}

//###########################################################//

func init() {
	{
		Methods.Parse.URL.Domain = parseDomainFromUrl
		Methods.Parse.URL.Author = parseAuthorFromUrl
		Methods.Parse.URL.Work = parseWorkFromUrl
		Methods.Parse.URL.Chapter = parseChapterFromUrl

		Methods.Parse.Author = parseAuthorFromData
		Methods.Parse.Work = parseWorkFromData
		Methods.Parse.Chapter = parseChapterFromData

		Methods.Parse.AuthorWorks = parseWorkFromAuthorData
		Methods.Parse.WorkChapters = parseChapterFromWorkData
	}
	{
		Methods.Net.Ping = netPing

		Methods.Net.Author = netLoadAuthor
		Methods.Net.Work = netLoadWork
		Methods.Net.Chapter = netLoadChapter
	}
	{
		Methods.Gen.URL.Domain = genUrlDomain
		Methods.Gen.URL.Author = genUrlAuthor
		Methods.Gen.URL.Work = genUrlWork
		Methods.Gen.URL.Chapter = genUrlWorkChapter
	}
}
