package sys

//###########################################################//

// WorkObj Information about the work
type WorkObj struct {
	Work PersonalityObj `json:"work"` // Work pointer.

	Language  LanguageType `json:"language"`  // Work language.
	Timestamp TimestampObj `json:"timestamp"` // Work timestamps.
	Hash      HashType     `json:"hash"`      // global checksum of the Work

	Tags       []TagType    `json:"tags"`       // теги
	Fandoms    []FandomType `json:"fandoms"`    // фендомы
	Personages []PersonType `json:"personages"` // персонажи

	Status StatusTag `json:"status"` // status of the Work
	Rating RatingTag `json:"rating"` // rating of the Work
	Focus  FocusTag  `json:"focus"`  // focus of the Work

	Description NormalisedTextType `json:"description"` // description of the Work
}

// ChapterObj chapter content
type ChapterObj struct {
	Chapter PersonalityObj `json:"chapter"` // Chapter pointer.

	Timestamp TimestampObj `json:"timestamp"` // Chapter timestamps.
	Size      SizeObj      `json:"size"`      // size of the Chapter
	Hash      HashType     `json:"hash"`      // checksum of the Chapter

	Data NormalisedTextType `json:"data"` // text of the Chapter
}

//###########################################################//

type BookObj struct {
	Domain DomainType     `json:"domain"` // хост-имя домена
	Author PersonalityObj `json:"author"` // Author pointer.

	Work     WorkObj      `json:"work"`     // information about the Work
	Size     SizeObj      `json:"size"`     // size of the Work
	Chapters []ChapterObj `json:"chapters"` //	chapters of the Work
}
