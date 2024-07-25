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

type MethodsSysObject struct{}

////////////

type MethodsParseObject struct {
	URL   MethodsParseUrlObject
	Field MethodsParseFieldObject

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
type MethodsParseFieldObject struct {
	Work    MethodsParseFieldWorkObject
	Chapter MethodsParseFieldChapterObject
}

type MethodsParseFieldWorkObject struct {
	Personality func(data *LoadWebDataType) (PersonalityObj, error)
	Timestamp   func(data *LoadWebDataType) (TimestampObj, error)
	Language    func(data *LoadWebDataType) (LanguageType, error)

	Tags       func(data *LoadWebDataType) ([]TagType, error)
	Fandoms    func(data *LoadWebDataType) ([]FandomType, error)
	Personages func(data *LoadWebDataType) ([]PersonType, error)

	Status func(data *LoadWebDataType) (StatusTag, error)
	Rating func(data *LoadWebDataType) (RatingTag, error)
	Focus  func(data *LoadWebDataType) (FocusTag, error)

	Description func(data *LoadWebDataType) (NormalisedTextType, error)
}

type MethodsParseFieldChapterObject struct {
	Personality func(data *LoadWebDataType) (PersonalityObj, error)
	Timestamp   func(data *LoadWebDataType) (TimestampObj, error)
	Data        func(data *LoadWebDataType) (NormalisedTextType, error)
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
	URL  MethodsGenUrlObject
	Text MethodsGenTextObject
}

type MethodsGenUrlObject struct {
	Domain  func(domain DomainType) string
	Author  func(domain DomainType, author UIDType) string
	Work    func(domain DomainType, work UIDType) string
	Chapter func(domain DomainType, work UIDType, chapter UIDType) string
}

type MethodsGenTextObject struct {
	Normalised func(data LoadWebDataType) (NormalisedTextType, error)
	Letters    func(data *NormalisedTextType) (uint64, error)
	Bytes      func(data *NormalisedTextType) uint64
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

		Methods.Parse.Field.Chapter.Personality = parseFieldChapterPersonalityFromData
		Methods.Parse.Field.Chapter.Timestamp = parseFieldChapterTimestampFromData
		Methods.Parse.Field.Chapter.Data = parseFieldChapterDataFromData

		Methods.Parse.Field.Work.Personality = parseFieldWorkPersonalityFromData
		Methods.Parse.Field.Work.Timestamp = parseFieldWorkTimestampFromData
		Methods.Parse.Field.Work.Language = parseFieldWorkLanguageFromData
		Methods.Parse.Field.Work.Tags = parseFieldWorkTagsFromData
		Methods.Parse.Field.Work.Fandoms = parseFieldWorkFandomsFromData
		Methods.Parse.Field.Work.Personages = parseFieldWorkPersonagesFromData
		Methods.Parse.Field.Work.Status = parseFieldWorkStatusFromData
		Methods.Parse.Field.Work.Rating = parseFieldWorkRatingFromData
		Methods.Parse.Field.Work.Focus = parseFieldWorkFocusFromData
		Methods.Parse.Field.Work.Description = parseFieldWorkDescriptionFromData
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

		Methods.Gen.Text.Letters = genLettersSum
		Methods.Gen.Text.Normalised = genNormalisedText
		Methods.Gen.Text.Bytes = genBytesSum
	}
}
