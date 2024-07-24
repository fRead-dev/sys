package ParserInterface

//###########################################################//

func Upload() error {
	if !IsInit() {
		return ErrGlobalNotInit
	}

	if IsAuthor() {
		data, err := Methods.Net.Author(protectParser.domain, protectParser.author.UID)
		if err != nil {
			return err
		}
		protectParser.load = &data
	}

	if IsWork() {
		data, err := Methods.Net.Work(protectParser.domain, protectParser.work.UID)
		if err != nil {
			return err
		}
		protectParser.load = &data
	}

	return nil
}

////////

func ReadWorksFromAuthor() (*[]PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if protectParser.load == nil {
		return nil, ErrGlobalNotLoad
	}
	if !IsAuthor() {
		return nil, ErrGlobalIsNotAuthor
	}

	arr, err := Methods.Parse.WorkChapters(protectParser.load)
	return &arr, err
}

func ReadChaptersFromWork() (*[]PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if protectParser.load == nil {
		return nil, ErrGlobalNotLoad
	}
	if !IsWork() {
		return nil, ErrGlobalIsNotWork
	}

	arr, err := Methods.Parse.WorkChapters(protectParser.load)
	return &arr, err
}

//###########################################################//
