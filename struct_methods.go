package ParserInterface

//###########################################################//

type ParserObj struct {
	domain DomainType
	work   PersonalityObj
	author PersonalityObj
}

//###########################################################//

type MethodsObject struct {
	Parse MethodsParseObject
	Net   MethodsNetObject
	Gen   MethodsGenObject
}

////////////

type MethodsParseObject struct {
	Domain func(url string) (DomainType, error)
	Author func(data *[]LoadWebDataType) (PersonalityObj, error)
}

////////

type MethodsNetObject struct {
	Ping func(url string) error
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
	Methods.Parse.Domain = parseDomainFromUrl
	Methods.Parse.Author = parseAuthorFromData

	Methods.Net.Ping = netPing

	Methods.Gen.URL.Domain = genUrlDomain
	Methods.Gen.URL.Author = genUrlAuthor
	Methods.Gen.URL.Work = genUrlWork
	Methods.Gen.URL.Chapter = genUrlWorkChapter
}
