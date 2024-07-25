package sys

//###########################################################//

func New(url string) error {
	domain, err := Methods.Parse.URL.Domain(url)
	if err != nil {
		return err
	}

	err = Methods.Net.Ping(Methods.Gen.URL.Domain(domain))
	if err != nil {
		return err
	}
	protectParser = &ParserObj{domain: domain}

	id, err := Methods.Parse.URL.Author(url)
	if err == nil {
		protectParser.author = &PersonalityObj{UID: id}
	}

	id, err = Methods.Parse.URL.Work(url)
	if err == nil {
		protectParser.work = &PersonalityObj{UID: id}
	}

	return nil
}

func Close() {
	if protectParser != nil {
		protectParser = nil
	}
}

////

func IsInit() bool {
	return protectParser != nil
}

func IsAuthor() bool {
	if protectParser == nil {
		return false
	}
	return protectParser.author != nil
}

func IsWork() bool {
	if protectParser == nil {
		return false
	}
	return protectParser.work != nil
}

//###########################################################//
