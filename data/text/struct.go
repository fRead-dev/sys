package text

//###########################################################//

type ParagraphObj struct {
	Text     string
	Entities []EntityObj
}

type EntityObj struct {
	Type   FormatType
	Offset int
	Length int
}

type TagObj struct {
	Position int
	IsOpen   bool
	Type     FormatType
}
