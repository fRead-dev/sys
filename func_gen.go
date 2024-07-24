package ParserInterface

import (
	"bytes"
)

//###########################################################//

func genUrlDomain(domain DomainType) string {
	return "https://" + string(domain)
}

func genUrlAuthor(domain DomainType, author UIDType) string {
	return "https://" + string(domain) + "/" + string(author)
}

func genUrlWork(domain DomainType, work UIDType) string {
	return "https://" + string(domain) + "/" + string(work)
}

func genUrlWorkChapter(domain DomainType, work UIDType, chapter UIDType) string {
	return "https://" + string(domain) + "/" + string(work) + "/" + string(chapter)
}

////////

func genNormalisedText(data LoadWebDataType) (NormalisedTextType, error) {
	return []byte{}, ErrGenTextLettersNil
}

func genLettersSum(data *NormalisedTextType) (uint64, error) {
	return 0, ErrGenTextNormalisedNil
}

////

func (text *MethodsGenTextObject) Hash(data *NormalisedTextType) HashType {
	var reader bytes.Reader

	_, err := reader.Read(*data)
	if err != nil {
		return ""
	}

	return Methods.SYS.Hash(&reader)
}
