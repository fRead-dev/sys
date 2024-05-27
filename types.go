package ParserInterface

import "time"

type TimestampObj struct {
	Create time.Time `json:"create"` // creation time.
	Update time.Time `json:"update"` // last updated time.
}

type SizeObj struct {
	Bytes   uint64 `json:"bytes"`   // size in bytes.
	Letters uint64 `json:"letters"` // character count excluding markup.
}

type PersonalityObj struct {
	ID   string `json:"id"`   // unique pointer within the domain.
	Name string `json:"name"` // name/title within the unique pointer.
}

//###########################################################//

// WorkObj Information about the work
type WorkObj struct {
	Work   PersonalityObj `json:"work"`   // Work pointer.
	Author PersonalityObj `json:"author"` // Author pointer.

	Language  string       `json:"language"`  // Work language.
	Timestamp TimestampObj `json:"timestamp"` // Work timestamps.
	Hash      string       `json:"hash"`      // global checksum of the Work

	Tags       []string `json:"tags"`       // specified tags
	Fandoms    []string `json:"fandoms"`    // specified fandoms
	Characters []string `json:"characters"` // specified characters

	Status StatusTag `json:"status"` // status of the Work
	Rating RatingTag `json:"rating"` // rating of the Work
	Focus  FocusTag  `json:"focus"`  // focus of the Work

	Description []byte `json:"description" compress:"true"` // description of the Work
}

// ChapterObj chapter content
type ChapterObj struct {
	Chapter PersonalityObj `json:"chapter"` // Chapter pointer.

	Timestamp TimestampObj `json:"timestamp"` // Chapter timestamps.
	Size      SizeObj      `json:"size"`      // size of the Chapter
	Hash      string       `json:"hash"`      // checksum of the Chapter

	Data []byte `json:"data" compress:"true"` // text of the Chapter
}

//###########################################################//

// GlobalObj Global information object about the work
type GlobalObj struct {
	Info     WorkObj      `json:"info"`     // information about the Work
	Size     SizeObj      `json:"size"`     // size of the Work
	Chapters []ChapterObj `json:"chapters"` //	chapters of the Work
}
