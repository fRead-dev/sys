package ParserInterface

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

//###########################################################//
